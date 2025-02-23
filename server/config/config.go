package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	DB_URL     string
	JWT_SECRET string
)

func LoadConfig() {
	err := godotenv.Load()
	// TODO: Load() doesnt override but i want it to. i think. not sure.
	if err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	DB_URL = os.Getenv("DATABASE_URL")
	JWT_SECRET = os.Getenv("JWT_SECRET")

	if DB_URL == "" || JWT_SECRET == "" {
		log.Fatal("Missing required environment variables")
	}
}
