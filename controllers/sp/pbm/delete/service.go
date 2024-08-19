package deleteSP

import (
	model "attendance-is/models"
)

type Service interface {
	DeleteSPService(input InputDeleteSP) string
}

type service struct {
	repository Repository
}

func NewDeleteSPService(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) DeleteSPService(input InputDeleteSP) string {
	sp := model.SP{ID: input.ID}

	return s.repository.DeleteSPRepository(&sp)
}
