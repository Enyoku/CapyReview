package middleware

import (
	"github.com/gin-gonic/gin"
)

func HeaderMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json; charset=utf-8")
	}
}
