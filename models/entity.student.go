package model

import (
	"time"
)

type Student struct {
	UserID            uint      `json:"user_id,omitempty" gorm:"default:null"`
	ID                uint      `gorm:"primaryKey"`
	ClassID           uint      `json:"class_id,omitempty" gorm:"default:null"`
	Class             Class     `gorm:"foreignKey:ClassID" json:"-"`
	Nim               string    `gorm:"type:varchar(25);unique"`
	Name              string    `gorm:"type:varchar(255)"`
	CompensationHours *uint     `json:"compensation_hours,omitempty" gorm:"-"`
	Address           string    `gorm:"type:varchar(255)"`
	Dob               time.Time `gorm:"type:date"`
	Year              uint
}
