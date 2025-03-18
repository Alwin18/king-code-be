package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost          string
	DBUser          string
	DBPassword      string
	DBName          string
	DBPort          string
	ServerPort      string
	SetMaxIdleConns string
	SetMaxOpenConns string
	SetMaxLifeTime  string
	SSLMode         string
}

func LoadEnv() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	config := &Config{}
	envVars := map[string]*string{
		"DB_HOST":      &config.DBHost,
		"DB_PORT":      &config.DBPort,
		"DB_USER":      &config.DBUser,
		"DB_PASSWORD":  &config.DBPassword,
		"DB_NAME":      &config.DBName,
		"SERVER_PORT":  &config.ServerPort,
		"SET_MAX_IDLE": &config.SetMaxIdleConns,
		"SET_MAX_OPEN": &config.SetMaxOpenConns,
		"SET_MAX_LIFE": &config.SetMaxLifeTime,
		"SSLMODE":      &config.SSLMode,
	}

	for key, ptr := range envVars {
		value := os.Getenv(key)
		if value == "" {
			log.Printf("Missing environment variable: %s", key)
		}
		*ptr = value
	}

	return config, nil
}
