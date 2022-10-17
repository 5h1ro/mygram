package middleware

import (
	"mygram/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		verify, err := helpers.Verify(c)

		if err != nil {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized, gin.H{
					"error":   "Unauthenticated",
					"message": err.Error(),
				},
			)
			return
		}

		c.Set("userData", verify)

		c.Next()
	}
}
