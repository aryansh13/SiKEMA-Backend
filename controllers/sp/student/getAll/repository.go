package getAllSP

import (
	model "attendance-is/models"

	"gorm.io/gorm"
)

type Repository interface {
	GetAllSPRepository(StudentID uint) (*[]model.SP, string)
}

type repository struct {
	db *gorm.DB
}

func NewGetAllSPRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) GetAllSPRepository(StudentID uint) (*[]model.SP, string) {
	var sp []model.SP
	result := r.db.Where("student_id = ?", StudentID).Find(&sp)

	if result.RowsAffected == 0 {
		return nil, "ERR_NOTFOUND_404"
	}

	if result.Error != nil {
		return nil, "ERR_UNEXPECTED_500"
	}

	return &sp, ""
}
