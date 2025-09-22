package main

import (
	"log"
	"net/http"

	"github.com/dakshesh14/golinky/internal/config"
	"github.com/dakshesh14/golinky/internal/container"
	"github.com/dakshesh14/golinky/internal/infrastructure/cache"
	"github.com/dakshesh14/golinky/internal/infrastructure/db"
	"github.com/dakshesh14/golinky/internal/routes"
	"github.com/go-chi/chi/v5"
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

	container := container.NewContainer(database, cache, cfg)

	router := chi.NewRouter()

	routes.SetupRoutes(router, container)

	log.Printf("Starting server on port %s", cfg.AppPort)
	log.Fatal(http.ListenAndServe(":"+cfg.AppPort, router))
}
