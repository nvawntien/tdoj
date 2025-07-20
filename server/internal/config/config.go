package config

import (
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
}

type ServerConfig struct {
	Port string
	Env  string
}

type DatabaseConfig struct {
	Db       *sqlx.DB
	Host     string
	Port     string
	Username string
	Password string
	DbName   string
}

func Load() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	return &Config{
		Server: ServerConfig{
			Port: getEnv("PORT", "8080"),
			Env:  getEnv("Env", "development"),
		},

		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "1526"),
			Username: getEnv("DB_USERNAME", "vantien1526"),
			Password: getEnv("DB_PASSWORD", "lpop3kiss"),
			DbName:   getEnv("DB_NAME", "tamduong"),
		},
	}, nil
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return fallback
}
