package authentication

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// JwtAuthMiddleware is a middleware that checks the validity of a JWT token.
func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := TokenValid(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "unauthorized",
			})
			return
		}
		c.Next()
	}
}
