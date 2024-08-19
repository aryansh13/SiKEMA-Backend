package getAbsent

import (
	model "attendance-is/models"
)

type Service interface {
	GetAbsentService() (*[]model.Absent, string)
}

type service struct {
	repository Repository
}

func NewGetAbsentService(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) GetAbsentService() (*[]model.Absent, string) {
	absent, err := s.repository.GetAbsent()
	if err != "" {
		return nil, err
	}

	return absent, ""
}
