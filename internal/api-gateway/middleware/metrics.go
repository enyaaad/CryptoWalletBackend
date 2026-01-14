package middleware

import (
	"strconv"
	"time"

	"github.com/enyaaad/CryptoWalletBackend/pkg/metrics"
	"github.com/gin-gonic/gin"
)

func MetricsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		duration := time.Since(start).Seconds()
		status := strconv.Itoa(c.Writer.Status())
		method := c.Request.Method
		endpoint := c.FullPath()

		if endpoint == "" {
			endpoint = c.Request.URL.Path
		}

		metrics.HTTPRequestsTotal.WithLabelValues(method, endpoint, status).Inc()
		metrics.HTTPRequestsDuration.WithLabelValues(method, endpoint).Observe(duration)
	}
}
