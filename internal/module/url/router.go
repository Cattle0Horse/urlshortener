package url

import (
	"github.com/Cattle0Horse/url-shortener/internal/global/middleware"
	"github.com/gin-gonic/gin"
)

func (u *ModuleUrl) InitRouter(r *gin.RouterGroup) {
	r.GET("/url/:code", Redirect) // 重定向到原始链接

	//添加鉴权中间件
	r.Group("", middleware.Auth()).
		GET("/urls", FetchAll).       // 获取用户的所有短链接
		POST("/url", Create).         // 创建短链接
		DELETE("/url/:code", Delete). // 删除短链接
		PATCH("/url/:code", Update)   // 更新短链接的有效期
}
