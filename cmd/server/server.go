package server

import (
	"fmt"
	"log/slog"

	"github.com/Cattle0Horse/url-shortener/config"
	"github.com/Cattle0Horse/url-shortener/internal/global/database"
	"github.com/Cattle0Horse/url-shortener/internal/global/logger"
	"github.com/Cattle0Horse/url-shortener/internal/global/middleware"
	"github.com/Cattle0Horse/url-shortener/internal/global/otel"
	"github.com/Cattle0Horse/url-shortener/internal/module"
	"github.com/Cattle0Horse/url-shortener/tools"
	"github.com/gin-gonic/gin"
)

var log *slog.Logger

func Init() {
	config.Init()
	log = logger.New("Server")

	database.Init()

	if config.Get().OTel.Enable {
		otel.Init()
	}

	for _, m := range module.Modules {
		log.Info(fmt.Sprintf("Init Module: %s", m.GetName()))
		m.Init()
	}
}

func Run() {
	gin.SetMode(string(config.Get().Mode))
	r := gin.New()

	switch config.Get().Mode {
	case config.ModeRelease:
		r.Use(middleware.Logger(logger.Get()))
	case config.ModeDebug:
		r.Use(gin.Logger())
	}

	r.Use(middleware.Recovery())

	if config.Get().OTel.Enable {
		r.Use(middleware.Trace())
	}

	for _, m := range module.Modules {
		log.Info(fmt.Sprintf("Init Router: %s", m.GetName()))
		m.InitRouter(r.Group(config.Get().Prefix))
	}
	err := r.Run(config.Get().Host + ":" + config.Get().Port)
	tools.PanicOnErr(err)
}
