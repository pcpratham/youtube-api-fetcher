package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL   string
	YouTubeAPIKey string
	Port          string
}

func Load() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file: %v", err)
	}

	return &Config{
		DatabaseURL:   os.Getenv("DATABASE_URL"),
		YouTubeAPIKey: os.Getenv("YOUTUBE_API_KEY"),
		Port:          getEnvWithDefault("PORT", "8080"),
	}
}

func getEnvWithDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}