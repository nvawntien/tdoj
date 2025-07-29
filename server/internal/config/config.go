package config

import (
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	Redis    RedisConfig
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

type RedisConfig struct {
	Rd       *redis.Client
	Host     string
	Port     string
	Password string
	Db       string
}

func Load() (*Config, error) {
	return &Config{
		Server: ServerConfig{
			Port: os.Getenv("PORT"),
			Env:  os.Getenv("ENV"),
		},

		Database: DatabaseConfig{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			Username: os.Getenv("DB_USERNAME"),
			Password: os.Getenv("DB_PASSWORD"),
			DbName:   os.Getenv("DB_NAME"),
		},

		Redis: RedisConfig{
			Host:     os.Getenv("REDIS_HOST"),
			Port:     os.Getenv("REDIS_PORT"),
			Password: os.Getenv("REDIS_PASSWORD"),
			Db:       os.Getenv("REDIS_DB"),
		},
	}, nil
}
