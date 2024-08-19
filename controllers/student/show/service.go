package showStudent

import (
	model "attendance-is/models"
	util "attendance-is/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetStudent(c *gin.Context) {
	var students []model.Student
	if err := util.DB.Find(&students).Error; util.HandleGORMError(c, err, http.StatusNotFound, "") {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": students,
	})
}

func ShowStudent(c *gin.Context) {
	id := c.Param("id")
	var student model.Student
	if err := util.DB.Where("id = ?", id).Find(&student).Error; util.HandleGORMError(c, err, http.StatusNotFound, "") {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": student,
	})
}

func GetStudentByClass(c *gin.Context) {
	// classId = c.Param("classId")
	// var students []model.Student
	// util.DB.Where
}
