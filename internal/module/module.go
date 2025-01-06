package module

import (
	"github.com/Cattle0Horse/url-shortener/internal/module/ping"
	"github.com/Cattle0Horse/url-shortener/internal/module/url"
	"github.com/Cattle0Horse/url-shortener/internal/module/user"
	"github.com/gin-gonic/gin"
)

type Module interface {
	GetName() string
	Init()
	InitRouter(r *gin.RouterGroup)
}

var Modules []Module

func registerModule(m []Module) {
	Modules = append(Modules, m...)
}

func init() {
	// Register your module here
	registerModule([]Module{
		&user.ModuleUser{},
		&ping.ModulePing{},
		&url.ModuleUrl{},
	})
}
