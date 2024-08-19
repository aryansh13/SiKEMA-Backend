package route

import (
	createSP "attendance-is/controllers/sp/pbm/create"
	deleteSP "attendance-is/controllers/sp/pbm/delete"
	getSP "attendance-is/controllers/sp/pbm/get"
	getAllSP "attendance-is/controllers/sp/pbm/getAll"
	updateSP "attendance-is/controllers/sp/pbm/update"
	createSPHandler "attendance-is/handlers/sp/pbm/create"
	deleteSPHandler "attendance-is/handlers/sp/pbm/delete"
	getSPHandler "attendance-is/handlers/sp/pbm/get"
	getAllSPHandler "attendance-is/handlers/sp/pbm/getAll"
	updateSPHandler "attendance-is/handlers/sp/pbm/update"
	middleware "attendance-is/middlewares"

	getSPStudent "attendance-is/controllers/sp/student/get"
	getAllSPStudent "attendance-is/controllers/sp/student/getAll"
	getSPHandlerStudent "attendance-is/handlers/sp/student/get"
	getAllSPHandlerStudent "attendance-is/handlers/sp/student/getAll"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitSPRoute(db *gorm.DB, r *gin.Engine) {
	getSPPBMRepository := getSP.NewGetSPRepository(db)
	getSPPBMService := getSP.NewGetSPService(getSPPBMRepository)
	getSPPBMHandler := getSPHandler.NewGetSPHandler(getSPPBMService)

	getAllSPPBMRepository := getAllSP.NewGetAllSPRepository(db)
	getAllSPPBMService := getAllSP.NewGetAllSPService(getAllSPPBMRepository)
	getAllSPPBMHandler := getAllSPHandler.NewGetAllSPHandler(getAllSPPBMService)

	createSPPBMRepository := createSP.NewCreateSPRepository(db)
	createSPPBMService := createSP.NewCreateSPService(createSPPBMRepository)
	createSPPBMHandler := createSPHandler.NewCreateSPHandler(createSPPBMService)

	updateSPPBMRepository := updateSP.NewUpdateSPRepository(db)
	updateSPPBMService := updateSP.NewUpdateSPService(updateSPPBMRepository)
	updateSPPBMHandler := updateSPHandler.NewUpdateSPHandler(updateSPPBMService)

	deleteSPPBMRepository := deleteSP.NewDeleteSPRepository(db)
	deleteSPPBMService := deleteSP.NewDeleteSPService(deleteSPPBMRepository)
	deleteSPPBMHandler := deleteSPHandler.NewDeleteSPHandler(deleteSPPBMService)

	getSPStudentRepository := getSPStudent.NewGetSPRepository(db)
	getSPStudentService := getSPStudent.NewGetSPService(getSPStudentRepository)
	getSPStudentHandler := getSPHandlerStudent.NewGetSPHandler(getSPStudentService)

	getAllSPStudentRepository := getAllSPStudent.NewGetAllSPRepository(db)
	getAllSPStudentService := getAllSPStudent.NewGetAllSPService(getAllSPStudentRepository)
	getAllSPStudentHandler := getAllSPHandlerStudent.NewGetAllSPHandler(getAllSPStudentService)

	groupPBM := r.Group("api/pbm/sp", middleware.Auth(), middleware.IsPBM())
	groupPBM.GET("", getSPPBMHandler.GetSPHandler)
	groupPBM.GET(":id", getAllSPPBMHandler.GetAllSPHandler)
	groupPBM.PATCH(":id", updateSPPBMHandler.UpdateSPHandler)
	groupPBM.POST("", createSPPBMHandler.CreateSPHandler)
	groupPBM.PUT(":id", updateSPPBMHandler.UpdateSPHandler)
	groupPBM.DELETE(":id", deleteSPPBMHandler.DeleteSPHandler)

	groupStudent := r.Group("api/student/:studentid/sp", middleware.Auth(), middleware.IsStudent())
	groupStudent.GET("", getSPStudentHandler.GetSPHandler)
	groupStudent.GET(":id", getAllSPStudentHandler.GetAllSPHandler)
}
