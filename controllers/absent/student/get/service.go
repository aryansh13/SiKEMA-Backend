package getAbsent

import (
	model "attendance-is/models"
)

type Service interface {
	GetAbsentService(input InputGetAbsent) (*[]model.Absent, string)
}

type service struct {
	repository Repository
}

func NewGetAbsentService(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) GetAbsentService(input InputGetAbsent) (*[]model.Absent, string) {
	student, err := s.repository.GetStudentByID(input.StudentID)
	if err != "" {
		return nil, err
	}

	absent, err := s.repository.GetAbsentByWhere("student_id = ?", student.ID)
	if err != "" {
		return nil, err
	}

	return absent, ""
}
