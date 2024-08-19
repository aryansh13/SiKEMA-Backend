package createSP

import (
	model "attendance-is/models"
	"errors"

	"gorm.io/gorm"
)

type Repository interface {
	CreateSPRepository(input *model.SP) (*model.SP, string)
}

type repository struct {
	db *gorm.DB
}

func NewCreateSPRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) CreateSPRepository(input *model.SP) (*model.SP, string) {
	if err := r.db.Create(&input).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, "ERR_CONFLICT_409"
		}
		return nil, "ERR_UNEXPECTED_500"
	}

	return input, ""
}
