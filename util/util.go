package util

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func RequiredEnvVars() {

	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	vars := []string{
		"POSTGRES_USER",
		"POSTGRES_PASSWORD",
		"POSTGRES_DB",
	}

	for _, name := range vars {
		value := os.Getenv(name)
		if value == "" {
			log.Fatalf("Error: Missing required environment variable '%s'", name)
		}
	}
}
