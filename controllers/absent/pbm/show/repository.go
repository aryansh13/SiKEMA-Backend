package showAbsent

import (
	model "attendance-is/models"

	"gorm.io/gorm"
)

type Repository interface {
	GetAbsentByID(id string) (*model.Absent, string)
}

type repository struct {
	db *gorm.DB
}

func NewShowAbsentRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) GetAbsentByID(id string) (*model.Absent, string) {
	var absent model.Absent
	if err := r.db.Model(absent).Preload("Student").Preload("Event").Preload("Event.Course").Preload("Event.Class").Preload("Event.Lecturer").Where("id = ?", id).Take(&absent).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, "ABSENT_NOTFOUND_404"
		} else {
			return nil, "ABSENT_UNEXPECTED_500 : " + err.Error()
		}
	}

	return &absent, ""
}
