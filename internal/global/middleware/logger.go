package middleware

import (
	"github.com/Cattle0Horse/url-shortener/internal/global/otel"
	"github.com/gin-gonic/gin"
	"log/slog"
	"time"
)

func Logger(log *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()

		// 处理请求
		c.Next()

		// 结束时间
		endTime := time.Now()
		latency := endTime.Sub(startTime)

		// 记录监控指标
		otel.RecordRequestMetrics(
			c.Request.Method,
			c.Request.URL.Path,
			c.Writer.Status(),
			latency,
		)

		// 记录请求日志
		log.Info("Request",
			"method", c.Request.Method,
			"path", c.Request.URL.Path,
			"status", c.Writer.Status(),
			"latency", latency,
			"client_ip", c.ClientIP(),
		)

		// 如果状态码 >= 500，记录错误日志
		if c.Writer.Status() >= 500 {
			log.Error("Server error",
				"method", c.Request.Method,
				"path", c.Request.URL.Path,
				"status", c.Writer.Status(),
				"latency", latency,
				"client_ip", c.ClientIP(),
			)
		}
	}
}
