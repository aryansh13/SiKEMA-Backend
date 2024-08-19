package getAllSP

import (
	model "attendance-is/models"
)

type Service interface {
	GetAllSPService() (*[]model.SP, string)
}

type service struct {
	repository Repository
}

func NewGetAllSPService(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) GetAllSPService() (*[]model.SP, string) {
	return s.repository.GetAllSPRepository()
}
