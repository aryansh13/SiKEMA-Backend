package saveQRCode

import (
	model "attendance-is/models"
	"strconv"
)

type Service interface {
	SaveQRCodeService(input InputSaveQRCode) (*model.QRCode, string)
}

type service struct {
	repository Repository
}

func NewSaveQRCodeService(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) SaveQRCodeService(input InputSaveQRCode) (*model.QRCode, string) {
	eventId, _ := strconv.ParseUint(input.EventID, 10, 64)
	newQRCode := model.QRCode{
		EventID:  uint(eventId),
		SSID:     input.SSID,
		Password: input.Password,
	}

	err := s.repository.SaveQRCode(&newQRCode)
	if err != "" {
		return nil, err
	}

	return &newQRCode, ""
}
