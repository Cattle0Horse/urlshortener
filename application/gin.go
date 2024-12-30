package application

import (
	"github.com/Cattle0Horse/url-shortener/internal/middleware"
	"github.com/gin-gonic/gin"
)

func NewGin() *gin.Engine {
	g := gin.Default()
	g.Use(middleware.Logger)
	g.Use(middleware.Cors)
	return g
}
