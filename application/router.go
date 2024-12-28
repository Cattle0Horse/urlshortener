package application

import (
	"github.com/Cattle0Horse/url-shortener/config"
	"github.com/Cattle0Horse/url-shortener/internal/controller"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func setupRouter(gin *gin.Engine, config *config.Config, db *gorm.DB) {
	// 公开的API
	publicRouter := gin.Group("")
	// 需要JWT认证的API
	protectedRouter := gin.Group("/api")

	// url缩短服务相关路由
	uc := controller.NewUrlController(config, db)
	publicRouter.GET("/:code", uc.Redirect) // 短链接重定向
	protectedRouter.
		POST("/url", uc.Create).                   // 创建短链接
		GET("/url", uc.FetchAll).                  // 获取用户的所有短链接
		DELETE("/url/:code", uc.Delete).           // 删除短链接
		PATCH("/url/:code", uc.UpdateByExpiryTime) // 更新短链接的有效期
}
