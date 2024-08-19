package validateAuth

import (
	model "attendance-is/models"

	"gorm.io/gorm"
)

type Repository interface {
	GetUserByID(id string) (*model.User, string)
}

type repository struct {
	db *gorm.DB
}

func NewValidateAuthRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) GetUserByID(id string) (*model.User, string) {
	var user model.User
	if err := r.db.Model(user).Preload("Student").Preload("Lecturer").Preload("PBM").Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err.Error()
	}

	return &user, ""
}
