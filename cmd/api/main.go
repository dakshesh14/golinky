package main

import (
	"log"

	"github.com/dakshesh14/golinky/internal/config"
	"github.com/dakshesh14/golinky/internal/container"
	"github.com/dakshesh14/golinky/internal/infrastructure/cache"
	"github.com/dakshesh14/golinky/internal/infrastructure/db"
)

func main() {
	cfg, err := config.LoadConfig()

	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	database, err := db.InitDb(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to db: %v", err)
	}

	cache, err := cache.NewCacheService(cfg.RedisURL)
	if err != nil {
		log.Fatalf("Failed to connect to cache: %v", err)
	}

	container.NewContainer(database, cache, cfg)

	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
}
