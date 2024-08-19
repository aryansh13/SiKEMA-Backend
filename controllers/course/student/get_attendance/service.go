package getCourseAttendance

import (
	model "attendance-is/models"
)

type Service interface {
	GetCourseAttendance(input InputGetCourseAttendance) (*[]model.Event, string)
}

type service struct {
	repository Repository
}

func NewGetCourseAttendanceService(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) GetCourseAttendance(input InputGetCourseAttendance) (*[]model.Event, string) {
	student, err := s.repository.GetStudent(input.StudentID)
	if err != "" {
		return nil, err
	}

	enrollment, err2 := s.repository.GetEnrollmentByID(student, input.CourseID)
	if err2 != "" {
		return nil, err2
	}

	events, err := s.repository.GetCourseEvent(&enrollment.Course, student)

	return events, err
}
