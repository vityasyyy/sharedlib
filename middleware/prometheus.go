package middleware

import (
	"fmt"
	"time"

	"github.com/vityasyyy/sharedlib/metrics"

	"github.com/gin-gonic/gin"
)

func PrometheusMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		status := fmt.Sprint(c.Writer.Status())
		path := c.FullPath()
		method := c.Request.Method
		duration := time.Since(start)
		metrics.HTTPReqs.WithLabelValues(method, path, status).Inc()
		metrics.HTTPDur.WithLabelValues(method, path).Observe(duration.Seconds())
	}
}
