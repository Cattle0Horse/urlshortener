package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

// RateLimiter 限流中间件（ token bucket 算法）
// r: 令牌桶的容量，即每秒生成令牌的数量
// b: 令牌桶的填充速度，即每秒向桶中放入令牌的数量
func RateLimiter(r rate.Limit, b int) gin.HandlerFunc {
	limiter := rate.NewLimiter(r, b)

	return func(c *gin.Context) {
		if !limiter.Allow() {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"code":    http.StatusTooManyRequests,
				"message": "too many requests",
			})
			return
		}
		c.Next()
	}
}

// AdaptiveRateLimiter 自适应限流
func AdaptiveRateLimiter(maxRPS int) gin.HandlerFunc {
	var (
		limiter    *rate.Limiter
		lastUpdate time.Time
	)

	return func(c *gin.Context) {
		// 根据当前负载动态调整限流值
		if time.Since(lastUpdate) > time.Minute {
			limiter = rate.NewLimiter(rate.Limit(maxRPS), maxRPS*2)
			lastUpdate = time.Now()
		}

		if !limiter.Allow() {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"code":    http.StatusTooManyRequests,
				"message": "too many requests",
			})
			return
		}
		c.Next()
	}
}
