package showEvent

import model "attendance-is/models"

type Service interface {
	ShowEventService(input InputShowEvent) (*model.Event, string)
}

type service struct {
	repository Repository
}

func NewShowEventService(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) ShowEventService(input InputShowEvent) (*model.Event, string) {
	event, err := s.repository.GetEventByID(input.LecturerID, input.EventID)
	if err != "" {
		return nil, err
	}
	return event, ""
}
