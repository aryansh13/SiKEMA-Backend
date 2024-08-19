package getAllExcuse

import (
	model "attendance-is/models"

	"gorm.io/gorm"
)

type Repository interface {
	GetAllExcuseRepository(student model.Student) (*[]model.Excuse, string)
}

type repository struct {
	db *gorm.DB
}

func NewGetAllExcuseRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) GetAllExcuseRepository(student model.Student) (*[]model.Excuse, string) {
	var excuses []model.Excuse
	if err := r.db.Model(excuses).Preload("Absent", "id = ?", student.ID).Find(&excuses).Error; err != nil {
		return nil, "EXCUSE_UNEXPECTED_500 : " + err.Error()
	}

	if len(excuses) <= 0 {
		return nil, "EXCUSE_NOTFOUND_404"
	}
	return &excuses, ""
}
