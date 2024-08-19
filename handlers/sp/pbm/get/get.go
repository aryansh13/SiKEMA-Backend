package getSPHandler

import (
	getSP "attendance-is/controllers/sp/pbm/get"
	util "attendance-is/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service getSP.Service
}

func NewGetSPHandler(service getSP.Service) *handler {
	return &handler{service: service}
}

func (h *handler) GetSPHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	input := getSP.InputGetSP{
		ID: uint(id),
	}

	res, err := h.service.GetSPService(input)
	if err != "" {
		switch err {
		case "ERR_CONFLICT_409":
			util.ErrorResponse(c, http.StatusConflict, "SP ID already exist")
			return
		case "ERR_NOTFOUND_404":
			util.ErrorResponse(c, http.StatusNotFound, "SP ID not found")
			return
		default:
			util.ErrorResponse(c, http.StatusInternalServerError, "Unexpected Error : "+err)
			return
		}
	}

	util.APIResponse(c, http.StatusOK, res, nil)
}
