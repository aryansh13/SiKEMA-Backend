package getEvent

import model "attendance-is/models"

type Service interface {
	GetEventService(input InputGetEvent) (*[]model.Event, string)
}

type service struct {
	repository Repository
}

func NewGetEventService(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) GetEventService(input InputGetEvent) (*[]model.Event, string) {
	event, err := s.repository.GetEvent(input.LecturerID, input.ClassID, input.CourseID)
	if err != "" {
		return nil, err
	}
	return event, ""
}
