package model

type Course struct {
	ID        uint        `json:"id,omitempty" gorm:"primaryKey"`
	Code      string      `json:"code,omitempty" gorm:"unique"`
	Name      string      `json:"name,omitempty"`
	Hours     int         `json:"hours"`
	Classes   *[]Class    `json:"classes,omitempty" gorm:"many2many:enrollments"`
	Lecturers *[]Lecturer `json:"lecturers,omitempty" gorm:"many2many:enrollment_lecturers"`
}
