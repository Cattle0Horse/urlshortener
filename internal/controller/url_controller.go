package controller

import (
	"net/http"

	"github.com/Cattle0Horse/url-shortener/internal/schema"
	"github.com/Cattle0Horse/url-shortener/internal/service"
	"github.com/gin-gonic/gin"
)

type UrlController struct {
	urlService *service.UrlService
}

func NewUrlController(urlService *service.UrlService) *UrlController {
	return &UrlController{
		urlService: urlService,
	}
}

// 创建一个短链接
// POST /api/url
func (uc *UrlController) Create(c *gin.Context) {
	var json schema.CreateUrlRequest
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, schema.ErrorResponse{Message: err.Error()})
		return
	}
	userID := c.GetInt("userID")

	shortUrl, err := uc.urlService.Create(c, json.OriginalUrl, json.Duration, uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, schema.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, schema.CreateUrlResponse{
		ShortUrl: shortUrl,
	})
}

// 重定向
// GET /:code
func (uc *UrlController) Redirect(c *gin.Context) {
	shortCode := c.Param("code")

	originalUrl, err := uc.urlService.FetchOriginalUrl(c, shortCode)
	if err != nil {
		// 404 错误
		c.JSON(http.StatusNotFound, schema.ErrorResponse{Message: err.Error()})
		return
	}
	// 永久重定向
	c.Redirect(http.StatusPermanentRedirect, originalUrl)
}

// 获取用户所有短链接
// GET /api/url?page=1&size=10
func (uc *UrlController) FetchAll(c *gin.Context) {
	userID := c.GetInt("userID")
	var q schema.FetchAllRequest
	if err := c.ShouldBindQuery(&q); err != nil {
		c.JSON(http.StatusBadRequest, schema.ErrorResponse{Message: err.Error()})
		return
	}
	if q.Page == 0 {
		q.Page = 1
	}

	if q.Size == 0 {
		q.Size = 10
	}

	urls, err := uc.urlService.FetchAllByUserID(c, uint(userID), q.Page, q.Size)
	if err != nil {
		c.JSON(http.StatusInternalServerError, schema.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, urls)
}

// 删除短链接
// DELETE /api/url/:code
func (uc *UrlController) Delete(c *gin.Context) {
	shortCode := c.Param("code")
	userID := c.GetInt("userID")

	if err := uc.urlService.Delete(c, shortCode, uint(userID)); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "Url deleted successfully")
}

// 更新短链接有效期
// PATCH /api/url/:code
func (uc *UrlController) UpdateByExpiryTime(c *gin.Context) {
	shortCode := c.Param("code")
	userID := c.GetInt("userID")
	var json schema.UpdateByExpiryTimeRequest

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, schema.ErrorResponse{Message: err.Error()})
		return
	}

	if err := uc.urlService.UpdateByExpiryTime(c, shortCode, json.ExpiryTime, uint(userID)); err != nil {
		c.JSON(http.StatusInternalServerError, schema.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, "Url updated successfully")
}
