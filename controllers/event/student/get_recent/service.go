package getRecentEvent

import (
	model "attendance-is/models"
	"strconv"
)

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
	student, err := s.repository.GetStudent(input.StudentID)
	if err != "" {
		return nil, err
	}

	classId := strconv.Itoa(int(student.ClassID))
	event, err := s.repository.GetRecentEvent(classId)
	if err != "" {
		return nil, err
	}
	return event, ""
}
