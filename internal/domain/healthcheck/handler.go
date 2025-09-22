package healthcheck

import (
	"net/http"
	"time"

	"github.com/google/uuid"

	"github.com/dakshesh14/golinky/internal/container"
	"github.com/dakshesh14/golinky/pkg/response"
)

type Handler struct {
	Container container.Container
}

func NewHandler(container container.Container) *Handler {
	return &Handler{
		Container: container,
	}
}

func (h *Handler) HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	sqlDB, err := h.Container.Db.DB()
	dbStatus := "error"
	if err == nil && sqlDB.Ping() == nil {
		dbStatus = "ok"
	}

	testingUUID := uuid.New().String()
	err = h.Container.Cache.Set(r.Context(), testingUUID, testingUUID, time.Duration(time.Millisecond))
	cacheStatus := "error"
	if err == nil {
		cacheStatus = "ok"
	}
	h.Container.Cache.Delete(r.Context(), testingUUID)

	resp := &HealthCheckResponse{
		DB:    dbStatus,
		Cache: cacheStatus,
	}

	response.WriteJSON(w, http.StatusOK, resp)
}
