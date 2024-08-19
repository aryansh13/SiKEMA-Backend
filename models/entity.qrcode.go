package model

import (
	"time"

	"gorm.io/gorm"
)

type QRCode struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	EventID   uint      `json:"event_id,omitempty" gorm:"primaryKey"`
	Event     Event     `json:"event,omitempty" gorm:"foreignKey:EventID"`
	SSID      string    `json:"ssid"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

func (entity *QRCode) BeforeCreate(db *gorm.DB) error {
	entity.CreatedAt = time.Now().Local()
	return nil
}
