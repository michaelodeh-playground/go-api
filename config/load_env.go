package config

import (
	"log"

	"github.com/joho/godotenv"
)

func Load() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

}
