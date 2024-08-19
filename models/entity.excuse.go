package model

import (
	"time"

	"gorm.io/gorm"
)

type Excuse struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	AbsentID   uint      `json:"absent_id"`
	Absent     *Absent   `json:"absent,omitempty" gorm:"foreignKey:AbsentID"`
	Excuse     string    `json:"excuse" gorm:"type:varchar(255)"`
	Attachment string    `json:"attachment" gorm:"type:varchar(255)"`
	Status     uint      `json:"status"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
	UpdatedAt  time.Time `json:"updated_at,omitempty"`
}

func (entity *Excuse) BeforeCreate(db *gorm.DB) error {
	entity.CreatedAt = time.Now().Local()
	return nil
}

func (entity *Excuse) BeforeUpdate(db *gorm.DB) error {
	entity.UpdatedAt = time.Now().Local()
	return nil
}
