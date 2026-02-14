package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppEnv        string
	UserGrpcAddr  string
	OrderHTTPPort string
	UserGrpcPort  string
}

func Load() *Config {
	_ = godotenv.Load()

	cfg := &Config{
		AppEnv:        os.Getenv("APP_ENV"),
		UserGrpcAddr:  os.Getenv("USER_GRPC_ADDR"),
		OrderHTTPPort: os.Getenv("ORDER_HTTP_PORT"),
		UserGrpcPort:  os.Getenv("USER_GRPC_PORT"),
	}

	if cfg.AppEnv == "" {
		log.Fatal("APP_ENV is required")
	}

	return cfg
}
