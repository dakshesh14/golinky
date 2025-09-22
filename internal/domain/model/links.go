package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Link struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	PK        uint      `json:"pk" gorm:"autoIncrement;not null"`
	URL       string    `json:"url" gorm:"type:text;not null"`
	Code      string    `json:"code" gorm:"type:varchar(32);not null"`
	CreatedAt time.Time `json:"created_at" gorm:"not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" gorm:"not null;default:CURRENT_TIMESTAMP"`
}

func (Link) TableName() string {
	return "links"
}

func (u *Link) BeforeUpdate(tx *gorm.DB) error {
	u.UpdatedAt = time.Now()
	return nil
}
