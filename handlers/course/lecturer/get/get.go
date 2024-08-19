package getCourseHandler

import (
	getCourse "attendance-is/controllers/course/lecturer/get"
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
	input.LecturerID = c.Param("lecturerid")
	input.ClassID = c.Query("class_id")
	input.CourseID = c.Query("course_id")

	res, err := h.service.GetCourseService(input)

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
