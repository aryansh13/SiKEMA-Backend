package addStudentEvent

type Service interface {
	AddStudentToEventService(input InputAddStudent) string
}

type service struct {
	repository Repository
}

func NewAddStudentToEventService(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) AddStudentToEventService(input InputAddStudent) string {
	event, err := s.repository.GetEventByID(input.EventId)
	if err != "" {
		return err
	}

	students, err := s.repository.GetStudentByWhere("nim IN ? AND class_id = ?", input.StudentId, event.ClassID)
	if err != "" {
		return err
	}

	err = s.repository.AddStudent(event, students)
	if err != "" {
		return err
	}

	return ""
}
