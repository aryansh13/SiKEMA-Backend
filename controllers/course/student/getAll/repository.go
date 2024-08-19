package getAllCourse

import (
	model "attendance-is/models"
	"errors"

	"gorm.io/gorm"
)

type Repository interface {
	GetStudent(StudentID string) (*model.Student, string)
	GetStudentCourse(student *model.Student) (*[]model.Enrollment, string)
	GetAllCourseRepository(StudentID string, ClassID string) (*[]model.Enrollment, string)
}

type repository struct {
	db *gorm.DB
}

func NewGetAllCourseRepository(db *gorm.DB) *repository {
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

func (r *repository) GetStudentCourse(student *model.Student) (*[]model.Enrollment, string) {
	var courses []model.Enrollment
	if err := r.db.Model(model.Enrollment{}).Preload("Course").Preload("Lecturers").Where("class_id = ?", student.ClassID).Find(&courses).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, "COURSE_NOTFOUND_404"
		}
		return nil, "COURSE_UNEXPECTED_500 : " + err.Error()
	}

	return &courses, ""
}

func (r *repository) GetAllCourseRepository(StudentID string, ClassID string) (*[]model.Enrollment, string) {
	// var class model.Class
	// r.db.Where("id = ?", ClassID).Find(&class)

	// var courses []model.Course
	// if err := r.db.Model(&class).Association("Courses").Find(&courses); err != nil {
	// 	return nil, "500_INTERNAL"
	// }

	// if len(courses) == 0 {
	// 	return nil, "404_NOTFOUND"
	// }

	var courses []model.Enrollment
	r.db.Model(model.Enrollment{}).Preload("Course").Preload("Lecturers").Where("class_id = ?", ClassID).Find(&courses)

	return &courses, ""
}
