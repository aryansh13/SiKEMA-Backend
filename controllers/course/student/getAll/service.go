package getAllCourse

import (
	model "attendance-is/models"
)

type Service interface {
	GetAllCourse(input InputGetAllCourse) (*[]model.Enrollment, string)
	GetAllCourseService(input InputGetAllCourse) (*[]model.Enrollment, string)
}

type service struct {
	repository Repository
}

func NewGetAllCourseService(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) GetAllCourse(input InputGetAllCourse) (*[]model.Enrollment, string) {
	student, err := s.repository.GetStudent(input.StudentID)
	if err != "" {
		return nil, err
	}

	courses, err2 := s.repository.GetStudentCourse(student)
	if err2 != "" {
		return nil, err2
	}

	return courses, err
}

func (s *service) GetAllCourseService(input InputGetAllCourse) (*[]model.Enrollment, string) {
	res, err := s.repository.GetAllCourseRepository(input.StudentID, input.ClassID)

	return res, err
}
