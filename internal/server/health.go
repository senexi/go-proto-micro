package server
import(
    "github.com/senexi/go-proto-micro/internal/db"
)
type HealthStatus struct {
	Status  bool `json:"status"`
	StatusDB   bool `json:"db"`
}


func (s *Server) Health() HealthStatus {
    return HealthStatus{
		Status: true,
		StatusDB: db.Health(),
	}
}