package getRecentEvent

import model "attendance-is/models"

type Service interface {
	GetRecentEventService(input InputGetRecentEvent) (*model.Event, string)
}

type service struct {
	repository Repository
}

func NewGetRecentEventService(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) GetRecentEventService(input InputGetRecentEvent) (*model.Event, string) {
	event, err := s.repository.GetRecentEvent(input.LecturerID)
	if err != "" {
		return nil, err
	}
	return event, ""
}
