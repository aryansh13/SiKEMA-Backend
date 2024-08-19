package route

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	middleware "attendance-is/middlewares"

	getAbsent "attendance-is/controllers/absent/student/get"
	getAbsentExcuse "attendance-is/controllers/absent/student/get_excuse"
	showAbsent "attendance-is/controllers/absent/student/show"

	getAbsentHandler "attendance-is/handlers/absent/student/get"
	getAbsentExcuseHandler "attendance-is/handlers/absent/student/get_excuse"
	showAbsentHandler "attendance-is/handlers/absent/student/show"

	getAbsentPBM "attendance-is/controllers/absent/pbm/get"
	getAbsentExcusePBM "attendance-is/controllers/absent/pbm/get_excuse"
	showAbsentPBM "attendance-is/controllers/absent/pbm/show"

	getAbsentHandlerPBM "attendance-is/handlers/absent/pbm/get"
	getAbsentExcuseHandlerPBM "attendance-is/handlers/absent/pbm/get_excuse"
	showAbsentHandlerPBM "attendance-is/handlers/absent/pbm/show"
)

func InitAbsentRoute(db *gorm.DB, r *gin.Engine) {
	getAbsentRepo := getAbsent.NewGetAbsentRepository(db)
	getAbsentService := getAbsent.NewGetAbsentService(getAbsentRepo)
	getAbsentHandler := getAbsentHandler.NewGetAbsentHandler(getAbsentService)

	showAbsentRepo := showAbsent.NewShowAbsentRepository(db)
	showAbsentService := showAbsent.NewShowAbsentService(showAbsentRepo)
	showAbsentHandler := showAbsentHandler.NewShowAbsentHandler(showAbsentService)

	getAbsentExcuseRepo := getAbsentExcuse.NewGetAbsentExcuseRepository(db)
	getAbsentExcuseService := getAbsentExcuse.NewGetAbsentExcuseService(getAbsentExcuseRepo)
	getAbsentExcuseHandler := getAbsentExcuseHandler.NewGetAbsentExcuseHandler(getAbsentExcuseService)

	getAbsentPBMRepo := getAbsentPBM.NewGetAbsentRepository(db)
	getAbsentPBMService := getAbsentPBM.NewGetAbsentService(getAbsentPBMRepo)
	getAbsentPBMHandler := getAbsentHandlerPBM.NewGetAbsentHandler(getAbsentPBMService)

	showAbsentPBMRepo := showAbsentPBM.NewShowAbsentRepository(db)
	showAbsentPBMService := showAbsentPBM.NewShowAbsentService(showAbsentPBMRepo)
	showAbsentPBMHandler := showAbsentHandlerPBM.NewShowAbsentHandler(showAbsentPBMService)

	getAbsentExcusePBMRepo := getAbsentExcusePBM.NewGetAbsentExcuseRepository(db)
	getAbsentExcusePBMService := getAbsentExcusePBM.NewGetAbsentExcuseService(getAbsentExcusePBMRepo)
	getAbsentExcusePBMHandler := getAbsentExcuseHandlerPBM.NewShowAbsentHandler(getAbsentExcusePBMService)

	group := r.Group("/api/pbm/absent", middleware.Auth(), middleware.IsPBM())
	group.GET("", getAbsentPBMHandler.GetAbsentHandler)
	group.GET(":id/excuse", getAbsentExcusePBMHandler.ShowAbsentHandler)
	group.GET(":id", showAbsentPBMHandler.ShowAbsentHandler)

	studentGroup := r.Group("/api/student/:studentid", middleware.Auth(), middleware.IsStudent())
	studentGroup.GET("absent", getAbsentHandler.GetAbsentHandler)
	studentGroup.GET("absent/:id/excuse", getAbsentExcuseHandler.GetAbsentExcuseHandler)
	studentGroup.GET("event/:eventid/absent", showAbsentHandler.ShowAbsentHandler)
}
