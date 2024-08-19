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
	db := r.db.Model(model.Excuse{})

	var count int64
	if err := db.Where("id = ?", input.ID).Count(&count).Find(&excuse).Error; err != nil {
		return nil, "EXCUSE_UNEXPECTED_500 : " + err.Error()
	}

	if count == 0 {
		return nil, "EXCUSE_NOTFOUND_404"
	}

	return &excuse, ""
}
