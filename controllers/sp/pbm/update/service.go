package updateSP

import model "attendance-is/models"

type Service interface {
	UpdateSPService(input InputUpdateSP) (*model.SP, string)
}

type service struct {
	repository Repository
}

func NewUpdateSPService(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) UpdateSPService(input InputUpdateSP) (*model.SP, string) {
	newSP := model.SP{
		StudentID: input.StudentID,
		Semester:  input.Semester,
		Level:     input.Level,
		Status:    input.Status,
	}

	res, err := s.repository.UpdateSPRepository(&newSP)

	return res, err
}
