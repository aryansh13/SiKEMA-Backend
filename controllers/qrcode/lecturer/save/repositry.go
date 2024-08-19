package saveQRCode

import (
	model "attendance-is/models"

	"gorm.io/gorm"
)

type Repository interface {
	SaveQRCode(qrCode *model.QRCode) string
}

type repository struct {
	db *gorm.DB
}

func NewSaveQRCodeRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) SaveQRCode(qrCode *model.QRCode) string {
	if err := r.db.Save(qrCode).Error; err != nil {
		return "QRCODE_UNEXPECTED_500 : " + err.Error()
	}
	return ""
}
