package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vityasyyy/sharedlib/logger"
)

func ReqLoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		requestID := c.GetString("request_id")

		// Create a request-scoped logger
		reqLogger := logger.Log.With().
			Str("request_id", requestID).
			Str("method", c.Request.Method).
			Str("path", c.Request.URL.Path).
			Logger()

		logger.AttachLogger(c, reqLogger)

		c.Next()

		duration := time.Since(start)
		reqLogger.Info().
			Int("status", c.Writer.Status()).
			Dur("duration", duration).
			Str("ip_address", c.ClientIP()).
			Msg("Request completed")
	}
}
