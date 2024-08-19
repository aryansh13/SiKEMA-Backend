package util

import (
	model "attendance-is/models"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connection(user string, pass string, host string, port string, dbname string) error {
	dsn := user + ":" + pass + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local"

	if db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}); err == nil {
		DB = db
		return err
	} else {
		return nil
	}
}

func Migrate() {
	DB.Debug().AutoMigrate(
		&model.User{},
		&model.Student{},
		&model.Class{},
		&model.Course{},
		&model.Lecturer{},
		&model.Event{},
		&model.Absent{},
		&model.Excuse{},
		&model.Kompen{},
		&model.SP{},
		&model.PBM{},
		// &model.Enrollment{},
	)
}

func HandleGORMError(c *gin.Context, err error, status int, message string) bool {
	if err != nil {
		c.JSON(status, gin.H{
			"message": message,
			"error":   err.Error(),
		})
		return true
	}
	return false
}
