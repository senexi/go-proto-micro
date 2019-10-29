package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var logLevel string
var jsonLogging bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-proto-micro",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

    rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ./config.yaml)")
    rootCmd.PersistentFlags().StringVarP(&logLevel, "verbosity", "v", log.InfoLevel.String(), "Log level (debug, info, warn, error, fatal, panic")
    viper.BindPFlag("logging.verbosity", rootCmd.PersistentFlags().Lookup("verbosity"))
    rootCmd.PersistentFlags().BoolVarP(&jsonLogging, "json-logging", "j", false, "enable logging in json format")
    viper.BindPFlag("logging.json-logging", rootCmd.PersistentFlags().Lookup("json-logging"))

    rootCmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
        setLogFormat()
		if err := setLogLevel(); err != nil {
			return err
        }
		return nil
    }

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
    // rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		ex, err := os.Executable()
		if err != nil {
			panic(err)
		}
		exPath := filepath.Dir(ex)
		log.Info(exPath)
		viper.AddConfigPath(exPath)
		viper.SetConfigName("config")
	}
	viper.SetEnvPrefix("gmp")
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err == nil {
		log.WithFields(log.Fields{"config_file": viper.ConfigFileUsed()}).Info("using config")
	} else {
		log.Warn(err)
	}
}

//setUpLogs set the log output ans the log level
func setLogLevel() error {
    level := viper.GetString("logging.verbosity")
	lvl, err := log.ParseLevel(level)
	if err != nil {
		return err
	}
    log.SetLevel(lvl)
    log.WithFields(log.Fields{"loglevel": logLevel}).Info("set loglevel")
	return nil
}

func setLogFormat() {
    if(viper.GetBool("logging.json-logging")) {
        log.SetFormatter(&log.JSONFormatter{})
    }
}