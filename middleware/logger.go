package middleware

import (
	"time"

	"github.com/Cattle0Horse/url-shortener/pkg/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Logger(c *gin.Context) {
	start := time.Now()

	c.Next()

	req := c.Request

	res := c.Writer

	fields := []zap.Field{
		zap.String("remote_ip", c.ClientIP()),
		zap.String("latency", time.Since(start).String()),
		zap.String("host", req.Host),
		zap.String("request", req.Method+" "+req.RequestURI),
		zap.Int("status", res.Status()),
		zap.Int("size", res.Size()),
		zap.String("user_agent", req.UserAgent()),
	}

	id := req.Header.Get("X-Request-ID")

	if id != "" {
		fields = append(fields, zap.String("request_id", id))
	}

	n := res.Status()
	switch {
	case n >= 500:
		logger.Error("Server error", fields...)
	case n >= 400:
		logger.Warn("Client error", fields...)
	case n >= 300:
		logger.Info("Redirection", fields...)
	default:
		logger.Info("Success", fields...)
	}
}
