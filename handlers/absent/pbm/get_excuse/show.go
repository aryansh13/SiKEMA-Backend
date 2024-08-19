package getAbsentExcuseHandler

import (
	getAbsentExcuse "attendance-is/controllers/absent/pbm/get_excuse"
	util "attendance-is/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service getAbsentExcuse.Service
}

func NewShowAbsentHandler(service getAbsentExcuse.Service) *handler {
	return &handler{service: service}
}

func (h *handler) ShowAbsentHandler(c *gin.Context) {
	var input getAbsentExcuse.InputGetAbsentExcuse
	input.AbsentID = c.Param("id")

	res, err := h.service.GetAbsentExcuseService(input)
	if err != "" {
		switch err {
		case "ABSENT_NOTFOUND_404":
			util.ErrorResponse(c, http.StatusNotFound, "Absent record not found")
			return
		case "EXCUSE_NOTFOUND_404":
			util.ErrorResponse(c, http.StatusNotFound, "Excuse record not found")
			return
		default:
			util.ErrorResponse(c, http.StatusInternalServerError, err)
			return
		}
	}
	util.APIResponse(c, http.StatusOK, res, nil)
}
