package createExcuseHandler

import (
	createExcuse "attendance-is/controllers/excuse/student/create"
	util "attendance-is/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service createExcuse.Service
}

func NewCreateExcuseHandler(service createExcuse.Service) *handler {
	return &handler{service: service}
}

func (h *handler) CreateExcuseHandler(c *gin.Context) {
	var input createExcuse.InputCreateExcuse
	c.ShouldBindJSON(&input)

	res, err := h.service.CreateExcuseService(&input)
	if err != "" {
		switch err {
		case "EXCUSE_CONFLICT_409":
			util.ErrorResponse(c, http.StatusNotFound, "Excuse ID already exist")
			return
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
