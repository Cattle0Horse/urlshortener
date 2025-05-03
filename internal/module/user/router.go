package user

import (
	"github.com/gin-gonic/gin"
)

func (u *ModuleUser) InitRouter(r *gin.RouterGroup) {

	r.POST("/auth/login", Login).
		POST("/auth/register", Create)
	// GET("/auth/me", Me)

}
