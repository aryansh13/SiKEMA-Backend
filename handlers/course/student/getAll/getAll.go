package getAllCourseHandler

import (
	getAllCourse "attendance-is/controllers/course/student/getAll"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service getAllCourse.Service
}

func NewGetAllCourseHandler(service getAllCourse.Service) *handler {
	return &handler{service: service}
}

func (h *handler) GetAllCourseHandler(c *gin.Context) {
	var input getAllCourse.InputGetAllCourse
	input.StudentID = c.Param("studentid")
	input.ClassID = c.Param("classid")

	res, err := h.service.GetAllCourse(input)

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
