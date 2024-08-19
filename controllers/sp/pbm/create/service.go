package createSP

import model "attendance-is/models"

type Service interface {
	CreateSPService(input InputCreateSP) (*model.SP, string)
}

type service struct {
	repository Repository
}

func NewCreateSPService(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) CreateSPService(input InputCreateSP) (*model.SP, string) {
	newSP := model.SP{
		StudentID: input.StudentID,
		Semester:  input.Semester,
		Level:     input.Level,
		Status:    input.Status,
	}

	res, err := s.repository.CreateSPRepository(&newSP)

	return res, err
}
