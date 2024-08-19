package finalizeEvent

import model "attendance-is/models"

type Service interface {
	FinalizeEventService(input InputFinalizeEvent) string
}

type service struct {
	repository Repository
}

func NewFinalizeEventService(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) FinalizeEventService(input InputFinalizeEvent) string {
	event, err := s.repository.GetEventByID(input.EventID)
	if err != "" {
		return err
	}

	var IDs []uint
	if len(event.Students) > 0 {
		for _, element := range event.Students {
			IDs = append(IDs, element.ID)
		}
	} else {
		IDs = append(IDs, 0)
	}

	students, err := s.repository.GetStudentByWhere("class_id = ? AND id NOT IN ?", event.ClassID, IDs)
	if err != "" {
		return err
	}

	for _, student := range *students {
		_, err := s.repository.CreateAbsent(model.Absent{Student: student, Event: *event, Hours: uint(event.Course.Hours)})
		if err != "" {
			return err
		}
	}

	event.Status = 2

	err = s.repository.UpdateEvent(event)
	if err != "" {
		return err
	}

	return ""
}
