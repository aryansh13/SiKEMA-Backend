package getCourseAttendance

import (
	model "attendance-is/models"
	"errors"

	"gorm.io/gorm"
)

type Repository interface {
	GetStudent(StudentID string) (*model.Student, string)
	GetEnrollmentByID(student *model.Student, id string) (*model.Enrollment, string)
	GetCourseEvent(course *model.Course, student *model.Student) (*[]model.Event, string)
}

type repository struct {
	db *gorm.DB
}

func NewGetCourseAttendanceRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) GetStudent(StudentID string) (*model.Student, string) {
	var student model.Student
	if err := r.db.Where("id = ?", StudentID).Take(&student).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, "STUDENT_NOTFOUND_404"
		}
		return nil, "STUDENT_UNEXPECTED_500 : " + err.Error()
	}
	return &student, ""
}

func (r *repository) GetEnrollmentByID(student *model.Student, id string) (*model.Enrollment, string) {
	var courses model.Enrollment
	if err := r.db.Model(model.Enrollment{}).Preload("Class").Preload("Course").Where("class_id = ? AND course_id = ?", student.ClassID, id).Find(&courses).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, "COURSE_NOTFOUND_404"
		}
		return nil, "COURSE_UNEXPECTED_500 : " + err.Error()
	}

	return &courses, ""
}

func (r *repository) GetCourseEvent(course *model.Course, student *model.Student) (*[]model.Event, string) {
	var events []model.Event
	r.db.Model(events).Preload("Students", "id = ?", student.ID).Where("course_id = ? AND class_id = ? AND status = ?", course.ID, student.ClassID, 2).Find(&events)
	return &events, ""
}
