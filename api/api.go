package api

import (
	"GoPorject/config"
	"GoPorject/log"
	"database/sql"
	"fmt"
	"net/http"
)

func StartServer(cfg *config.Config, db *sql.DB) error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to GoProject!")

	})
	address := fmt.Sprintf(":%s", cfg.ServerPort)
	if cfg.ServerPort == "" {
		logger.Warning("Port is not available.")
	}
	logger.Info("Go application successfully started")

	return http.ListenAndServe(address, nil)
}
