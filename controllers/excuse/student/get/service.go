package getExcuse

import (
	model "attendance-is/models"
)

type Service interface {
	GetExcuseService(input InputGetExcuse) (*model.Excuse, string)
}

type service struct {
	repository Repository
}

func NewGetExcuseService(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) GetExcuseService(input InputGetExcuse) (*model.Excuse, string) {
	res, err := s.repository.GetExcuseRepository(input)

	return res, err
}
