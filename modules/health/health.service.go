package health

import "context"

type HealthService struct{}

func NewHealthService() *HealthService {
	return &HealthService{}
}

func (s *HealthService) CheckHealth(ctx context.Context) string {
	return "Service is healthy"
}
