package removeStudentEvent

type Service interface {
	RemoveStudentToEventService(input InputRemoveStudent) string
}

type service struct {
	repository Repository
}

func NewRemoveStudentToEventService(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) RemoveStudentToEventService(input InputRemoveStudent) string {
	event, err := s.repository.GetEventByID(input.EventId)
	if err != "" {
		return err
	}

	students, err := s.repository.GetStudentByWhere("nim IN ?", input.StudentId)
	if err != "" {
		return err
	}

	err = s.repository.RemoveStudent(event, students)
	if err != "" {
		return err
	}

	return ""
}
