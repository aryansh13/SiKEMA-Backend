package getEventHandler

import (
	getEvent "attendance-is/controllers/event/lecturer/get"
	util "attendance-is/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service getEvent.Service
}

func NewGetEventHandler(service getEvent.Service) *handler {
	return &handler{service: service}
}

func (h *handler) GetEventHandler(c *gin.Context) {
	input := getEvent.InputGetEvent{LecturerID: c.Param("lecturerid"), ClassID: c.Query("class_id"), CourseID: c.Query("course_id")}
	res, err := h.service.GetEventService(input)
	if err != "" {
		switch err {
		case "EVENT_NOTFOUND_404":
			util.ErrorResponse(c, http.StatusNotFound, "Record not found")
			return
		default:
			util.ErrorResponse(c, http.StatusInternalServerError, err)
			return
		}
	}
	util.APIResponse(c, http.StatusOK, res, nil)
}
