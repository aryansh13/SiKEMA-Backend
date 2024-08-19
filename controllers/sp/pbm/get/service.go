package getSP

import (
	model "attendance-is/models"
)

type Service interface {
	GetSPService(input InputGetSP) (*model.SP, string)
}

type service struct {
	repository Repository
}

func NewGetSPService(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) GetSPService(input InputGetSP) (*model.SP, string) {
	return s.repository.GetSPRepository(input.ID)
}
