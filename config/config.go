package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	AppPort          string
	PostgresUser     string
	PostgresPassword string
	PostgresDB       string
	PostgresHost     string
	PostgresPort     string
	PostgresSSLMode  string
}

func LoadConfig() *Config {
	if err := godotenv.Load(".env.example"); err != nil {
		log.Println("No .env.example file found, continue with default values")
	}

	return &Config{
		AppPort:          getEnv("APP_PORT", "3000"),
		PostgresUser:     getEnv("POSTGRES_USER", "postgres"),
		PostgresPassword: getEnv("POSTGRES_PASSWORD", "postgres"),
		PostgresDB:       getEnv("POSTGRES_DB", "postgres"),
		PostgresHost:     getEnv("POSTGRES_HOST", "db"),
		PostgresPort:     getEnv("POSTGRES_PORT", "5432"),
		PostgresSSLMode:  getEnv("POSTGRES_SSLMODE", "disable"),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue
}
