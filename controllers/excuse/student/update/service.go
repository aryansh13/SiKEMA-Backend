package updateExcuse

import model "attendance-is/models"

type Service interface {
	UpdateExcuseService(input *InputUpdateExcuse) (*model.Excuse, string)
}

type service struct {
	repository Repository
}

func NewUpdateExcuseService(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) UpdateExcuseService(input *InputUpdateExcuse) (*model.Excuse, string) {
	excuse, err := s.repository.GetExcuseByID(input.ID)
	if err != "" {
		return nil, err
	}

	if input.Attachment != "" {
		excuse.Attachment = input.Attachment
	}

	if input.Excuse != "" {
		excuse.Excuse = input.Excuse
	}

	err = s.repository.SaveExcuse(excuse)
	if err != "" {
		return nil, err
	}

	return excuse, ""
}
