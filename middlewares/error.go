package middleware

import (
	util "attendance-is/utils"
	"fmt"

	"github.com/gin-gonic/gin"
)

type CustomError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (ce *CustomError) Error() string {
	return fmt.Sprintf("Error %d: %s", ce.Code, ce.Message)
}

func NewCustomError(code int, message string) *CustomError {
	return &CustomError{
		Code:    code,
		Message: message,
	}
}

func GlobalErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		err := c.Errors.Last()
		if err != nil {
			switch e := err.Err.(type) {
			case *CustomError:
				util.ErrorResponse(c, e.Code, e.Message)
			default:
				util.ErrorResponse(c, 500, err.Error())
			}

			c.Abort()
		}
	}
}
