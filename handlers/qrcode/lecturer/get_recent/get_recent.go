package getRecentQRCodeHandler

import (
	getRecentQRCode "attendance-is/controllers/qrcode/lecturer/get_recent"
	util "attendance-is/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service getRecentQRCode.Service
}

func NewGetRecentQRCodeHandler(service getRecentQRCode.Service) *handler {
	return &handler{service: service}
}

func (h *handler) GetRecentQRCodeHandler(c *gin.Context) {
	var input getRecentQRCode.InputGetRecentQRCode
	input.EventID = c.Param("id")

	res, err := h.service.GetRecentQRCodeService(input)
	switch err {
	case "ERR_CONFLICT_409":
		util.ErrorResponse(c, http.StatusConflict, "SP ID already exist")
		return
	default:
		util.APIResponse(c, http.StatusOK, res, nil)
		return
	}
}
