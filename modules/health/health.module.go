package health

import (
	"github.com/gin-gonic/gin"
)

type HealthModule struct {
	Controller *HealthController
}

func NewHealthModule() *HealthModule {
	service := NewHealthService()
	controller := NewHealthController(service)

	return &HealthModule{
		Controller: controller,
	}
}

func (m *HealthModule) RegisterRoutes(router *gin.RouterGroup) {
	router.GET("/health", m.Controller.GetHealth)
}
