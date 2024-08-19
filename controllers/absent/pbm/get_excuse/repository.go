package getAbsentExcuse

import (
	model "attendance-is/models"

	"gorm.io/gorm"
)

type Repository interface {
	GetAbsentByID(id string) (*model.Absent, string)
	GetAbsentExcuse(absent *model.Absent) (*model.Excuse, string)
}

type repository struct {
	db *gorm.DB
}

func NewGetAbsentExcuseRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) GetAbsentByID(id string) (*model.Absent, string) {
	var absent model.Absent
	if err := r.db.Where("id = ?", id).Take(&absent).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, "ABSENT_NOTFOUND_404"
		} else {
			return nil, "ABSENT_UNEXPECTED_500 : " + err.Error()
		}
	}

	return &absent, ""
}

func (r *repository) GetAbsentExcuse(absent *model.Absent) (*model.Excuse, string) {
	var excuse model.Excuse
	if err := r.db.Where("absent_id = ?", absent.ID).Take(&excuse).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, "EXCUSE_NOTFOUND_404"
		} else {
			return nil, "EXCUSE_UNEXPECTED_500 : " + err.Error()
		}
	}

	return &excuse, ""
}
