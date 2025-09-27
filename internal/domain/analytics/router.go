package analytics

import (
	"github.com/go-chi/chi/v5"

	"github.com/dakshesh14/golinky/internal/container"
	"github.com/dakshesh14/golinky/internal/infrastructure/db/repository"
)

func SetupRoutes(container container.Container) *chi.Mux {
	router := chi.NewRouter()

	linkRepository := repository.NewLinkRepository(container.Db)
	linkClickRepository := repository.NewLinkClickRepository(container.Db)

	handler := NewHandler(container, linkRepository, linkClickRepository)

	router.Get("/links/{linkCode}/analysis", handler.GetLinkAnalysis)

	return router
}
