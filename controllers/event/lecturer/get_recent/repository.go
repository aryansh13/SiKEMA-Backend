package getRecentEvent

import (
	model "attendance-is/models"
	"errors"

	"gorm.io/gorm"
)

type Repository interface {
	GetRecentEvent(LecturerID string) (*model.Event, string)
}

type repository struct {
	db *gorm.DB
}

func NewGetRecentEventRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) GetRecentEvent(LecturerID string) (*model.Event, string) {
	var event model.Event
	if err := r.db.Model(&event).Preload("Class").Preload("Course").Preload("Students").Preload("Lecturer").Where("lecturer_id = ?", LecturerID).Where("status != ?", 2).Order("created_at desc").First(&event).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, "EVENT_NOTFOUND_404"
		}
		return nil, err.Error()
	}

	return &event, ""
}
