package getExcuseHandler

import (
	getExcuse "attendance-is/controllers/excuse/student/get"
	util "attendance-is/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service getExcuse.Service
}

func NewGetExcuseHandler(service getExcuse.Service) *handler {
	return &handler{service: service}
}

func (h *handler) GetExcuseHandler(c *gin.Context) {
	var input getExcuse.InputGetExcuse
	input.StudentID = c.Param("studentid")
	id, _ := strconv.Atoi(c.Param("id"))
	input.ID = uint(id)

	res, err := h.service.GetExcuseService(input)
	if err != "" {
		switch err {
		case "EXCUSE_NOTFOUND_404":
			util.ErrorResponse(c, http.StatusNotFound, "Record not found")
			return
		default:
			util.ErrorResponse(c, http.StatusInternalServerError, err)
			return
		}
	}
	util.APIResponse(c, http.StatusOK, res, nil)
}
