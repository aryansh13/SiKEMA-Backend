package createEvent

import (
	model "attendance-is/models"
	"strconv"
)

type Service interface {
	CreateEventService(input InputCreateEvent) (*model.Event, string)
}

type service struct {
	repository Repository
}

func NewCreateEventService(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) CreateEventService(input InputCreateEvent) (*model.Event, string) {
	lecturer, err := s.repository.GetLecturerByID(strconv.Itoa(int(input.LecturerId)))
	if err != "" {
		return nil, ""
	}

	enrollment, err := s.repository.GetEnrollmentByWhere(lecturer, strconv.Itoa(int(input.CourseId)), strconv.Itoa(int(input.ClassId)))
	if err != "" {
		return nil, ""
	}
	event := model.Event{
		ClassID:    enrollment.ClassID,
		CourseID:   enrollment.CourseID,
		LecturerID: input.LecturerId,
		Meet:       input.Meet,
	}

	err = s.repository.CreateEvent(&event)
	if err != "" {
		return nil, ""
	}
	return &event, ""
}
