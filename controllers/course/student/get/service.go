package getCourse

import (
	model "attendance-is/models"
)

type Service interface {
	GetCourse(input InputGetCourse) (*model.Enrollment, string)
	GetCourseService(input InputGetCourse) (*model.Course, string)
}

type service struct {
	repository Repository
}

func NewGetCourseService(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) GetCourse(input InputGetCourse) (*model.Enrollment, string) {
	student, err := s.repository.GetStudent(input.StudentID)
	if err != "" {
		return nil, err
	}

	course, err := s.repository.GetStudentCourse(student, input.CourseID)
	if err != "" {
		return nil, err
	}

	return course, err
}

func (s *service) GetCourseService(input InputGetCourse) (*model.Course, string) {
	res, err := s.repository.GetCourseRepository(input.StudentID, input.ClassID, input.CourseID)

	return res, err
}
