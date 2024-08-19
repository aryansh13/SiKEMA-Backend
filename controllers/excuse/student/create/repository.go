package createExcuse

import (
	model "attendance-is/models"

	"gorm.io/gorm"
)

type Repository interface {
	CreateExcuseRepository(input *model.Excuse) (*model.Excuse, string)
}

type repository struct {
	db *gorm.DB
}

func NewCreateExcuseRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) CreateExcuseRepository(input *model.Excuse) (*model.Excuse, string) {

	db := r.db.Model(&model.Excuse{})

	var count int64
	db.Where("id = ?", input.AbsentID).Count(&count)
	if count != 0 {
		return nil, "EXCUSE_CONFLICT_409"
	}

	if err := db.Create(&input).Error; err != nil {
		return nil, "EXCUSE_UNEXPECTED_500"
	}
	return input, ""
}
