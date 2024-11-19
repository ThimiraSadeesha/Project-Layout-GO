package api

import (
	"GoPorject/config"
	"GoPorject/log"
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

func StartServer(cfg *config.Config, db *sql.DB) error {
	if cfg.ServerPort == "" {
		log.Fatal("Server port not provided in config.")
	}

	address := fmt.Sprintf(":%s", cfg.ServerPort)
	logger.Info("Starting server on " + address)
	return http.ListenAndServe(address, nil)
}
