package getAllSPHandler

import (
	getAllSP "attendance-is/controllers/sp/student/getAll"
	util "attendance-is/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service getAllSP.Service
}

func NewGetAllSPHandler(service getAllSP.Service) *handler {
	return &handler{service: service}
}

func (h *handler) GetAllSPHandler(c *gin.Context) {
	student_id, _ := strconv.Atoi(c.Param("studentid"))

	input := getAllSP.InputGetAllSP{
		StudentID: uint(student_id),
	}

	res, err := h.service.GetAllSPService(input)
	switch err {
	case "ERR_CONFLICT_409":
		util.ErrorResponse(c, http.StatusConflict, "SP ID already exist")
		return
	default:
		util.APIResponse(c, http.StatusOK, res, nil)
		return
	}
}
