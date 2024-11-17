package main

import (
	"GoPorject/api"
	"GoPorject/config"
	"GoPorject/log"
	"log"
)

func main() {
	cfg := config.LoadConfig()
	db := config.Connect(cfg)
	defer db.Close()
	logger.Info("Starting Go application")
	if err := api.StartServer(cfg, db); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
