package middleware

import (
	"strings"

	"github.com/Cattle0Horse/url-shortener/internal/global/errs"
	"github.com/Cattle0Horse/url-shortener/internal/global/jwt"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			errs.Fail(c, errs.EmptyAuthorization)
			c.Abort()
			return
		}
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			errs.Fail(c, errs.ErrTokenInvalid)
			c.Abort()
			return
		}
		if payload, valid := jwt.ParseToken(parts[1]); !valid {
			errs.Fail(c, errs.ErrTokenInvalid)
			c.Abort()
			return
		} else {
			c.Set("payload", payload)
		}
		c.Next()
	}
}
