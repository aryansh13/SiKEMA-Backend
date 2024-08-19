package class

import (
	schema "attendance-is/schemas/class"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *handler) GetClassHandler(c *gin.Context) {
	meta := schema.GetStudentMeta{Page: 1, ItemPerPage: 10}

	if c.Query("page") != "" {
		meta.Page, _ = strconv.ParseInt(c.Query("page"), 0, 64)
	}

	if c.Query("itemperpage") != "" {
		meta.ItemPerPage, _ = strconv.ParseInt(c.Query("itemperpage"), 0, 64)
	}

	class, err := h.service.GetClass(&meta)
	// var response []schema.GetClassResponse
	// for _, element := range *class {
	// 	response = append(response, element.ToGetClassResponse())
	// }

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"meta": meta,
		"data": class,
	})
}
