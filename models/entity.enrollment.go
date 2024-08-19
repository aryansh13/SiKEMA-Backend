package model

type Enrollment struct {
	CourseID  uint        `json:"course_id,omitempty" gorm:"primaryKey"`
	ClassID   uint        `json:"class_id,omitempty" gorm:"primaryKey"`
	Course    Course      `json:"course,omitempty" gorm:"foreignKey:CourseID"`
	Class     Class       `json:"class,omitempty" gorm:"foreignKey:ClassID"`
	Lecturers *[]Lecturer `json:"lecturers,omitempty" gorm:"many2many:enrollment_lecturers"`
}
