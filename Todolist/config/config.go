package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	POSTGRES_CONNECTION string `env:"POSTGRES_CONNECTION"`
	PORT                string `env:"PORT"`
}

func LoadEnv() (*Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("No .env file found or error loading it (that's okay if using real env vars)")
	}

	POSTGRES_CONNECTION := os.Getenv("POSTGRES_CONNECTION")
	if POSTGRES_CONNECTION == "" {
		return nil, fmt.Errorf("SERVICES is required")
	}
	PORT := os.Getenv("PORT")
	if PORT == "" {
		panic("PORT is not set")
	}

	cfg := &Config{
		POSTGRES_CONNECTION: POSTGRES_CONNECTION,
		PORT:                PORT,
	}

	return cfg, nil
}
