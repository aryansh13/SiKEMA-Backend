package getRecentQRCode

import (
	model "attendance-is/models"

	"gorm.io/gorm"
)

type Repository interface {
	GetRecentQRCode(eventId string) (*model.QRCode, string)
}

type repository struct {
	db *gorm.DB
}

func NewGetRecentQRCodeRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) GetRecentQRCode(eventId string) (*model.QRCode, string) {
	var qrCode model.QRCode
	if err := r.db.Model(model.QRCode{}).Order("created_at DESC").Where("event_id = ?", eventId).First(&qrCode).Error; err != nil {
		return nil, "QR_UNEXPECTED_500 : " + err.Error()
	}
	return &qrCode, ""
}
