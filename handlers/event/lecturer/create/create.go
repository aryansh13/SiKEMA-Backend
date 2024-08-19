package createEventHandler

import (
	createEvent "attendance-is/controllers/event/lecturer/create"
	util "attendance-is/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service createEvent.Service
}

func NewCreateEventHandler(service createEvent.Service) *handler {
	return &handler{service: service}
}

func (h *handler) CreateEventHandler(c *gin.Context) {
	var input createEvent.InputCreateEvent
	c.ShouldBindJSON(&input)
	lecturerId, _ := strconv.Atoi(c.Param("lecturerid"))
	input.LecturerId = uint(lecturerId)

	res, err := h.service.CreateEventService(input)
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
	util.APIResponse(c, http.StatusAccepted, res, nil)
}
