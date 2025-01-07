package server

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/Cattle0Horse/url-shortener/config"
	"github.com/Cattle0Horse/url-shortener/internal/global/database"
	"github.com/Cattle0Horse/url-shortener/internal/global/jwt"
	"github.com/Cattle0Horse/url-shortener/internal/global/logger"
	"github.com/Cattle0Horse/url-shortener/internal/global/middleware"
	"github.com/Cattle0Horse/url-shortener/internal/global/redis"
	"github.com/Cattle0Horse/url-shortener/internal/module"
	"github.com/Cattle0Horse/url-shortener/pkg/tools"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var log *slog.Logger

func Init() {
	config.Init()
	jwt.Init(config.Get().JWT)

	log = logger.NewModule("Server")

	database.Init()
	redis.Init()

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
	r.Use(cors.New(cors.Config{
		AllowOrigins:     config.Get().Cors.AllowOrigins,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	r.Use(middleware.Recovery())

	for _, m := range module.Modules {
		log.Info(fmt.Sprintf("Init Router: %s", m.GetName()))
		m.InitRouter(r.Group(cfg.Prefix))
	}
	err := r.Run(cfg.Host + ":" + cfg.Port)
	tools.PanicOnErr(err)
}
