package deleteSP

import (
	deleteSP "attendance-is/controllers/sp/pbm/delete"
	util "attendance-is/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service deleteSP.Service
}

func NewDeleteSPHandler(service deleteSP.Service) *handler {
	return &handler{service: service}
}

func (h *handler) DeleteSPHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	input := deleteSP.InputDeleteSP{ID: uint(id)}
	err := h.service.DeleteSPService(input)

	switch err {
	case "ERR_CONFLICT_409":
		util.ErrorResponse(c, http.StatusConflict, "SP ID already exist")
		return
	default:
		util.APIResponse(c, http.StatusOK, nil, nil)
		return
	}
}
