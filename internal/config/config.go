package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	AppEnv           string
	AppPort          string
	DatabaseURL      string
	RedisURL         string
	RateLimitEnabled bool
	RateLimitPerMin  int
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
		AppEnv:           GetEnvOrDefault("APP_ENVIRONMENT", "local"),
		AppPort:          GetEnvOrDefault("APP_PORT", "8080"),
		DatabaseURL:      databaseURL,
		RedisURL:         redisURL,
		RateLimitEnabled: GetEnvOrDefaultBool("RATE_LIMIT_ENABLED", false),
		RateLimitPerMin:  GetEnvOrDefaultInt("RATE_LIMIT_PER_MIN", 60),
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

func GetEnvOrDefaultBool(key string, defaultValue bool) bool {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}

	return value == "1" || strings.ToLower(value) == "true"
}

func GetEnvOrDefaultInt(key string, defaultValue int) int {
	valueStr := os.Getenv(key)
	if valueStr == "" {
		return defaultValue
	}

	value, err := strconv.Atoi(valueStr)
	if err != nil {
		log.Panicf("invalid value provided: %q", valueStr)
	}

	return value
}
