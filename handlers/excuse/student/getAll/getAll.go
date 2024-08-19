package getExcuseHandler

import (
	getAllExcuse "attendance-is/controllers/excuse/student/getAll"
	util "attendance-is/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service getAllExcuse.Service
}

func NewGetAllExcuseHandler(service getAllExcuse.Service) *handler {
	return &handler{service: service}
}

func (h *handler) GetAllExcuseHandler(c *gin.Context) {
	var input getAllExcuse.InputGetAllExcuse
	input.StudentID = c.Param("studentid")

	res, err := h.service.GetAllExcuseService(input)
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

// }
