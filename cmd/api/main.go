package main

import (
	"log"

	"github.com/dakshesh14/golinky/internal/config"
	"github.com/dakshesh14/golinky/internal/infrastructure/cache"
	"github.com/dakshesh14/golinky/internal/infrastructure/db"
)

func main() {
	cfg, err := config.LoadConfig()

	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	_, err = db.InitDb(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to db: %v", err)
	}

	_, err = cache.NewCacheService(cfg.RedisURL)
	if err != nil {
		log.Fatalf("Failed to connect to cache: %v", err)
	}

	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

}
