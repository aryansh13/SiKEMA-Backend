package createSPHandler

import (
	createSP "attendance-is/controllers/sp/pbm/create"
	util "attendance-is/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service createSP.Service
}

func NewCreateSPHandler(service createSP.Service) *handler {
	return &handler{service: service}
}

func (h *handler) CreateSPHandler(c *gin.Context) {
	var input createSP.InputCreateSP
	c.ShouldBindJSON(&input)

	res, err := h.service.CreateSPService(input)
	if err != "" {
		switch err {
		case "ERR_CONFLICT_409":
			util.ErrorResponse(c, http.StatusConflict, "SP ID already exist")
			return
		case "ERR_UNEXPECTED_500":
			util.ErrorResponse(c, http.StatusInternalServerError, "There's an error creating new resource")
			return
		}
	}

	util.APIResponse(c, http.StatusOK, res, nil)
}
