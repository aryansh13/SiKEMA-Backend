package getExcuse

import (
	model "attendance-is/models"

	"gorm.io/gorm"
)

type Repository interface {
	GetExcuseRepository(input InputGetExcuse) (*model.Excuse, string)
}

type repository struct {
	db *gorm.DB
}

func NewGetExcuseRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) GetExcuseRepository(input InputGetExcuse) (*model.Excuse, string) {
	var excuse model.Excuse
	if err := r.db.Model(excuse).Preload("Absent", "student_id = ?", input.StudentID).Where("id = ?", input.ID).Take(&excuse).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, "EXCUSE_NOTFOUND_404"
		}
		return nil, "EXCUSE_UNEXPECTED_500"
	}
	return &excuse, ""
}
