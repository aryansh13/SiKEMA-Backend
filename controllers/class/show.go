package class

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) ShowClassHandler(c *gin.Context) {
	id := c.Param("id")
	class, err := h.service.ShowClass(id)

	if err != nil {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": class,
	})
}
