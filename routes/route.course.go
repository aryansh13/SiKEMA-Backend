package route

import (
	middleware "attendance-is/middlewares"

	getCourseStudent "attendance-is/controllers/course/student/get"
	getAllCourseStudent "attendance-is/controllers/course/student/getAll"
	getCourseHandler "attendance-is/handlers/course/student/get"
	getAllCourseHandler "attendance-is/handlers/course/student/getAll"

	getCourseAttendanceStudent "attendance-is/controllers/course/student/get_attendance"
	getCourseAttendanceHandler "attendance-is/handlers/course/student/get_attendance"

	getCourseLecturer "attendance-is/controllers/course/lecturer/get"
	getCourseLecturerHandler "attendance-is/handlers/course/lecturer/get"

	getCoursePBM "attendance-is/controllers/course/pbm/get"
	getCoursePBMHandler "attendance-is/handlers/course/pbm/get"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitCourseRoute(db *gorm.DB, r *gin.Engine) {

	getCourseRepository := getCourseStudent.NewGetCourseRepository(db)
	getCourseService := getCourseStudent.NewGetCourseService(getCourseRepository)
	getCourseHandler := getCourseHandler.NewGetCourseHandler(getCourseService)

	getCourseAttendanceRepository := getCourseAttendanceStudent.NewGetCourseAttendanceRepository(db)
	getCourseAttendanceService := getCourseAttendanceStudent.NewGetCourseAttendanceService(getCourseAttendanceRepository)
	getCourseAttendanceHandler := getCourseAttendanceHandler.NewGetCourseAttendanceHandler(getCourseAttendanceService)

	getAllCourseRepository := getAllCourseStudent.NewGetAllCourseRepository(db)
	getAllCourseService := getAllCourseStudent.NewGetAllCourseService(getAllCourseRepository)
	getAllCourseHandler := getAllCourseHandler.NewGetAllCourseHandler(getAllCourseService)

	getCourseLecturerRepository := getCourseLecturer.NewGetCourseRepository(db)
	getCourseLecturerService := getCourseLecturer.NewGetCourseService(getCourseLecturerRepository)
	getCourseLecturerHandler := getCourseLecturerHandler.NewGetCourseHandler(getCourseLecturerService)

	getCoursePBMRepository := getCoursePBM.NewGetCourseRepository(db)
	getCoursePBMService := getCoursePBM.NewGetCourseService(getCoursePBMRepository)
	getCoursePBMHandler := getCoursePBMHandler.NewGetCourseHandler(getCoursePBMService)

	groupStudent := r.Group("api/student/:studentid/course", middleware.Auth(), middleware.IsStudent())
	groupStudent.GET("", getAllCourseHandler.GetAllCourseHandler)
	groupStudent.GET(":id/attendance", getCourseAttendanceHandler.GetCourseAttendanceHandler)
	groupStudent.GET(":id", getCourseHandler.GetCourseHandler)

	groupLecturer := r.Group("api/lecturer/:lecturerid/course", middleware.Auth(), middleware.IsLecturer())
	groupLecturer.GET("", getCourseLecturerHandler.GetCourseHandler)

	groupPBM := r.Group("api/pbm/course", middleware.Auth(), middleware.IsPBM())
	groupPBM.GET("", getCoursePBMHandler.GetCourseHandler)
}
