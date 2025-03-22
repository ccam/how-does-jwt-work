package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := ValidateJWT(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
			c.Abort()
			return
		}

		error := ValidateAdiminRoleJWT(c)
		if error != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Admins only, loser"})
			c.Abort()
			return
		}
		c.Next()
	}
}
