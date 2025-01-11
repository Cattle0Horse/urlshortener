// Provide init function and global variables for user module,
// it will be called by `cmd/server/server.go`
package user

import (
	"log/slog"

	"github.com/Cattle0Horse/url-shortener/internal/global/logger"
)

var log *slog.Logger

type ModuleUser struct{}

func (u *ModuleUser) GetName() string {
	return "User"
}

func (u *ModuleUser) Init() {
	log = logger.NewModule("User")
}

func selfInit() {
	u := &ModuleUser{}
	u.Init()
}
