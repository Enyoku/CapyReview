package api

import (
	"github.com/gin-gonic/gin"
)

func headerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json; charset=utf-8")
	}
}
