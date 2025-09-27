package model

import (
	"time"

	"github.com/google/uuid"
)

type LinkClick struct {
	ID        int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	LinkID    uuid.UUID `json:"link_id" gorm:"type:uuid;not null;index"`
	ClickedAt time.Time `json:"clicked_at" gorm:"not null;default:CURRENT_TIMESTAMP"`
	IPAddress string    `json:"ip_address" gorm:"type:inet"`
	UserAgent string    `json:"user_agent" gorm:"type:text"`
}

type LinkClickEvent struct {
	LinkCode  string    `json:"linkCode"`
	IP        string    `json:"ip"`
	UserAgent string    `json:"userAgent"`
	TS        time.Time `json:"ts"`
}
