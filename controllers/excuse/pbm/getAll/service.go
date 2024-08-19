package getAllExcuse

import model "attendance-is/models"

type Service interface {
	GetAllExcuseService() (*[]model.Excuse, string)
}

type service struct {
	repository Repository
}

func NewGetAllExcuseService(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) GetAllExcuseService() (*[]model.Excuse, string) {
	res, _ := s.repository.GetAllExcuseRepository()
	return res, ""
}
