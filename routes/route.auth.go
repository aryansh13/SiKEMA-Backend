package route

import (
	loginAuth "attendance-is/controllers/auth/login"
	validateAuth "attendance-is/controllers/auth/validate"
	loginAuthHandler "attendance-is/handlers/auth/login"
	validateAuthHandler "attendance-is/handlers/auth/validate"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitAuthRoute(db *gorm.DB, r *gin.Engine) {
	loginAuthRepository := loginAuth.NewLoginAuthRepository(db)
	loginAuthService := loginAuth.NewLoginAuthService(loginAuthRepository)
	loginAuthHandler := loginAuthHandler.NewLoginAuthHandler(loginAuthService)

	validateAuthRepository := validateAuth.NewValidateAuthRepository(db)
	validateAuthService := validateAuth.NewValidateAuthService(validateAuthRepository)
	validateAuthHandler := validateAuthHandler.NewValidateAuthHandler(validateAuthService)

	group := r.Group("api/auth")
	group.POST("login", loginAuthHandler.LoginAuthHandler)
	group.GET("validate", validateAuthHandler.ValidateAuthHandler)
}
