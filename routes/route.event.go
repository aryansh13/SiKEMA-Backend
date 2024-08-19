package route

import (
	getEvent "attendance-is/controllers/event/lecturer/get"
	getEventHandler "attendance-is/handlers/event/lecturer/get"

	getRecentEventLecturer "attendance-is/controllers/event/lecturer/get_recent"
	getRecentEventLecturerHandler "attendance-is/handlers/event/lecturer/get_recent"

	getRecentEventStudent "attendance-is/controllers/event/student/get_recent"
	getRecentEventStudentHandler "attendance-is/handlers/event/student/get_recent"

	showEvent "attendance-is/controllers/event/lecturer/show"
	showEventHandler "attendance-is/handlers/event/lecturer/show"

	addStudentEvent "attendance-is/controllers/event/lecturer/add_student"
	addStudentEventHandler "attendance-is/handlers/event/lecturer/add_student"

	removeStudentEvent "attendance-is/controllers/event/lecturer/remove_student"
	removeStudentEventHandler "attendance-is/handlers/event/lecturer/remove_student"

	createEvent "attendance-is/controllers/event/lecturer/create"
	createEventHandler "attendance-is/handlers/event/lecturer/create"

	updateEvent "attendance-is/controllers/event/lecturer/update"
	updateEventHandler "attendance-is/handlers/event/lecturer/update"

	finalizeEvent "attendance-is/controllers/event/lecturer/finalize"
	finalizeEventHandler "attendance-is/handlers/event/lecturer/finalize"

	getRecentQRCode "attendance-is/controllers/qrcode/lecturer/get_recent"
	getRecentQRCodeHandler "attendance-is/handlers/qrcode/lecturer/get_recent"

	saveQRCode "attendance-is/controllers/qrcode/lecturer/save"
	saveQRCodeHandler "attendance-is/handlers/qrcode/lecturer/save"

	middleware "attendance-is/middlewares"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitEventRoute(db *gorm.DB, r *gin.Engine) {
	getEventRepo := getEvent.NewGetEventRepository(db)
	getEventService := getEvent.NewGetEventService(getEventRepo)
	getEventHandler := getEventHandler.NewGetEventHandler(getEventService)

	getRecentEventLecturerRepo := getRecentEventLecturer.NewGetRecentEventRepository(db)
	getRecentEventLecturerService := getRecentEventLecturer.NewGetRecentEventService(getRecentEventLecturerRepo)
	getRecentEventLecturerHandler := getRecentEventLecturerHandler.NewGetRecentEventHandler(getRecentEventLecturerService)

	getRecentEventStudentRepo := getRecentEventStudent.NewGetRecentEventRepository(db)
	getRecentEventStudentService := getRecentEventStudent.NewGetRecentEventService(getRecentEventStudentRepo)
	getRecentEventStudentHandler := getRecentEventStudentHandler.NewGetRecentEventHandler(getRecentEventStudentService)

	showEventRepo := showEvent.NewShowEventRepository(db)
	showEventService := showEvent.NewShowEventService(showEventRepo)
	showEventHandler := showEventHandler.NewShowEventHandler(showEventService)

	addStudentEventRepo := addStudentEvent.NewAddStudentEventRepository(db)
	addStudentEventService := addStudentEvent.NewAddStudentToEventService(addStudentEventRepo)
	addStudentEventHandler := addStudentEventHandler.NewAddStudentEventHandler(addStudentEventService)

	removeStudentEventRepo := removeStudentEvent.NewRemoveStudentEventRepository(db)
	removeStudentEventService := removeStudentEvent.NewRemoveStudentToEventService(removeStudentEventRepo)
	removeStudentEventHandler := removeStudentEventHandler.NewRemoveStudentEventHandler(removeStudentEventService)

	createEventRepo := createEvent.NewCreateEventRepository(db)
	createEventService := createEvent.NewCreateEventService(createEventRepo)
	createEventHandler := createEventHandler.NewCreateEventHandler(createEventService)

	updateEventRepo := updateEvent.NewUpdateEventRepository(db)
	updateEventService := updateEvent.NewUpdateEventService(updateEventRepo)
	updateEventHandler := updateEventHandler.NewUpdateEventHandler(updateEventService)

	finalizeEventRepo := finalizeEvent.NewFinalizeEventRepository(db)
	finalizeEventService := finalizeEvent.NewFinalizeEventService(finalizeEventRepo)
	finalizeEventHandler := finalizeEventHandler.NewFinalizeEventHandler(finalizeEventService)

	getRecentQRCodeRepo := getRecentQRCode.NewGetRecentQRCodeRepository(db)
	getRecentQRCodeService := getRecentQRCode.NewGetRecentQRCodeService(getRecentQRCodeRepo)
	getRecentQRCodeHandler := getRecentQRCodeHandler.NewGetRecentQRCodeHandler(getRecentQRCodeService)

	saveQRCodeRepo := saveQRCode.NewSaveQRCodeRepository(db)
	saveQRCodeService := saveQRCode.NewSaveQRCodeService(saveQRCodeRepo)
	saveQRCodeHandler := saveQRCodeHandler.NewSaveQRCodeHandler(saveQRCodeService)

	groupLecturer := r.Group("/api/lecturer/:lecturerid/event", middleware.Auth(), middleware.IsLecturer())

	groupLecturer.GET("", getEventHandler.GetEventHandler)
	groupLecturer.GET(":id/qrcode", getRecentQRCodeHandler.GetRecentQRCodeHandler)
	groupLecturer.POST(":id/qrcode", saveQRCodeHandler.SaveQRCodeHandler)
	groupLecturer.GET("/recent", getRecentEventLecturerHandler.GetRecentEventHandler)
	groupLecturer.GET(":id", showEventHandler.ShowEventHandler)
	groupLecturer.PATCH(":id", updateEventHandler.UpdateEventHandler)
	groupLecturer.PATCH(":id/finalize", finalizeEventHandler.FinalizeEventHandler)
	groupLecturer.POST("", createEventHandler.CreateEventHandler)
	groupLecturer.POST(":id/student", addStudentEventHandler.AddStudentEventHandler)
	groupLecturer.DELETE(":id/student", removeStudentEventHandler.RemoveStudentEvent)

	groupStudent := r.Group("/api/student/:studentid/event", middleware.Auth(), middleware.IsStudent())
	// groupStudent.GET("")
	groupStudent.GET("/recent", getRecentEventStudentHandler.GetRecentEventHandler)
}
