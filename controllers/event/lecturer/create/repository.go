package createEvent

import (
	model "attendance-is/models"

	"gorm.io/gorm"
)

type Repository interface {
	GetLecturerByID(ID string) (*model.Lecturer, string)
	GetEnrollmentByWhere(lecturer *model.Lecturer, courseId string, classId string) (*model.Enrollment, string)
	CreateEvent(event *model.Event) string
}

type repository struct {
	db *gorm.DB
}

func NewCreateEventRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) GetLecturerByID(ID string) (*model.Lecturer, string) {
	var lecturer model.Lecturer
	if err := r.db.Where("id = ?", ID).Take(&lecturer).Error; err != nil {
		return nil, err.Error()
	}
	return &lecturer, ""
}

func (r *repository) GetEnrollmentByWhere(lecturer *model.Lecturer, courseId string, classId string) (*model.Enrollment, string) {
	var courses []model.Enrollment
	if err := r.db.Model(lecturer).Preload("Class").Preload("Course").Where("course_id = ? AND class_id = ?", courseId, classId).Association("Enrollments").Find(&courses); err != nil {

		return nil, err.Error()
	}
	if len(courses) < 1 {
		return nil, "COURSE_NOTFOUND_404"
	}
	return &courses[0], ""

}

func (r *repository) CreateEvent(event *model.Event) string {
	if err := r.db.Create(&event).Error; err != nil {
		if err == gorm.ErrDuplicatedKey {
			return "EVENT_CONFLICT_409"
		}
		return "EVENT_UNEXPECTED_500 : " + err.Error()
	}
	return ""
}
