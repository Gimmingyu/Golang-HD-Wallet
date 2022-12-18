package api

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// SetRequestUUID : Sample Middleware
func SetRequestUUID() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Before Request
		uid := uuid.New()
		c.Set("uuid", uid)

		// After Request
		c.Next()
	}
}
