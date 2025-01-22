package config

import (
	"log"
	"os"
	"strconv"

	"github.com/Okemwag/medihub/pkg/database"
	"github.com/joho/godotenv"
)

func LoadConfig() database.Config {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables directly")
	}

	return database.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     atoi(os.Getenv("DB_PORT")),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}
}

func atoi(value string) int {
	v, err := strconv.Atoi(value)
	if err != nil {
		log.Fatalf("Invalid integer value for config: %v", err)
	}
	return v
}
