package routes

import (
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"

	"github.com/dakshesh14/golinky/internal/container"
	"github.com/dakshesh14/golinky/internal/domain/analytics"
	"github.com/dakshesh14/golinky/internal/domain/healthcheck"
	"github.com/dakshesh14/golinky/internal/domain/link"
	lmiddleware "github.com/dakshesh14/golinky/internal/middleware"
)

func SetupRoutes(router *chi.Mux, container *container.Container) *chi.Mux {
	if container.Config.RateLimitEnabled {
		router.Use(lmiddleware.GlobalRateLimit(container.Config.RateLimitPerMin))
	}

	if container.Config.AppEnv == "local" {
		router.Use(middleware.Logger)
	}

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Recoverer)

	router.Mount("/api/healthcheck", healthcheck.SetupRoutes(*container))
	router.Mount("/", link.SetupRoutes(*container))
	router.Mount("/api/analytics", analytics.SetupRoutes(*container))

	return router
}
