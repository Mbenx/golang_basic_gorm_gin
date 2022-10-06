package middlewares

import (
	"golang_basic_gorm_gin/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenSting := c.GetHeader("Authorization")
		if tokenSting == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Request need access token",
				"status":  "Unauthorized",
			})

			c.Abort()
			return
		}
		err := auth.ValidateToken(tokenSting)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": err.Error(),
				"status":  "Unauthorized",
			})

			c.Abort()
			return
		}
		c.Next()
	}
}
