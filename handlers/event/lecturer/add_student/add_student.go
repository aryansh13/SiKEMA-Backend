package addStudentEventHandler

import (
	addStudentEvent "attendance-is/controllers/event/lecturer/add_student"
	util "attendance-is/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service addStudentEvent.Service
}

func NewAddStudentEventHandler(service addStudentEvent.Service) *handler {
	return &handler{service: service}
}

func (h *handler) AddStudentEventHandler(c *gin.Context) {
	var input addStudentEvent.InputAddStudent
	c.ShouldBindJSON(&input)
	input.EventId = c.Param("id")

	err := h.service.AddStudentToEventService(input)
	if err != "" {
		switch err {
		case "EVENT_NOTFOUND_404":
			util.ErrorResponse(c, http.StatusNotFound, "Event ID not found")
			return
		default:
			util.ErrorResponse(c, http.StatusInternalServerError, err)
			return
		}
	}
	util.APIResponse(c, http.StatusAccepted, nil, nil)
}
