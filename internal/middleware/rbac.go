package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RequireScope(required string) gin.HandlerFunc {
	return func(c *gin.Context) {
		claims := c.MustGet("claims").(*Claims)

		for _, s := range claims.Scope {
			if s == required {
				c.Next()
				return
			}
		}

		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"error": "insufficient permissions",
		})
	}
}
