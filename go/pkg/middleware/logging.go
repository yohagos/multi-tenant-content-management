package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yohagos/multi-content-management/pkg/logger"
	"go.uber.org/zap"
)

func Logging(log *logger.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery

		c.Next()

		latency := time.Since(start)
		status := c.Writer.Status()
		clientIP := c.ClientIP()
		method := c.Request.Method
		userAgent := c.Request.UserAgent()

		if query != "" {
			path = path + "?" + query
		}

		fields := []zap.Field{
			zap.Int("status", status),
			zap.String("method", method),
			zap.String("path", path),
			zap.String("ip", clientIP),
			zap.Duration("latency", latency),
			zap.String("user_agent", userAgent),
		}

		if len(c.Errors) > 0 {
			fields = append(fields, zap.String("errors", c.Errors.String()))
			log.Error("Request failed", fields...)
			return
		}

		if status >= 500 {
			log.Error("Server error", fields...)
		} else if status >= 400 {
			log.Warn("Client error", fields...)
		} else {
			log.Info("Request completed", fields...)
		}
	}
}