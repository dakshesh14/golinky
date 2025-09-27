package analytics

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/dakshesh14/golinky/internal/container"
	"github.com/dakshesh14/golinky/internal/infrastructure/db/repository"
	"github.com/dakshesh14/golinky/pkg/response"
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

type Handler struct {
	container     container.Container
	linkRepo      repository.LinkRepositoryInterface
	linkClickRepo repository.LinkClickRepositoryInterface
}

func NewHandler(c container.Container, linkRepo repository.LinkRepositoryInterface, linkClickRepo repository.LinkClickRepositoryInterface) *Handler {
	return &Handler{
		container:     c,
		linkRepo:      linkRepo,
		linkClickRepo: linkClickRepo,
	}
}

func (h *Handler) GetLinkAnalysis(w http.ResponseWriter, r *http.Request) {
	linkCode := chi.URLParam(r, "linkCode")

	limit := int64(-1)
	offset := int64(-1)

	if v := r.URL.Query().Get("limit"); v != "" {
		if parsed, err := strconv.ParseInt(v, 10, 64); err == nil {
			limit = parsed
		}
	}
	if v := r.URL.Query().Get("offset"); v != "" {
		if parsed, err := strconv.ParseInt(v, 10, 64); err == nil {
			offset = parsed
		}
	}

	link, err := h.linkRepo.GetByCode(r.Context(), linkCode)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.WriteError(w, http.StatusNotFound, fmt.Sprintf("link with code %q not found", linkCode), nil)
			return
		}

		response.WriteError(w, http.StatusInternalServerError, "failed to get link", err.Error())
		return
	}

	totalCount, err := h.linkClickRepo.CountByLinkID(r.Context(), link.ID.String())
	if err != nil {
		response.WriteError(w, http.StatusInternalServerError, "failed to get analysis", err.Error())
		return
	}

	analysises, err := h.linkClickRepo.GetByLinkID(r.Context(), link.ID.String(), int(limit), int(offset))
	if err != nil {
		response.WriteError(w, http.StatusInternalServerError, "failed to get analysis", err.Error())
		return
	}

	analysisResponse := make([]LinkClickResponse, len(analysises))
	for i, analysis := range analysises {
		analysisResponse[i] = ToLinkClickResponse(&analysis)
	}

	resp := &LinkClickAnalysisResponse{
		TotalCount: totalCount,
		Data:       analysisResponse,
	}

	response.WriteJSON(w, http.StatusOK, resp)
}
