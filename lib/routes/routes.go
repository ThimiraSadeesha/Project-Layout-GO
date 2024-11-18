package routes

import (
	"GoPorject/lib/modules/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, userController *controllers.UserController) {
	api := router.Group("/api")
	api.GET("/users", userController.GetWelcomeMessage)
}
