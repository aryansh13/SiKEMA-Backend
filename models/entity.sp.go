package model

type SP struct {
	ID        uint `gorm:"primaryKey"`
	StudentID uint
	Student   Student `gorm:"foreignKey:StudentID"`
	Semester  uint
	Level     uint
	Status    uint
}
