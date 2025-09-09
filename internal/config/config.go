package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppEnv      string
	AppPort     string
	DatabaseURL string
	RedisURL    string
}

func LoadConfig() (*Config, error) {
	_ = godotenv.Load()

	databaseURL, err := MustGetEnv("DATABASE_URL")
	if err != nil {
		return nil, err
	}

	redisURL, err := MustGetEnv("REDIS_URL")
	if err != nil {
		return nil, err
	}

	cfg := &Config{
		AppEnv:      GetEnvOrDefault("APP_ENVIRONMENT", "local"),
		AppPort:     GetEnvOrDefault("APP_PORT", "8080"),
		DatabaseURL: databaseURL,
		RedisURL:    redisURL,
	}

	return cfg, nil
}

func MustGetEnv(key string) (string, error) {
	value := os.Getenv(key)
	if value == "" {
		return "", fmt.Errorf("environment variable %s not set", key)
	}

	return value, nil
}

func GetEnvOrDefault(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}

	return value
}
