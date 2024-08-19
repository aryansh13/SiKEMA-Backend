package updateSP

import (
	model "attendance-is/models"
	"errors"

	"gorm.io/gorm"
)

type Repository interface {
	UpdateSPRepository(input *model.SP) (*model.SP, string)
}

type repository struct {
	db *gorm.DB
}

func NewUpdateSPRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) UpdateSPRepository(input *model.SP) (*model.SP, string) {
	if err := r.db.Save(&input).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, "ERR_CONFLICT_409"
		}
		return nil, "ERR_UNEXPECTED_500"
	}

	return input, ""
}
