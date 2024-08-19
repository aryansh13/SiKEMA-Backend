package showAbsentHandler

import (
	showAbsent "attendance-is/controllers/absent/pbm/show"
	util "attendance-is/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service showAbsent.Service
}

func NewShowAbsentHandler(service showAbsent.Service) *handler {
	return &handler{service: service}
}

func (h *handler) ShowAbsentHandler(c *gin.Context) {
	var input showAbsent.InputShowAbsent
	input.AbsentID = c.Param("id")

	res, err := h.service.ShowAbsentService(input)
	if err != "" {
		switch err {
		case "ABSENT_NOTFOUND_404":
			util.ErrorResponse(c, http.StatusNotFound, "Absent record not found")
			return
		default:
			util.ErrorResponse(c, http.StatusInternalServerError, err)
			return
		}
	}
	util.APIResponse(c, http.StatusOK, res, nil)
}
