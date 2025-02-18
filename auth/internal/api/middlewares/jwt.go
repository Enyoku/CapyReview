package middleware

import (
	"authService/internal/auth"
	"authService/internal/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RequiredRole(config *config.JWT, allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenCookie, err := c.Request.Cookie("token")
		if err != nil || tokenCookie == nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "missing or invalid token"})
			c.Abort()
			return
		}

		claims, err := auth.ParseToken(tokenCookie.Value, *config)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}

		roleMatch := false
		for _, role := range allowedRoles {
			if claims.Role == role {
				roleMatch = true
				break
			}
		}

		if !roleMatch {
			c.JSON(http.StatusForbidden, gin.H{"error": "insufficient permissions"})
			c.Abort()
			return
		}

		c.Set("user_id", claims.UserId)
		c.Set("role", claims.Role)

		c.Next()
	}
}
