package link

import (
	"time"

	"github.com/dakshesh14/golinky/internal/domain/model"
)

type LinkResponse struct {
	ID        string `json:"id"`
	PK        uint   `json:"pk"`
	URL       string `json:"url"`
	Code      string `json:"code"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func ToLinkResponse(l *model.Link) LinkResponse {
	return LinkResponse{
		ID:        l.ID.String(),
		PK:        l.PK,
		URL:       l.URL,
		Code:      l.Code,
		CreatedAt: l.CreatedAt.Format(time.RFC3339),
		UpdatedAt: l.UpdatedAt.Format(time.RFC3339),
	}
}
