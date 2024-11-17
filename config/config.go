package config

import (
	"GoPorject/log"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type Config struct {
	DBUser     string
	DBPass     string
	DBName     string
	DBHost     string
	DBPort     string
	ServerPort string
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		logger.Warning("Warning: .env file not found, using environment variables")
	}

	return &Config{
		DBUser:     os.Getenv("DB_USER"),
		DBPass:     os.Getenv("DB_PASS"),
		DBName:     os.Getenv("DB_NAME"),
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		ServerPort: os.Getenv("SERVER_PORT"),
	}
}

func Connect(cfg *Config) *sql.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.DBUser, cfg.DBPass, cfg.DBHost, cfg.DBPort, cfg.DBName)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to connect to the database: %v", err))
	}
	logger.Info("Database connection initialized")
	return db
}
