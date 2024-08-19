package getCourse

import (
	model "attendance-is/models"
)

type Service interface {
	GetCourseService(input InputGetCourse) (*[]model.Enrollment, string)
}

type service struct {
	repository Repository
}

func NewGetCourseService(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) GetCourseService(input InputGetCourse) (*[]model.Enrollment, string) {
	lecturer, err := s.repository.GetLecturerByID(input.LecturerID)
	if err != "" {
		return nil, err
	}

	enrollments, err := s.repository.GetEnrollmentByLecturer(lecturer, input.ClassID, input.CourseID)
	if err != "" {
		return nil, err
	}

	return enrollments, ""
}
