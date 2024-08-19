package removeStudentEvent

import (
	model "attendance-is/models"

	"gorm.io/gorm"
)

type Repository interface {
	GetStudentByWhere(query interface{}, args ...interface{}) (*[]model.Student, string)
	GetEventByID(ID string) (*model.Event, string)
	RemoveStudent(event *model.Event, students *[]model.Student) string
}

type repository struct {
	db *gorm.DB
}

func NewRemoveStudentEventRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) GetStudentByWhere(query interface{}, args ...interface{}) (*[]model.Student, string) {
	var student []model.Student
	err := r.db.Where(query, args...).Find(&student)
	if err.Error != nil {
		return nil, "STUDENT_UNEXPECTED_500 : " + err.Error.Error()
	}
	if err.RowsAffected == 0 {
		return nil, "STUDENT_NOTFOUND_404"
	}
	return &student, ""
}

func (r *repository) GetEventByID(ID string) (*model.Event, string) {
	var event model.Event
	if err := r.db.Where("id = ?", ID).Take(&event).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, "EVENT_NOTFOUND_404"
		}
		return nil, "EVENT_UNEXPECTED_500 : " + err.Error()
	}
	return &event, ""
}

func (r *repository) RemoveStudent(event *model.Event, students *[]model.Student) string {
	if err := r.db.Model(event).Association("Students").Delete(students); err != nil {
		return err.Error()
	}

	return ""
}
