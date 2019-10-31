package management

import (
)

type HealthStatus struct {
	Status   bool `json:"status"`
}

func (s *Server) Health() HealthStatus {
	return HealthStatus{
		Status:   true,
	}
}
