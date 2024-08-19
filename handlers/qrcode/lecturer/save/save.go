package saveQRCodeHandler

import (
	saveQRCode "attendance-is/controllers/qrcode/lecturer/save"
	util "attendance-is/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service saveQRCode.Service
}

func NewSaveQRCodeHandler(service saveQRCode.Service) *handler {
	return &handler{service: service}
}

func (h *handler) SaveQRCodeHandler(c *gin.Context) {
	var input saveQRCode.InputSaveQRCode
	c.ShouldBindJSON(&input)
	input.EventID = c.Param("id")

	res, err := h.service.SaveQRCodeService(input)
	switch err {
	case "ERR_CONFLICT_409":
		util.ErrorResponse(c, http.StatusConflict, "SP ID already exist")
		return
	default:
		util.APIResponse(c, http.StatusOK, res, nil)
		return
	}
}
