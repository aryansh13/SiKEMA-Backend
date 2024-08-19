package getCourse

import (
	model "attendance-is/models"

	"gorm.io/gorm"
)

type Repository interface {
	GetEnrollment() (*[]model.Enrollment, string)
}

type repository struct {
	db *gorm.DB
}

func NewGetCourseRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) GetEnrollment() (*[]model.Enrollment, string) {
	var enrollments []model.Enrollment
	if err := r.db.Model(&enrollments).Preload("Class").Preload("Course").Preload("Lecturers").Find(&enrollments).Error; err != nil {
		return nil, "COURSE_UNEXPECTED_500 : " + err.Error()
	}

	if len(enrollments) < 1 {
		return nil, "COURSE_NOTFOUND_404"
	}

	return &enrollments, ""
}
