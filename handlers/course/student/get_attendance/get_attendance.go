package getCourseAttendanceHandler

import (
	getCourseAttendance "attendance-is/controllers/course/student/get_attendance"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service getCourseAttendance.Service
}

func NewGetCourseAttendanceHandler(service getCourseAttendance.Service) *handler {
	return &handler{service: service}
}

func (h *handler) GetCourseAttendanceHandler(c *gin.Context) {
	var input getCourseAttendance.InputGetCourseAttendance
	input.StudentID = c.Param("studentid")
	input.CourseID = c.Param("id")

	// res, err := h.service.GetCourseAttendanceService(input)
	res, err := h.service.GetCourseAttendance(input)

	if err != "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": res,
	})
}
