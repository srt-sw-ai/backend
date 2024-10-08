package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
    DBHost     string
    DBUser     string
    DBPassword string
    DBName     string
    JWTSecret  string
}

func LoadConfig() (config Config, err error) {
    err = godotenv.Load(".env")
    if err != nil {
        log.Fatalf("Error loading .env file")
        return
    }

    config = Config{
        DBHost:     os.Getenv("DB_HOST"),
        DBUser:     os.Getenv("DB_USER"),
        DBPassword: os.Getenv("DB_PASSWORD"),
        DBName:     os.Getenv("DB_NAME"),
        JWTSecret:  os.Getenv("JWT_SECRET"),
    }

    return
}