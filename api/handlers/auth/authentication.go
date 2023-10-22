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
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}
		c.Next()
	}
}
