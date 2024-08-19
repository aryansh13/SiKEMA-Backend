package getAllExcuse

import (
	model "attendance-is/models"
)

type Service interface {
	GetAllExcuseService(input InputGetAllExcuse) (*[]model.Excuse, string)
}

type service struct {
	repository Repository
}

func NewGetAllExcuseService(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) GetAllExcuseService(input InputGetAllExcuse) (*[]model.Excuse, string) {
	student := model.Student{Nim: input.StudentID}

	res, err := s.repository.GetAllExcuseRepository(student)

	return res, err
}
