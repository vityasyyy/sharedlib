package middleware

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
)

// Timeout adds a timeout to all HTTP requests
func Timeout(duration time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), duration)
		defer cancel()

		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
