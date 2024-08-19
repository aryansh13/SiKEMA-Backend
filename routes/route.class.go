package route

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	class "attendance-is/controllers/class"
	service "attendance-is/services"
)

func InitClassRoute(db *gorm.DB, r *gin.Engine) {
	classService := service.NewClassService(db)
	classHandler := class.NewClassHandler(classService)

	group := r.Group("api/class")
	group.DELETE(":id/student/:studentid", classHandler.RemoveStudentFromClassHandler)
	group.GET("", classHandler.GetClassHandler)
	group.GET(":id", classHandler.ShowClassHandler)
	group.PATCH(":id", classHandler.UpdateClassHandler)
	group.PATCH(":id/student", classHandler.UpdateStudentFromClass)
	group.POST("", classHandler.CreateClassHandler)
	group.POST(":id/student", classHandler.AddStudentToClassHandler)
	group.PUT(":id", classHandler.UpdateClassHandler)
	group.PUT(":id/student", classHandler.UpdateStudentFromClass)
}
