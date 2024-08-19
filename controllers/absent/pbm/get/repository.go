package getAbsent

import (
	model "attendance-is/models"

	"gorm.io/gorm"
)

type Repository interface {
	GetAbsent() (*[]model.Absent, string)
}

type repository struct {
	db *gorm.DB
}

func NewGetAbsentRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) GetAbsent() (*[]model.Absent, string) {
	var absent []model.Absent
	if err := r.db.Model(model.Absent{}).Preload("Student").Preload("Event").Preload("Event.Course").Preload("Event.Class").Preload("Event.Lecturer").Find(&absent).Error; err != nil {
		return nil, "ABSENT_UNEXPECTED_500 : " + err.Error()
	}

	if len(absent) < 1 {
		return nil, "ABSENT_NOTFOUND_404"
	}

	return &absent, ""
}
