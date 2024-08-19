package compensationGet

import (
	model "attendance-is/models"
)

type Service interface {
	GetCompensation() (*[]model.Student, string)
}

type service struct {
	repository Repository
}

func NewCompensationGetService(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) GetCompensation() (*[]model.Student, string) {
	students, _ := s.repository.GetStudent()

	for i := range *students {
		absents, _ := s.repository.GetAbsent(&(*students)[i])
		hours := 0
		for _, absent := range *absents {
			excuse, _ := s.repository.GetExcuse(&absent)
			if excuse == nil {
				hours += int(absent.Hours)
				continue
			}

			if excuse.Status != 2 {
				hours += int(absent.Hours)
			}
		}
		compensationHours := uint(hours)
		(*students)[i].CompensationHours = &compensationHours
	}

	return students, ""
}
