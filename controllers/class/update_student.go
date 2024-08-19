package class

import (
	schema "attendance-is/schemas/class"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) UpdateStudentFromClass(c *gin.Context) {
	var input schema.UpdateStudentFromClassRequest
	c.ShouldBindJSON(&input)

	class, err := h.service.UpdateStudentFromClass(input)

	if err != nil {
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": class,
	})
}
