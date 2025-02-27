package api

import (
	"GoPorject/config"
	_ "GoPorject/log"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func StartServer(cfg *config.Config, db *sql.DB, r *gin.Engine) error {
	if cfg.ServerPort == "" {
		log.Fatal("Server port not provided in config.")
	}

	// Use Gin's method to run the server
	address := fmt.Sprintf(":%s", cfg.ServerPort)
	log.Printf("Starting server on %s\n", address) // You can use your logger here if you prefer
	return r.Run(address)
}
