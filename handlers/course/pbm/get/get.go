package getCourseHandler

import (
	getCourse "attendance-is/controllers/course/pbm/get"
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
	res, err := h.service.GetCourseService()

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
