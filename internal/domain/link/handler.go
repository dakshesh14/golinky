package link

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"

	"github.com/dakshesh14/golinky/internal/container"
	"github.com/dakshesh14/golinky/internal/domain/model"
	"github.com/dakshesh14/golinky/internal/infrastructure/db/repository"
	"github.com/dakshesh14/golinky/pkg/response"
	"github.com/dakshesh14/golinky/pkg/shorten"
)

const IDCounter = 99999

type Handler struct {
	container      container.Container
	linkRepository repository.LinkRepositoryInterface
}

func NewHandler(c container.Container, repo repository.LinkRepositoryInterface) *Handler {
	return &Handler{
		container:      c,
		linkRepository: repo,
	}
}

func (h *Handler) CreateShortURL(w http.ResponseWriter, r *http.Request) {
	var req ShortenURLRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.WriteError(w, http.StatusBadRequest, "invalid request body", map[string]any{
			"error": err.Error(),
		})
		return
	}

	lastLink, err := h.linkRepository.GetLast(r.Context())
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		response.WriteError(w, http.StatusInternalServerError, "failed to fetch last link", map[string]any{
			"error": err.Error(),
		})
		return
	}

	idToEncode := IDCounter
	if lastLink != nil {
		idToEncode = int(lastLink.PK)
	}

	shortCode := shorten.EncodeBase62(idToEncode)

	link := &model.Link{
		URL:  req.URL,
		Code: shortCode,
	}

	if err := h.linkRepository.Create(r.Context(), link); err != nil {
		response.WriteError(w, http.StatusInternalServerError, "failed to create link", map[string]any{
			"error": err.Error(),
		})
		return
	}

	response.WriteJSON(w, http.StatusOK, ToLinkResponse(link))
}

func (h *Handler) RedirectToOriginal(w http.ResponseWriter, r *http.Request) {
	code := chi.URLParam(r, "code")

	if cacheLink, err := h.container.Cache.Get(r.Context(), code); err == nil {
		http.Redirect(w, r, cacheLink, http.StatusFound)
		return
	}

	link, err := h.linkRepository.GetByCode(r.Context(), code)
	switch {
	case err == nil:
		h.container.Cache.Set(r.Context(), code, link.URL, time.Minute)
		http.Redirect(w, r, link.URL, http.StatusFound)

	case errors.Is(err, gorm.ErrRecordNotFound):
		response.WriteError(w, http.StatusNotFound,
			fmt.Sprintf("link with code %q not found", code), nil)

	default:
		response.WriteError(w, http.StatusInternalServerError,
			"failed to fetch link", map[string]any{"error": err.Error()})
	}
}
