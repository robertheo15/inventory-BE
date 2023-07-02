package config

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvFile() {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file")
	}

	log.Println("Success load .env file")
}
