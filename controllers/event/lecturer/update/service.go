package updateEvent

type Service interface {
	UpdateEventService(input InputUpdateEvent) string
}

type service struct {
	repository Repository
}

func NewUpdateEventService(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) UpdateEventService(input InputUpdateEvent) string {
	event, err := s.repository.GetEventByID(input.EventID)
	if err != "" {
		return err
	}

	event.Status = input.Status

	err = s.repository.SaveEvent(event)
	if err != "" {
		return err
	}

	return ""
}
