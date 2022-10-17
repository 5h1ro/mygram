package middleware

import (
	"mygram/database"
	"mygram/user"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func IsExistMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, db := database.GetDB()
		userData := c.MustGet("userData").(jwt.MapClaims)
		User := user.User{}

		err := db.Select("id").First(&User, int(userData["id"].(float64))).Error
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "failed",
				"error":  "Token invalid",
			})
			return
		}

		c.Next()
	}
}
