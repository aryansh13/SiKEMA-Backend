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
	absent, err := s.repository.GetAbsentByID(input.AbsentID)
	if err != "" {
		return nil, err
	}

	return absent, ""
}
