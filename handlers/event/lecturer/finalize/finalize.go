package finalizeEventHandler

import (
	finalizeEvent "attendance-is/controllers/event/lecturer/finalize"
	util "attendance-is/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service finalizeEvent.Service
}

func NewFinalizeEventHandler(service finalizeEvent.Service) *handler {
	return &handler{service: service}
}

func (h *handler) FinalizeEventHandler(c *gin.Context) {
	var input finalizeEvent.InputFinalizeEvent
	c.ShouldBindJSON(&input)
	input.EventID = c.Param("id")

	err := h.service.FinalizeEventService(input)
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
