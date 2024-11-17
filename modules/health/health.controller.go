package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthController struct {
	service *HealthService
}

func NewHealthController(service *HealthService) *HealthController {
	return &HealthController{service: service}
}

func (c *HealthController) GetHealth(ctx *gin.Context) {
	healthStatus := c.service.CheckHealth(ctx)
	ctx.JSON(http.StatusOK, gin.H{"status": healthStatus})
}
