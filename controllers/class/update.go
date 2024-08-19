package class

import (
	schema "attendance-is/schemas/class"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *handler) UpdateClassHandler(c *gin.Context) {
	data := schema.UpdateClassRequest{}
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))
	data.ID = uint(id)

	data2, err := h.service.UpdateClass(data)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	// c.Status(http.StatusNoContent)
	c.JSON(http.StatusOK, gin.H{
		"data": data2,
	})
}
