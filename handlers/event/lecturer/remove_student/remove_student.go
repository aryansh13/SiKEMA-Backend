package removeStudentEventHandler

import (
	removeStudentEvent "attendance-is/controllers/event/lecturer/remove_student"
	util "attendance-is/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service removeStudentEvent.Service
}

func NewRemoveStudentEventHandler(service removeStudentEvent.Service) *handler {
	return &handler{service: service}
}

func (h *handler) RemoveStudentEvent(c *gin.Context) {
	var input removeStudentEvent.InputRemoveStudent
	c.ShouldBindJSON(&input)

	input.EventId = c.Param("id")

	err := h.service.RemoveStudentToEventService(input)
	if err != "" {
		switch err {
		case "STUDENT_NOTFOUND_404":
			util.ErrorResponse(c, http.StatusNotFound, "Student NIM not found")
			return
		case "EVENT_NOTFOUND_404":
			util.ErrorResponse(c, http.StatusNotFound, "Event ID not found")
			return
		default:
			util.ErrorResponse(c, http.StatusInternalServerError, err)
			return
		}
	}
	c.Status(http.StatusNoContent)
}
