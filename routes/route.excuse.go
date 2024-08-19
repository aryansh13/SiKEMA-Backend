package route

import (
	// excuse "attendance-is/controllers/excuse"
	// service "attendance-is/services"

	// middleware "attendance-is/middlewares"

	createExcuse "attendance-is/controllers/excuse/student/create"
	getExcuse "attendance-is/controllers/excuse/student/get"
	getAllExcuse "attendance-is/controllers/excuse/student/getAll"
	updateExcuse "attendance-is/controllers/excuse/student/update"
	createExcuseHandler "attendance-is/handlers/excuse/student/create"
	getExcuseHandler "attendance-is/handlers/excuse/student/get"
	getAllExcuseHandler "attendance-is/handlers/excuse/student/getAll"
	updateExcuseHandler "attendance-is/handlers/excuse/student/update"
	middleware "attendance-is/middlewares"

	getExcuseByPBM "attendance-is/controllers/excuse/pbm/get"
	getAllExcuseByPBM "attendance-is/controllers/excuse/pbm/getAll"
	updateExcuseByPBM "attendance-is/controllers/excuse/pbm/update"
	getExcuseByPBMHandler "attendance-is/handlers/excuse/pbm/get"
	getAllExcuseByPBMHandler "attendance-is/handlers/excuse/pbm/getAll"
	updateExcuseByPBMHandler "attendance-is/handlers/excuse/pbm/update"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitExcuseRoute(db *gorm.DB, router *gin.Engine) {
	getExcuseRepository := getExcuse.NewGetExcuseRepository(db)
	getExcuseService := getExcuse.NewGetExcuseService(getExcuseRepository)
	getExcuseHandler := getExcuseHandler.NewGetExcuseHandler(getExcuseService)

	getAllExcuseRepository := getAllExcuse.NewGetAllExcuseRepository(db)
	getAllExcuseService := getAllExcuse.NewGetAllExcuseService(getAllExcuseRepository)
	getAllExcuseHandler := getAllExcuseHandler.NewGetAllExcuseHandler(getAllExcuseService)

	createExcuseRepository := createExcuse.NewCreateExcuseRepository(db)
	createExcuseService := createExcuse.NewCreateExcuseService(createExcuseRepository)
	createExcuseHandler := createExcuseHandler.NewCreateExcuseHandler(createExcuseService)

	updateExcuseRepository := updateExcuse.NewUpdateExcuseRepository(db)
	updateExcuseService := updateExcuse.NewUpdateExcuseService(updateExcuseRepository)
	updateExcuseHandler := updateExcuseHandler.NewUpdateExcuseHandler(updateExcuseService)

	getExcuseByPBMRepository := getExcuseByPBM.NewGetExcuseRepository(db)
	getExcuseByPBMService := getExcuseByPBM.NewGetExcuseService(getExcuseByPBMRepository)
	getExcuseByPBMHandler := getExcuseByPBMHandler.NewGetExcuseHandler(getExcuseByPBMService)

	getAllExcuseByPBMRepository := getAllExcuseByPBM.NewGetAllExcuseRepository(db)
	getAllExcuseByPBMService := getAllExcuseByPBM.NewGetAllExcuseService(getAllExcuseByPBMRepository)
	getAllExcuseByPBMHandler := getAllExcuseByPBMHandler.NewGetAllExcuseHandler(getAllExcuseByPBMService)

	updateExcuseByPBMRepository := updateExcuseByPBM.NewUpdateExcuseRepository(db)
	updateExcuseByPBMService := updateExcuseByPBM.NewUpdateExcuseService(updateExcuseByPBMRepository)
	updateExcuseByPBMHandler := updateExcuseByPBMHandler.NewUpdateExcuseHandler(updateExcuseByPBMService)

	groupStudent := router.Group("api/student/:studentid/excuse", middleware.Auth(), middleware.IsStudent())
	groupStudent.GET("", getAllExcuseHandler.GetAllExcuseHandler)
	groupStudent.GET(":id", getExcuseHandler.GetExcuseHandler)
	groupStudent.POST("", createExcuseHandler.CreateExcuseHandler)
	groupStudent.PATCH(":id", updateExcuseHandler.UpdateExcuseHandler)
	groupStudent.PUT(":id", updateExcuseHandler.UpdateExcuseHandler)

	groupPBM := router.Group("api/pbm/excuse", middleware.Auth(), middleware.IsPBM())
	groupPBM.GET("", getExcuseByPBMHandler.GetExcuseHandler)
	groupPBM.GET(":id", getAllExcuseByPBMHandler.GetAllExcuseHandler)
	groupPBM.PATCH(":id", updateExcuseByPBMHandler.UpdateExcuseHandler)
	groupPBM.PUT(":id", updateExcuseByPBMHandler.UpdateExcuseHandler)
}
