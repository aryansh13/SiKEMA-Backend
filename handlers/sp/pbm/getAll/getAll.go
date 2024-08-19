package getAllSPHandler

import (
	getAllSP "attendance-is/controllers/sp/pbm/getAll"
	util "attendance-is/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service getAllSP.Service
}

func NewGetAllSPHandler(service getAllSP.Service) *handler {
	return &handler{service: service}
}

func (h *handler) GetAllSPHandler(c *gin.Context) {
	res, err := h.service.GetAllSPService()
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
