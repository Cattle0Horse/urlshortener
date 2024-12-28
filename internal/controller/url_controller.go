package controller

import (
	"net/http"

	"github.com/Cattle0Horse/url-shortener/config"
	"github.com/Cattle0Horse/url-shortener/internal/domain"
	"github.com/Cattle0Horse/url-shortener/internal/repository"
	"github.com/Cattle0Horse/url-shortener/internal/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UrlController struct {
	UrlService domain.UrlService
}

func NewUrlController(config *config.Config, db *gorm.DB) *UrlController {
	ur := repository.NewUrlRepository(db)
	uc := &UrlController{
		UrlService: service.NewUrlService(ur, config.App.ContextTimeout, config.App.DefaultDuration, config.App.ServerAddress),
	}
	return uc
}

// 创建一个短链接
// POST /api/url
func (uc *UrlController) Create(c *gin.Context) {
	var json domain.CreateUrlRequest
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Messgae: err.Error()})
		return
	}
	userID := c.GetHeader("x-user-id") // 获取请求头中的x-user-id字段

	shortUrl, err := uc.UrlService.Create(c, json.OriginalUrl, json.Duration, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Messgae: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, domain.CreateUrlResponse{
		ShortUrl: shortUrl,
	})
}

// 重定向
// GET /:code
func (uc *UrlController) Redirect(c *gin.Context) {
	shortCode := c.Param("code")

	originalUrl, err := uc.UrlService.FetchOriginalUrl(c, shortCode)
	if err != nil {
		// 404 错误
		c.JSON(http.StatusNotFound, domain.ErrorResponse{Messgae: err.Error()})
		return
	}
	// 永久重定向
	c.Redirect(http.StatusPermanentRedirect, originalUrl)
}

// 获取用户所有短链接
// GET /api/url?page=1&size=10
func (uc *UrlController) FetchAll(c *gin.Context) {
	userID := c.GetHeader("x-user-id") // 获取请求头中的x-user-id字段
	var q domain.FetchAllRequest
	if err := c.ShouldBindQuery(&q); err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Messgae: err.Error()})
		return
	}
	urls, err := uc.UrlService.FetchAllByUserID(c, userID, q.Page, q.Size)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Messgae: err.Error()})
		return
	}
	c.JSON(http.StatusOK, urls)
}

// 删除短链接
// DELETE /api/url/:code
func (uc *UrlController) Delete(c *gin.Context) {
	shortCode := c.Param("code")
	userID := c.GetHeader("x-user-id") // 获取请求头中的x-user-id字段

	if err := uc.UrlService.Delete(c, shortCode, userID); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "Url deleted successfully")
}

// 更新短链接有效期
// PATCH /api/url/:code
func (uc *UrlController) UpdateByExpiryTime(c *gin.Context) {
	shortCode := c.Param("code")
	userID := c.GetHeader("x-user-id") // 获取请求头中的x-user-id字段
	var json domain.UpdateByExpiryTimeRequest

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Messgae: err.Error()})
		return
	}

	if err := uc.UrlService.UpdateByExpiryTime(c, shortCode, json.ExpiryTime, userID); err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Messgae: err.Error()})
		return
	}

	c.JSON(http.StatusOK, "Url updated successfully")
}
