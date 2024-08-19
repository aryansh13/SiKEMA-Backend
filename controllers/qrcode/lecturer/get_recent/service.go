package getRecentQRCode

import (
	model "attendance-is/models"
)

type Service interface {
	GetRecentQRCodeService(input InputGetRecentQRCode) (*model.QRCode, string)
}

type service struct {
	repository Repository
}

func NewGetRecentQRCodeService(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) GetRecentQRCodeService(input InputGetRecentQRCode) (*model.QRCode, string) {
	qr, err := s.repository.GetRecentQRCode(input.EventID)
	if err != "" {
		return nil, err
	}

	return qr, ""
}
