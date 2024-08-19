package model

import "gorm.io/gorm"

type Kompen struct {
	gorm.Model
	StudentID  uint    `json:"student_id,omitempty" gorm:"foreignKey:StudentID"`
	Student    Student `json:"student,omitempty"`
	Semester   uint    `json:"semester,omitempty"`
	AmountLeft uint    `json:"amount_left,omitempty"`
	Status     bool    `json:"status,omitempty"`
}
