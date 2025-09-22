package repository

import (
	"context"

	"gorm.io/gorm"

	"github.com/dakshesh14/golinky/internal/domain/model"
)

type LinkRepositoryInterface interface {
	Create(ctx context.Context, link *model.Link) error
	Get(ctx context.Context, id string) (*model.Link, error)
	GetLast(ctx context.Context) (*model.Link, error)
	GetByCode(ctx context.Context, code string) (*model.Link, error)
}

type LinkRepository struct {
	db *gorm.DB
}

func NewLinkRepository(db *gorm.DB) LinkRepositoryInterface {
	return &LinkRepository{db: db}
}

func (r *LinkRepository) Create(ctx context.Context, link *model.Link) error {
	return r.db.WithContext(ctx).Create(link).Error
}

func (r *LinkRepository) Get(ctx context.Context, id string) (*model.Link, error) {
	var link model.Link
	if err := r.db.WithContext(ctx).First(&link, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &link, nil
}

func (r *LinkRepository) GetLast(ctx context.Context) (*model.Link, error) {
	var link model.Link
	if err := r.db.WithContext(ctx).
		Order("pk DESC").
		First(&link).Error; err != nil {
		return nil, err
	}

	return &link, nil
}

func (r *LinkRepository) GetByCode(ctx context.Context, code string) (*model.Link, error) {
	var link model.Link
	if err := r.db.WithContext(ctx).
		Where("code = ?", code).
		First(&link).Error; err != nil {
		return nil, err
	}

	return &link, nil
}
