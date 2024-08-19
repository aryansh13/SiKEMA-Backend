package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string    `json:"email,omitempty" gorm:"type:varchar(255)"`
	Password string    `json:"-" gorm:"type:varchar(255)"`
	Type     uint      `json:"type,omitempty"`
	Student  *Student  `json:"student,omitempty" gorm:"foreignKey:UserID"`
	Lecturer *Lecturer `json:"lecturer,omitempty" gorm:"foreignKey:UserID"`
	PBM      *PBM      `json:"pbm,omitempty" gorm:"foreignKey:UserID"`
}
