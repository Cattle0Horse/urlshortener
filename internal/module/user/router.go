package user

import (
	"github.com/gin-gonic/gin"
)

func (u *ModuleUser) InitRouter(r *gin.RouterGroup) {

	r.POST("/api/auth/login", Login).
		POST("/api/auth/register", Create)

}
