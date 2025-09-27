package repository

import (
	"context"

	"gorm.io/gorm"

	"github.com/dakshesh14/golinky/internal/domain/model"
)

type LinkClickRepositoryInterface interface {
	Create(ctx context.Context, click *model.LinkClick) error
	CountByLinkID(ctx context.Context, linkID string) (int64, error)
	GetByLinkID(ctx context.Context, linkID string, limit int, offset int) ([]model.LinkClick, error)
}

type LinkClickRepository struct {
	db *gorm.DB
}

func NewLinkClickRepository(db *gorm.DB) LinkClickRepositoryInterface {
	return &LinkClickRepository{db: db}
}

func (r *LinkClickRepository) Create(ctx context.Context, click *model.LinkClick) error {
	return r.db.WithContext(ctx).Create(click).Error
}

func (r *LinkClickRepository) CountByLinkID(ctx context.Context, linkID string) (int64, error) {
	var count int64
	if err := r.db.WithContext(ctx).
		Model(&model.LinkClick{}).
		Where("link_id = ?", linkID).
		Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (r *LinkClickRepository) GetByLinkID(ctx context.Context, linkID string, limit int, offset int) ([]model.LinkClick, error) {
	var clicks []model.LinkClick
	if err := r.db.WithContext(ctx).
		Where("link_id = ?", linkID).
		Order("clicked_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&clicks).Error; err != nil {
		return nil, err
	}
	return clicks, nil
}
