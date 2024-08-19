package model

import "time"

type PBM struct {
	UserID  uint   `json:"user_id,omitempty" gorm:"default:null"`
	ID      uint   `json:"id" gorm:"primaryKey"`
	Nip     string `json:"nip" gorm:"type:varchar(50);index"`
	Name    string `json:"name" gorm:"type:varchar(100)"`
	Address string `json:"address,omitempty" gorm:"type:varchar(255)"`
	Phone   string `json:"phone" gorm:"type:varchar(25)"`
	Dob     time.Time
}
