package loginAuthHandler

import (
	loginAuth "attendance-is/controllers/auth/login"
	model "attendance-is/models"
	util "attendance-is/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type response struct {
	Token    string     `json:"access_token"`
	UserData model.User `json:"user_data"`
}

type handler struct {
	service loginAuth.Service
}

func NewLoginAuthHandler(service loginAuth.Service) *handler {
	return &handler{service: service}
}

func (h *handler) LoginAuthHandler(c *gin.Context) {
	var input loginAuth.InputLoginAuth
	c.ShouldBindJSON(&input)

	user, token, err := h.service.LoginAuthService(input)
	response := response{Token: token, UserData: *user}
	switch err {
	case "LOGIN_UNAUTHENTICATED_401":
		util.ErrorResponse(c, http.StatusUnauthorized, "Unauthenticated")
		return
	case "":
		util.APIResponse(c, http.StatusOK, response, nil)
		return
	default:
		util.ErrorResponse(c, http.StatusUnauthorized, "Unauthenticated :"+err)
		return
	}
}
