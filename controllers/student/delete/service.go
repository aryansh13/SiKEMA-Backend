package deleteStudent

import (
	model "attendance-is/models"
	util "attendance-is/utils"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DeleteStudent(c *gin.Context) {
	id := c.Param("id")
	if err := util.DB.Delete(model.Student{}, id).Error; util.HandleGORMError(c, err, http.StatusInternalServerError, "") {
		return
	}

	c.Status(http.StatusNoContent)
}

func MassDeleteStudent(c *gin.Context) {
	var input MassInput
	c.ShouldBindJSON(&input)

	var students []model.Student
	if err := util.DB.Where("id IN ?", input.Id).Delete(&students).Error; util.HandleGORMError(c, err, (map[bool]int{true: http.StatusNotFound, false: http.StatusInternalServerError})[errors.Is(err, gorm.ErrRecordNotFound)], "") {
		return
	}

	c.Status(http.StatusNoContent)
}
