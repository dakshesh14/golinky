package routes

import (
	"github.com/dakshesh14/golinky/internal/container"
	"github.com/dakshesh14/golinky/internal/domain/healthcheck"
	"github.com/go-chi/chi/v5"
)

func SetupRoutes(router *chi.Mux, container *container.Container) *chi.Mux {
	router.Mount("/api/healthcheck", healthcheck.SetupRoutes(*container))

	return router
}
