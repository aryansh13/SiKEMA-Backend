package getEvent

import (
	model "attendance-is/models"

	"gorm.io/gorm"
)

type Repository interface {
	GetEvent(lecturerId string, classId string, courseId string) (*[]model.Event, string)
}

type repository struct {
	db *gorm.DB
}

func NewGetEventRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) GetEvent(lecturerId string, classId string, courseId string) (*[]model.Event, string) {
	var event []model.Event
	db := r.db.Model(&event).Preload("Class").Preload("Course").Preload("Students").Preload("Lecturer").Where("lecturer_id = ?", lecturerId)

	if classId != "" {
		db = db.Where("class_id = ?", classId)
	}

	if courseId != "" {
		db = db.Where("course_id = ?", courseId)
	}

	if err := db.Find(&event).Error; err != nil {
		return nil, "EVENT_UNEXPECTED_500 : " + err.Error()
	}

	if len(event) < 1 {
		return nil, "EVENT_NOTFOUND_404"
	}
	return &event, ""
}
