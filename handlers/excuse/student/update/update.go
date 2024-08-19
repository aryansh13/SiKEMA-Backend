package excuseHandler

import (
	updateExcuse "attendance-is/controllers/excuse/student/update"
	util "attendance-is/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service updateExcuse.Service
}

func NewUpdateExcuseHandler(service updateExcuse.Service) *handler {
	return &handler{service: service}
}

func (h *handler) UpdateExcuseHandler(c *gin.Context) {
	var input updateExcuse.InputUpdateExcuse
	c.ShouldBindJSON(&input)
	input.ID = c.Param("id")

	_, err := h.service.UpdateExcuseService(&input)
	if err != "" {
		switch err {
		case "EXCUSE_NOTFOUND_404":
			util.ErrorResponse(c, http.StatusNotFound, "Record not found")
			return
		default:
			util.ErrorResponse(c, http.StatusInternalServerError, err)
			return
		}
	}
	util.APIResponse(c, http.StatusAccepted, nil, nil)

}
