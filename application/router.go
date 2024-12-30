package application

import (
	"log"

	"github.com/Cattle0Horse/url-shortener/config"
	"github.com/Cattle0Horse/url-shortener/internal/cache"
	"github.com/Cattle0Horse/url-shortener/internal/controller"
	"github.com/Cattle0Horse/url-shortener/internal/middleware"
	"github.com/Cattle0Horse/url-shortener/internal/repo"
	"github.com/Cattle0Horse/url-shortener/internal/service"
	"github.com/Cattle0Horse/url-shortener/pkg/emailsender"
	"github.com/Cattle0Horse/url-shortener/pkg/hasher"
	"github.com/Cattle0Horse/url-shortener/pkg/jwt"
	"github.com/Cattle0Horse/url-shortener/pkg/randnum"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func setupRouter(gin *gin.Engine, config *config.Config, db *gorm.DB) error {
	pemailsender, err := newEmailSender(config.Email)
	if err != nil {
		return err
	}
	log.Println("Email sender initialized")
	pjwt := newJWT(config.JWT)
	log.Println("JWT initialized")
	predisCache, err := newUserCacher(config.Redis)
	if err != nil {
		return err
	}
	log.Println("Redis cache initialized")

	// 公开的API
	publicRouter := gin.Group("")
	// 需要JWT认证的API
	protectedRouter := gin.Group("/api", middleware.JwtAuthMiddleware(pjwt))

	// url缩短服务相关路由
	urlRepo := repo.NewUrlRepository(db)
	urlService := service.NewUrlService(urlRepo, config.App.DefaultDuration, config.App.BaseUrl)
	urlController := controller.NewUrlController(urlService)
	publicRouter.GET("/:code", urlController.Redirect) // 短链接重定向
	protectedRouter.
		POST("/url", urlController.Create).                   // 创建短链接
		GET("/url", urlController.FetchAll).                  // 获取用户的所有短链接
		DELETE("/url/:code", urlController.Delete).           // 删除短链接
		PATCH("/url/:code", urlController.UpdateByExpiryTime) // 更新短链接的有效期

	// 用户认证相关路由
	userRepo := repo.NewUserRepository(db)
	userService := service.NewUserService(userRepo, newPasswordHash(), pjwt, predisCache, pemailsender, newRandNum())
	userController := controller.NewUserController(userService)
	publicRouter.
		POST("/api/auth/login", userController.Login).                 // 用户登录
		POST("/api/auth/register", userController.Register).           // 用户注册
		POST("/api/auth/forget", userController.ResetPassword).        // 忘记密码
		GET("/api/auth/register/:email", userController.SendEmailCode) // 发送注册验证码

	return nil
}

func newEmailSender(cfg *config.EmailConfig) (*emailsender.EmailSend, error) {
	return emailsender.NewEmailSend(
		cfg.HostAddress,
		cfg.HostPort,
		cfg.Username,
		cfg.Password,
		cfg.Subject,
		cfg.TestMail)
}

func newJWT(cfg *config.JWTConfig) *jwt.JWT {
	return jwt.NewJWT([]byte(cfg.AccessTokenSecret), cfg.AccessTokenDuration)
}

func newUserCacher(cfg *config.RedisConfig) (*cache.RedisCache, error) {
	return cache.NewRedisCache(cfg.Address, cfg.Password, cfg.DB, cfg.UrlDuration, cfg.EmailCodeDuration)
}

func newRandNum() *randnum.RandNum {
	// 用于生成6位email验证码
	return randnum.NewRandNum(6)
}

func newPasswordHash() *hasher.PasswordHash {
	return hasher.NewPasswordHash()
}
