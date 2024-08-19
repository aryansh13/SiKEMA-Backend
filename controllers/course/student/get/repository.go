package getCourse

import (
	model "attendance-is/models"
	"errors"

	"gorm.io/gorm"
)

type Repository interface {
	GetStudent(StudentID string) (*model.Student, string)
	GetStudentCourse(student *model.Student, ID string) (*model.Enrollment, string)
	GetCourseRepository(StudentID string, ClassID string, CourseID string) (*model.Course, string)
}

type repository struct {
	db *gorm.DB
}

func NewGetCourseRepository(db *gorm.DB) *repository {
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

func (r *repository) GetStudentCourse(student *model.Student, ID string) (*model.Enrollment, string) {
	var courses model.Enrollment
	if err := r.db.Model(model.Enrollment{}).Preload("Course").Preload("Lecturers").Where("class_id = ? AND course_id = ?", student.ClassID, ID).Take(&courses).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, "COURSE_NOTFOUND_404"
		}
		return nil, "COURSE_UNEXPECTED_500 : " + err.Error()
	}

	return &courses, ""
}

func (r *repository) GetCourseRepository(StudentID string, ClassID string, CourseID string) (*model.Course, string) {
	var class model.Class
	r.db.Where("id = ?", ClassID).Find(&class)

	var courses model.Course
	if err := r.db.Model(&class).Where("id = ?", CourseID).Association("Courses").Find(&courses); err != nil {
		return nil, "500_INTERNAL"
	}

	if courses.ID == 0 {
		return nil, "404_NOTFOUND"
	}

	return &courses, ""
}
