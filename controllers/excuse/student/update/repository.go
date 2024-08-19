package updateExcuse

import (
	model "attendance-is/models"

	"gorm.io/gorm"
)

type Repository interface {
	GetExcuseByID(id string) (*model.Excuse, string)
	SaveExcuse(excuse *model.Excuse) string
}

type repository struct {
	db *gorm.DB
}

func NewUpdateExcuseRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) GetExcuseByID(id string) (*model.Excuse, string) {
	var excuse model.Excuse
	if err := r.db.Where("id = ?", id).Take(&excuse).Error; err != nil {
		return nil, "EXCUSE_NOTFOUND_404"
	}
	return &excuse, ""
}

func (r *repository) SaveExcuse(excuse *model.Excuse) string {

	if err := r.db.Save(&excuse).Error; err != nil {
		return "EXCUSE_UNEXPECTED_500 : " + err.Error()
	}
	return ""
}
