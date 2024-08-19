package getCourse

import (
	model "attendance-is/models"
)

type Service interface {
	GetCourseService() (*[]model.Enrollment, string)
}

type service struct {
	repository Repository
}

func NewGetCourseService(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) GetCourseService() (*[]model.Enrollment, string) {
	enrollments, err := s.repository.GetEnrollment()
	if err != "" {
		return nil, err
	}

	return enrollments, ""
}
