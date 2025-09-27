package analytics

import (
	"context"
	"encoding/json"
	"log"

	"github.com/dakshesh14/golinky/internal/container"
	"github.com/dakshesh14/golinky/internal/domain/model"
	"github.com/dakshesh14/golinky/internal/infrastructure/cache"
	"github.com/dakshesh14/golinky/internal/infrastructure/db/repository"
	"github.com/dakshesh14/golinky/pkg/constants"
)

type AnalyticsConsumer struct {
	cache         cache.CacheService
	linkRepo      repository.LinkRepositoryInterface
	linkClickRepo repository.LinkClickRepositoryInterface
}

func NewAnalyticsConsumer(container *container.Container) *AnalyticsConsumer {
	linkRepo := repository.NewLinkRepository(container.Db)
	linkClickRepo := repository.NewLinkClickRepository(container.Db)

	return &AnalyticsConsumer{cache: container.Cache, linkRepo: linkRepo, linkClickRepo: linkClickRepo}
}

func (c *AnalyticsConsumer) Start(ctx context.Context) {
	log.Println("Analytics consumer started...")
	for {
		select {
		case <-ctx.Done():
			log.Println("Stopping analytics consumer...")
			return
		default:
			raw, err := c.cache.Dequeue(ctx, constants.AnalyticsQueue, 0)
			if err != nil {
				if ctx.Err() != nil {
					return
				}
				log.Printf("Failed to pop analytics event: %v", err)
				continue
			}

			var event model.LinkClickEvent
			if err := json.Unmarshal([]byte(raw), &event); err != nil {
				log.Printf("Invalid analytics event: %v", err)
				continue
			}

			link, err := c.linkRepo.GetByCode(ctx, event.LinkCode)
			if err != nil {
				log.Printf("Failed to get link: %v", err)
				continue
			}

			click := model.LinkClick{
				LinkID:    link.ID,
				IPAddress: event.IP,
				UserAgent: event.UserAgent,
				ClickedAt: event.TS,
			}

			if err := c.linkClickRepo.Create(ctx, &click); err != nil {
				log.Printf("Failed to save click: %v", err)
			}
		}
	}
}
