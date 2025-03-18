package main

import (
	"GoPorject/config"
	"GoPorject/route"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	cfg := config.LoadConfig()
	db := config.Connect(cfg)
	defer db.Close()
	r := gin.Default()
	gin.SetMode(gin.DebugMode)
	err := r.SetTrustedProxies([]string{"127.0.0.1", "192.168.23.230"})
	if err != nil {
		log.Fatalf("Error setting trusted proxies: %v", err)
	}
	route.SetupRoutes(r)
	r.Run(":4000")
}
