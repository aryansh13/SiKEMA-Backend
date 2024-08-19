package getRecentEvent

import (
	model "attendance-is/models"
	"errors"

	"gorm.io/gorm"
)

type Repository interface {
	GetStudent(StudentID string) (*model.Student, string)
	GetRecentEvent(classId string) (*model.Event, string)
}

type repository struct {
	db *gorm.DB
}

func NewGetRecentEventRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) GetStudent(StudentID string) (*model.Student, string) {
	var student model.Student
	if err := r.db.Where("id = ?", StudentID).Take(&student).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, "STUDENT_NOTFOUND_404"
		}
		return nil, "STUDENT_UNEXPECTED_500 : " + err.Error()
	}
	return &student, ""
}

func (r *repository) GetRecentEvent(classId string) (*model.Event, string) {
	var event model.Event
	if err := r.db.Model(&event).Preload("Class").Preload("Course").Preload("Students").Preload("Lecturer").Where("class_id = ?", classId).Where("status != ?", 2).Order("created_at desc").First(&event).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, "EVENT_NOTFOUND_404"
		}
		return nil, err.Error()
	}

	return &event, ""
}
