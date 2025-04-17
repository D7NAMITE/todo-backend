package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	DatabaseURL string
}

func LoadConfig() *Config {
	// Get .env secret
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	return &Config{
		DatabaseURL: os.Getenv("DATABASE_URL"),
	}
}
