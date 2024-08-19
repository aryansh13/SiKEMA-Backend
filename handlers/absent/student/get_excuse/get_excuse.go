package getAbsentExcuseHandler

import (
	getAbsentExcuse "attendance-is/controllers/absent/student/get_excuse"
	util "attendance-is/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service getAbsentExcuse.Service
}

func NewGetAbsentExcuseHandler(service getAbsentExcuse.Service) *handler {
	return &handler{service: service}
}

func (h *handler) GetAbsentExcuseHandler(c *gin.Context) {
	var input getAbsentExcuse.InputGetAbsentExcuse
	input.StudentID = c.Param("studentid")
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
