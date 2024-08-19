package validateAuthHandler

import (
	validateAuth "attendance-is/controllers/auth/validate"
	util "attendance-is/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service validateAuth.Service
}

func NewValidateAuthHandler(service validateAuth.Service) *handler {
	return &handler{service: service}
}

func (h *handler) ValidateAuthHandler(c *gin.Context) {
	token := util.ExtractToken(c)
	if token == "" {
		util.ErrorResponse(c, http.StatusUnauthorized, "Unauthenticated")
	}

	res, err := h.service.ValidateAuth(token)
	switch err {
	case "":
		util.APIResponse(c, http.StatusOK, res, nil)
		return
	default:
		util.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}
}
