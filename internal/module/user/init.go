package user

import (
	"log/slog"

	"github.com/Cattle0Horse/url-shortener/internal/global/logger"
	"github.com/Cattle0Horse/url-shortener/test"
)

var log *slog.Logger

type ModuleUser struct{}

func (u *ModuleUser) GetName() string {
	return "User"
}

func (u *ModuleUser) Init() {
	switch test.IsTest() {
	case false:
		log = logger.New("User")
	case true:
		log = logger.Get()
	}
}

func selfInit() {
	u := &ModuleUser{}
	u.Init()
}
