package updateEventHandler

import (
	updateEvent "attendance-is/controllers/event/lecturer/update"
	util "attendance-is/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service updateEvent.Service
}

func NewUpdateEventHandler(service updateEvent.Service) *handler {
	return &handler{service: service}
}

func (h *handler) UpdateEventHandler(c *gin.Context) {
	var input updateEvent.InputUpdateEvent
	c.ShouldBindJSON(&input)
	input.EventID = c.Param("id")

	err := h.service.UpdateEventService(input)
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
	c.Status(http.StatusNoContent)
}
