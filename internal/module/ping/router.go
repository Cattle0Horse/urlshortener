package ping

import (
	"github.com/gin-gonic/gin"
)

func (p *ModulePing) InitRouter(r *gin.RouterGroup) {
	r.GET("/ping", Ping)
}
