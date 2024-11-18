package main

import (
	"GoPorject/api"
	"GoPorject/config"
	"GoPorject/lib/modules/controllers"
	"GoPorject/lib/modules/services"
	"GoPorject/lib/routes"
	"GoPorject/log"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	cfg := config.LoadConfig()
	db := config.Connect(cfg)
	logger.Info("Starting Go application")
	userService := services.NewUserService()
	userController := controllers.NewUserController(userService)
	router := gin.Default()
	routes.RegisterRoutes(router, userController)

	if err := api.StartServer(cfg, db); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
