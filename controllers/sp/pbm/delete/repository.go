package deleteSP

import (
	model "attendance-is/models"

	"gorm.io/gorm"
)

type Repository interface {
	DeleteSPRepository(input *model.SP) string
}

type repository struct {
	db *gorm.DB
}

func NewDeleteSPRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) DeleteSPRepository(input *model.SP) string {
	var sp model.SP
	db := r.db.Model(input)

	res := db.Where("id = ?", input.ID).Find(&sp)
	if res.RowsAffected == 0 {
		return "ERR_NOTFOUND_404"
	}

	res = db.Delete(&sp)
	if res.Error != nil {
		return "ERR_UNEXPECTED_500"
	}

	return ""
}
