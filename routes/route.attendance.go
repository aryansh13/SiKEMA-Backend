package route

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	middleware "attendance-is/middlewares"
)

func InitAttendanceRoute(db *gorm.DB, r *gin.Engine) {
	studentGroup := r.Group("/api/student/:studentid", middleware.Auth(), middleware.IsStudent())
	studentGroup.GET("course/:courseid/attendance")
}
