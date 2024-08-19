package getAbsentExcuse

import (
	model "attendance-is/models"
)

type Service interface {
	GetAbsentExcuseService(input InputGetAbsentExcuse) (*model.Excuse, string)
}

type service struct {
	repository Repository
}

func NewGetAbsentExcuseService(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) GetAbsentExcuseService(input InputGetAbsentExcuse) (*model.Excuse, string) {
	absent, err := s.repository.GetAbsentByID(input.AbsentID)
	if err != "" {
		return nil, err
	}

	excuse, err := s.repository.GetAbsentExcuse(absent)
	if err != "" {
		return nil, err
	}

	return excuse, ""
}
