package showAbsent

import (
	model "attendance-is/models"
)

type Service interface {
	ShowAbsentService(input InputShowAbsent) (*model.Absent, string)
}

type service struct {
	repository Repository
}

func NewShowAbsentService(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) ShowAbsentService(input InputShowAbsent) (*model.Absent, string) {
	student, err := s.repository.GetStudentByID(input.StudentID)
	if err != "" {
		return nil, err
	}

	absent, err := s.repository.GetAbsentByWhere("student_id = ? AND event_id = ?", student.ID, input.EventID)
	if err != "" {
		return nil, err
	}

	return absent, ""
}
