package container

import (
	"github.com/dakshesh14/golinky/internal/config"
	"github.com/dakshesh14/golinky/internal/infrastructure/cache"
	"gorm.io/gorm"
)

type Container struct {
	Db     *gorm.DB
	Cache  cache.CacheService
	Config *config.Config
}

func NewContainer(db *gorm.DB, cache cache.CacheService, cfg *config.Config) *Container {
	return &Container{
		Db:     db,
		Cache:  cache,
		Config: cfg,
	}
}
