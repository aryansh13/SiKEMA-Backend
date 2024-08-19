package getCourse

import (
	model "attendance-is/models"

	"gorm.io/gorm"
)

type Repository interface {
	GetLecturerByID(id string) (*model.Lecturer, string)
	GetEnrollmentByLecturer(lecturer *model.Lecturer, classId string, courseId string) (*[]model.Enrollment, string)
}

type repository struct {
	db *gorm.DB
}

func NewGetCourseRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) GetLecturerByID(id string) (*model.Lecturer, string) {
	var lecturer model.Lecturer

	if err := r.db.Where("id = ?", id).Take(&lecturer).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, "LECTURER_NOTFOUND_404"
		}
		return nil, "LECTURER_UNEXPECTED_500" + err.Error()
	}

	return &lecturer, ""
}

func (r *repository) GetEnrollmentByLecturer(lecturer *model.Lecturer, classId string, courseId string) (*[]model.Enrollment, string) {
	var enrollments []model.Enrollment

	db := r.db.Model(&lecturer).Preload("Class").Preload("Course")
	if classId != "" {
		db = db.Where("class_id = ?", classId)
	}

	if courseId != "" {
		db = db.Where("course_id = ?", courseId)
	}

	if err := db.Association("Enrollments").Find(&enrollments); err != nil {
		return nil, "COURSE_UNEXPECTED_500 : " + err.Error()
	}

	if len(enrollments) < 1 {
		return nil, "COURSE_NOTFOUND_404"
	}

	return &enrollments, ""
}
