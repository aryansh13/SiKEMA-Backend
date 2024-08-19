package model

import "gorm.io/gorm"

type Class struct {
	ID           uint      `json:"id,omitempty" gorm:"primaryKey"`
	Name         string    `json:"name,omitempty" gorm:"type:varchar(25);unique"`
	StudentCount uint      `json:"student_count,omitempty" gorm:"-"`
	Students     []Student `json:"students,omitempty" gorm:"foreignKey:ClassID"`
	Courses      []Course  `json:"courses,omitempty" gorm:"many2many:enrollments"`
}

func (entity *Class) AfterFind(tx *gorm.DB) (err error) {
	if entity.StudentCount == 0 {
		entity.StudentCount = uint(tx.Model(entity).Association("Students").Count())
	}
	return
}
