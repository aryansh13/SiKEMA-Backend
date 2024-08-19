package getCourseHandler

import (
	getCourse "attendance-is/controllers/course/student/get"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service getCourse.Service
}

func NewGetCourseHandler(service getCourse.Service) *handler {
	return &handler{service: service}
}

func (h *handler) GetCourseHandler(c *gin.Context) {
	var input getCourse.InputGetCourse
	input.StudentID = c.Param("studentid")
	input.ClassID = c.Param("classid")
	input.CourseID = c.Param("id")

	// res, err := h.service.GetCourseService(input)
	res, err := h.service.GetCourse(input)

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
