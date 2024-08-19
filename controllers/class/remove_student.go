package class

import (
	schema "attendance-is/schemas/class"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *handler) RemoveStudentFromClassHandler(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 0)
	studentId, _ := strconv.ParseUint(c.Param("studentid"), 10, 0)
	input := schema.RemoveStudentFromClassRequest{ID: uint(id), Student: uint(studentId)}

	class, err := h.service.RemoveStudentFromClass(input)

	if err != nil {
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": class,
	})
}
