package updateSP

import (
	updateSP "attendance-is/controllers/sp/pbm/update"
	util "attendance-is/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service updateSP.Service
}

func NewUpdateSPHandler(service updateSP.Service) *handler {
	return &handler{service: service}
}

func (h *handler) UpdateSPHandler(c *gin.Context) {
	var input updateSP.InputUpdateSP
	c.ShouldBindJSON(&input)

	res, err := h.service.UpdateSPService(input)
	switch err {
	case "ERR_CONFLICT_409":
		util.ErrorResponse(c, http.StatusConflict, "SP ID already exist")
		return
	case "ERR_UNEXPECTED_500":
		util.ErrorResponse(c, http.StatusInternalServerError, "There's an error creating new resource")
		return
	default:
		util.APIResponse(c, http.StatusOK, res, nil)
		return
	}
}
