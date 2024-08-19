package compensationGet

import (
	model "attendance-is/models"

	"gorm.io/gorm"
)

type Repository interface {
	GetStudent() (*[]model.Student, string)
	GetAbsent(student *model.Student) (*[]model.Absent, string)
	GetExcuse(absent *model.Absent) (*model.Excuse, string)
}

type repository struct {
	db *gorm.DB
}

func NewCompensationGetRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) GetStudent() (*[]model.Student, string) {
	var students []model.Student
	if err := r.db.Find(&students).Error; err != nil {
		return nil, "STUDENT_UNEXPECTED_500 : " + err.Error()
	}

	return &students, ""
}

func (r *repository) GetAbsent(student *model.Student) (*[]model.Absent, string) {
	var absents []model.Absent
	if err := r.db.Where("student_id = ?", student.ID).Find(&absents).Error; err != nil {
		return nil, "ABSENT_UNEXPECTED_500 : " + err.Error()
	}

	return &absents, ""
}

func (r *repository) GetExcuse(absent *model.Absent) (*model.Excuse, string) {
	var excuse model.Excuse
	if err := r.db.Where("absent_id = ?", absent.ID).Find(&excuse).Error; err != nil {
		return nil, "EXCUSE_UNEXPECTED_500 : " + err.Error()
	}

	return &excuse, ""
}
