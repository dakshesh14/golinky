package link

import (
	"github.com/dakshesh14/golinky/internal/container"
	"github.com/dakshesh14/golinky/internal/infrastructure/db/repository"
	"github.com/go-chi/chi/v5"
)

func SetupRoutes(container container.Container) *chi.Mux {
	router := chi.NewRouter()

	linkRepository := repository.NewLinkRepository(container.Db)

	handler := NewHandler(container, linkRepository)

	router.Post("/api/links", handler.CreateShortURL)
	router.Get("/{code}", handler.RedirectToOriginal)

	return router
}
