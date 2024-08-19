package route

import (
	// updateClass "attendance-is/controllers/class/update"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitStudentRoute(db *gorm.DB, r *gin.Engine) {
	// group := r.Group("/api/class")
	// group.DELETE(":id", updateClass.RemoveStudent)
	// group.GET("")
	// group.GET(":id")
	// group.PATCH(":id", updateClass.UpdateClass)
	// group.PATCH(":id/student", updateClass.ReplaceStudent)
	// group.POST("")
	// group.POST(":id/student", updateClass.AppendStudent)
	// group.PUT(":id", updateClass.UpdateClass)
	// group.PUT(":id/student", updateClass.ReplaceStudent)
}
