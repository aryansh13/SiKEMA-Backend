package showAbsent

import (
	model "attendance-is/models"

	"gorm.io/gorm"
)

type Repository interface {
	GetStudentByID(ID string) (*model.Student, string)
	GetAbsentByWhere(query interface{}, args ...interface{}) (*model.Absent, string)
}

type repository struct {
	db *gorm.DB
}

func NewShowAbsentRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) GetStudentByID(ID string) (*model.Student, string) {
	var student model.Student
	if err := r.db.Where("id = ?", ID).Take(&student).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, "STUDENT_NOTFOUND_404"
		}
	}
	return &student, ""
}

func (r *repository) GetAbsentByWhere(query interface{}, args ...interface{}) (*model.Absent, string) {
	var absent model.Absent
	if err := r.db.Model(absent).Preload("Event").Preload("Event.Course").Preload("Event.Lecturer").Where(query, args...).Take(&absent).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, "ABSENT_NOTFOUND_404"
		}
	}
	return &absent, ""
}
