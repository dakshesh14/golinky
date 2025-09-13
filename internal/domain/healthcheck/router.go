package healthcheck

import (
	"github.com/dakshesh14/golinky/internal/container"
	"github.com/go-chi/chi/v5"
)

func SetupRoutes(container container.Container) *chi.Mux {
	router := chi.NewRouter()

	handler := NewHandler(container)

	router.Get("/", handler.HealthCheckHandler)

	return router
}
