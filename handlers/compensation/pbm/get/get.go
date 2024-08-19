package getCompensationHandler

import (
	getCompensation "attendance-is/controllers/compensation/pbm/get"
	util "attendance-is/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service getCompensation.Service
}

func NewGetCompensationHandler(service getCompensation.Service) *handler {
	return &handler{service: service}
}

func (h *handler) GetCompensationHandler(c *gin.Context) {
	res, err := h.service.GetCompensation()
	switch err {
	case "ERR_CONFLICT_409":
		util.ErrorResponse(c, http.StatusConflict, "SP ID already exist")
		return
	default:
		util.APIResponse(c, http.StatusOK, res, nil)
		return
	}
}
