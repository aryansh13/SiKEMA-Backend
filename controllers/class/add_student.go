package class

import (
	schema "attendance-is/schemas/class"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) AddStudentToClassHandler(c *gin.Context) {
	var input schema.CreateClassRequest
	c.ShouldBindJSON(&input)

	class, err := h.service.CreateClass(input)

	if err != nil {
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": class,
	})
}
