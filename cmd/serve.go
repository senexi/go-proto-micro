package cmd

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/rakyll/statik/fs"
	pb "github.com/senexi/go-proto-micro/generated/api"
	_ "github.com/senexi/go-proto-micro/generated/statik"
	"github.com/senexi/go-proto-micro/internal/providers/db"
	s "github.com/senexi/go-proto-micro/internal/providers/management"
	ps "github.com/senexi/go-proto-micro/internal/providers/service"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "run server",
	Long:  `Run the server`,
	Run: func(cmd *cobra.Command, args []string) {
		go runHTTPServer()
        go runGrpcGateway()
        if (mock) {
            runMockServer()
        } else {
            runServer()
        }
	},
}

var server s.Server
var mock bool

func init() {
    serveCmd.Flags().BoolVarP(&mock, "mock", "m", false, "run server in mock mode")
	rootCmd.AddCommand(serveCmd)
}

func runServer() {
	port := viper.GetString("server.port")
	database := viper.GetString("database.name")
	databasePort := viper.GetString("database.port")
	user := viper.GetString("database.user")
	password := viper.GetString("database.password")
	databaseURL := viper.GetString("database.url")

	db.Connect(database, user, password, fmt.Sprintf("%s:%s", databaseURL, databasePort))

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	log.WithFields(log.Fields{"port": port}).Info("started grpc server")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	partnerService := ps.NewPartnerService()
	pb.RegisterPartnerServiceServer(s, partnerService)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func runMockServer() {
    port := viper.GetString("server.port")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	log.WithFields(log.Fields{"port": port}).Info("started MOCKED grpc server")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	partnerService := ps.NewMockPartnerService()
	pb.RegisterPartnerServiceServer(s, partnerService)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func runGrpcGateway() {
	port := viper.GetString("server.gateway.port")

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := pb.RegisterPartnerServiceHandlerFromEndpoint(ctx, mux, fmt.Sprintf(":%s", port), opts)
	if err != nil {
		log.Error(err)
	}
	log.WithFields(log.Fields{"port": port}).Info("serving grpc gateway")
	http.ListenAndServe(fmt.Sprintf(":%s", port), mux)
}

func runHTTPServer() {
	port := viper.GetString("server.web.port")
	log.WithFields(log.Fields{"port": port}).Info("serving http endpoint")
	health()
	metrics()
	web()
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}

func web() {
	endpoint := "/web/"
	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}

	http.Handle(endpoint, http.StripPrefix(endpoint, http.FileServer(statikFS)))
	log.WithFields(log.Fields{"endpoint": fmt.Sprintf("%sswagger-ui", endpoint)}).Info("serving swagger-ui")
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	js, err := json.Marshal(server.Health())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func health() {
	endpoint := "/management/health"
	http.HandleFunc(endpoint, healthHandler)
	log.WithFields(log.Fields{"endpoint": endpoint}).Info("serving health")
}

func metrics() {
	endpoint := "/management/metrics"
	http.Handle(endpoint, promhttp.Handler())
	log.WithFields(log.Fields{"endpoint": endpoint}).Info("serving metrics")
}
