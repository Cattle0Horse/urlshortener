package url

import (
	"log/slog"

	"github.com/Cattle0Horse/url-shortener/config"
	"github.com/Cattle0Horse/url-shortener/internal/global/logger"
	"github.com/Cattle0Horse/url-shortener/test"
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
