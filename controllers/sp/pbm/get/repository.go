package getSP

import (
	model "attendance-is/models"

	"gorm.io/gorm"
)

type Repository interface {
	GetSPRepository(ID uint) (*model.SP, string)
}

type repository struct {
	db *gorm.DB
}

func NewGetSPRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) GetSPRepository(ID uint) (*model.SP, string) {
	var sp model.SP
	result := r.db.Where("id = ?", ID).First(&sp)

	if result.RowsAffected == 0 {
		return nil, "ERR_NOTFOUND_404"
	}

	if result.Error != nil {
		return nil, "ERR_UNEXPECTED_500"
	}

	return &sp, ""
}
