package routes

import (
	"github.com/dakshesh14/golinky/internal/container"
	"github.com/dakshesh14/golinky/internal/domain/analytics"
	"github.com/dakshesh14/golinky/internal/domain/healthcheck"
	"github.com/dakshesh14/golinky/internal/domain/link"
	"github.com/go-chi/chi/v5"
)

func SetupRoutes(router *chi.Mux, container *container.Container) *chi.Mux {
	router.Mount("/api/healthcheck", healthcheck.SetupRoutes(*container))
	router.Mount("/", link.SetupRoutes(*container))
	router.Mount("/api/analytics", analytics.SetupRoutes(*container))

	return router
}
