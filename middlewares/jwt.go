package middleware

import (
	util "attendance-is/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		token := util.ExtractToken(c)
		if token == "" || util.ValidateToken(token) != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthenticated",
			})
			defer c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		claim, _ := util.ExtractTokenClaim(token)
		c.Set("user", claim)
		c.Next()
	})
}

func checkClaim(c *gin.Context, allowedTypes []string, paramID string) bool {
	val, exist := c.Get("user")
	claim, _ := val.(util.JWTClaim)

	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthenticated",
		})
		c.AbortWithStatus(http.StatusUnauthorized)
		return false
	}

	for _, allowedType := range allowedTypes {
		if claim.Type == allowedType {
			if claim.TypeID == paramID {
				return true
			} else if paramID == "" {
				return true
			}
		}

	}

	c.JSON(http.StatusForbidden, gin.H{
		"message": "Unauthorized",
	})
	c.AbortWithStatus(http.StatusForbidden)
	return false
}

func IsStudent() gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		if checkClaim(c, []string{"student", "0"}, c.Param("studentid")) {
			c.Next()
		}
	})
}

func IsLecturer() gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		if checkClaim(c, []string{"lecturer", "1"}, c.Param("lecturerid")) {
			c.Next()
		}
	})
}

func IsPBM() gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		if checkClaim(c, []string{"pbm", "2"}, "") {
			c.Next()
		}
	})
}
