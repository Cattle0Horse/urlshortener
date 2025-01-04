package url

import (
	"log/slog"

	"github.com/Cattle0Horse/url-shortener/config"
	"github.com/Cattle0Horse/url-shortener/internal/global/logger"
	"github.com/Cattle0Horse/url-shortener/internal/global/middleware"
	"github.com/Cattle0Horse/url-shortener/test"
	"github.com/gin-gonic/gin"
)

var (
	log     *slog.Logger
	baseUrl string
)

type ModuleUrl struct{}

func (p *ModuleUrl) GetName() string {
	return "Url"
}

func (p *ModuleUrl) Init() {
	switch test.IsTest() {
	case false:
		log = logger.New("Url")
	case true:
		log = logger.Get()
	}
	// http协议
	if config.Get().Port == "8080" {
		baseUrl = "http://" + config.Get().Host + config.Get().Prefix
	} else {
		baseUrl = "http://" + config.Get().Host + ":" + config.Get().Port + config.Get().Prefix
	}
}

func (p *ModuleUrl) RegisterRoutes(router *gin.Engine) {
	// 应用限流中间件
	router.Use(middleware.AdaptiveRateLimiter(1000)) // 初始1000 RPS

	urlGroup := router.Group("/url")
	{
		urlGroup.POST("/shorten", p.shortenUrl)
		urlGroup.GET("/:code", p.redirectUrl)
	}
}

func (p *ModuleUrl) shortenUrl(c *gin.Context) {
	// 短链生成逻辑
}

func (p *ModuleUrl) redirectUrl(c *gin.Context) {
	// 长链重定向逻辑
}
