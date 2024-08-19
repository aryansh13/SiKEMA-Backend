package getAbsentHandler

import (
	getAbsent "attendance-is/controllers/absent/pbm/get"
	util "attendance-is/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service getAbsent.Service
}

func NewGetAbsentHandler(service getAbsent.Service) *handler {
	return &handler{service: service}
}

func (h *handler) GetAbsentHandler(c *gin.Context) {
	res, err := h.service.GetAbsentService()
	if err != "" {
		switch err {
		case "ABSENT_NOTFOUND_404":
			util.ErrorResponse(c, http.StatusNotFound, "Record not found")
			return
		default:
			util.ErrorResponse(c, http.StatusInternalServerError, err)
			return
		}
	}
	util.APIResponse(c, http.StatusOK, res, nil)
}
