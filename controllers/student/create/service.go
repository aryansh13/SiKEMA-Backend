package createStudent

import (
	model "attendance-is/models"
	util "attendance-is/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateStudent(c *gin.Context) {
	var input Input
	c.ShouldBindJSON(&input)

	var class model.Class
	util.DB.Where("id = ?").Find(&class)

	newStudent := model.Student{
		ClassID: class.ID,
		Nim:     input.Nim,
		Name:    input.Name,
	}

	if err := util.DB.Create(&newStudent).Error; util.HandleGORMError(c, err, http.StatusInternalServerError, "") {
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": newStudent,
	})
}
