package server

import (
	"fmt"
	"log/slog"

	"github.com/Cattle0Horse/url-shortener/config"
	"github.com/Cattle0Horse/url-shortener/internal/global/database/mysql"
	"github.com/Cattle0Horse/url-shortener/internal/global/jwt"
	"github.com/Cattle0Horse/url-shortener/internal/global/logger"
	"github.com/Cattle0Horse/url-shortener/internal/global/middleware"
	"github.com/Cattle0Horse/url-shortener/internal/module"
	"github.com/Cattle0Horse/url-shortener/tools"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var log *slog.Logger

func Init() {
	config.Init()
	jwt.Init(config.Get().JWT)

	log = logger.New("Server")

	mysql.Init()

	for _, m := range module.Modules {
		log.Info(fmt.Sprintf("Init Module: %s", m.GetName()))
		m.Init()
	}
}

func Run() {
	cfg := config.Get().Server
	gin.SetMode(string(cfg.Mode))
	r := gin.New()

	switch cfg.Mode {
	case config.ModeRelease:
		r.Use(middleware.Logger(logger.Get()))
	case config.ModeDebug:
		r.Use(gin.Logger())
	}

	// 跨域
	r.Use(cors.Default())
	r.Use(middleware.Recovery())

	for _, m := range module.Modules {
		log.Info(fmt.Sprintf("Init Router: %s", m.GetName()))
		m.InitRouter(r.Group(cfg.Prefix))
	}
	err := r.Run(cfg.Host + ":" + cfg.Port)
	tools.PanicOnErr(err)
}
