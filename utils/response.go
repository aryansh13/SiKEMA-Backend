package util

import "github.com/gin-gonic/gin"

type response struct {
	Meta interface{} `json:"meta,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

func APIResponse(c *gin.Context, StatusCode int, Data interface{}, Meta interface{}) {
	response := response{
		Meta: Meta,
		Data: Data,
	}

	if StatusCode >= 400 {
		defer c.AbortWithStatusJSON(StatusCode, response)
	} else {
		c.JSON(StatusCode, response)
	}
}

func ErrorResponse(c *gin.Context, StatusCode int, Msg string) {
	defer c.AbortWithStatusJSON(StatusCode, gin.H{"message": Msg})
}
