package application

import (
	"github.com/Cattle0Horse/url-shortener/config"
	"github.com/Cattle0Horse/url-shortener/database"
	"github.com/Cattle0Horse/url-shortener/pkg/logger"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type App struct {
	config *config.Config
	mysql  *gorm.DB
	gin    *gin.Engine
}

func NewApp(configFilePath string) (*App, error) {
	config, err := config.NewConfig(configFilePath)
	if err != nil {
		return nil, err
	}

	logger.InitLogger(&config.Logger)

	mysql, err := database.NewMysql(&config.Database)
	if err != nil {
		return nil, err
	}

	gin := NewGin()
	setupRouter(gin, config, mysql)

	return &App{
		config: config,
		mysql:  mysql,
		gin:    gin,
	}, nil
}

func (a *App) Start() {
	a.gin.Run(":" + a.config.App.Port)
}
