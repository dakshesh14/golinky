package analytics

import (
	"time"

	"github.com/dakshesh14/golinky/internal/domain/model"
)

type LinkClickAnalysisResponse struct {
	TotalCount int64               `json:"total_count"`
	Data       []LinkClickResponse `json:"data"`
}

type LinkClickResponse struct {
	ID        int64  `json:"id"`
	LinkID    string `json:"link_id"`
	ClickedAt string `json:"clicked_at"`
	IPAddress string `json:"ip_address"`
	UserAgent string `json:"user_agent"`
}

func ToLinkClickResponse(lc *model.LinkClick) LinkClickResponse {
	return LinkClickResponse{
		ID:        lc.ID,
		LinkID:    lc.LinkID.String(),
		ClickedAt: lc.ClickedAt.Format(time.RFC3339),
		IPAddress: lc.IPAddress,
		UserAgent: lc.UserAgent,
	}
}
