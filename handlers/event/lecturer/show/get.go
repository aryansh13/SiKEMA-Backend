package showEventHandler

import (
	showEvent "attendance-is/controllers/event/lecturer/show"
	util "attendance-is/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service showEvent.Service
}

func NewShowEventHandler(service showEvent.Service) *handler {
	return &handler{service: service}
}

func (h *handler) ShowEventHandler(c *gin.Context) {
	var input showEvent.InputShowEvent
	input.EventID = c.Param("id")
	input.LecturerID = c.Param("lecturerid")
	res, err := h.service.ShowEventService(input)
	if err != "" {
		switch err {
		case "EVENT_NOTFOUND_404":
			util.ErrorResponse(c, http.StatusNotFound, "Event record not found")
			return
		default:
			util.ErrorResponse(c, http.StatusInternalServerError, err)
			return
		}
	}
	util.APIResponse(c, http.StatusOK, res, nil)
}
