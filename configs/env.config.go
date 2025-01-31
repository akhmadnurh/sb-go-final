package configs

import (
	"OneTix/structs"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() *structs.Env {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	return &structs.Env{
		AppName:   os.Getenv("APP_NAME"),
		AppPort:   os.Getenv("APP_PORT"),
		JWTSecret: os.Getenv("JWT_SECRET"),
		DocsURL:   os.Getenv("DOCS_URL"),
		DBURL:     os.Getenv("DATABASE_URL"),
	}
}
