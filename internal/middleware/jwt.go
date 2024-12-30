package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/Cattle0Horse/url-shortener/internal/schema"
	"github.com/Cattle0Horse/url-shortener/pkg/jwt"
	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware(jwt *jwt.JWT) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		t := strings.Split(authHeader, " ")
		if len(t) != 2 || t[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, schema.ErrorResponse{Message: "Not authorized"})
			c.Abort()
			return
		}

		authToken := t[1]
		claims, err := jwt.ParseToken(authToken)
		if err != nil {
			c.JSON(http.StatusUnauthorized, schema.ErrorResponse{Message: err.Error()})
			c.Abort()
			return
		}
		c.Set("userID", claims.UserID)
		fmt.Printf("claims.UserID: %v\n", claims.UserID)
		c.Set("email", claims.Email)
		c.Next()
	}
}
