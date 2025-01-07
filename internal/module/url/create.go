package url

import (
	"errors"
	"time"

	"github.com/Cattle0Horse/url-shortener/config"
	"github.com/Cattle0Horse/url-shortener/internal/global/constant"
	"github.com/Cattle0Horse/url-shortener/internal/global/database"
	"github.com/Cattle0Horse/url-shortener/internal/global/errs"
	"github.com/Cattle0Horse/url-shortener/internal/global/jwt"
	"github.com/Cattle0Horse/url-shortener/internal/model"
	"github.com/gin-gonic/gin"
)

type CreateRequest struct {
	OriginalUrl string `json:"original_url" binding:"required,url,min=10,max=255"`
	Duration    *int   `json:"duration" binding:"omitempty,min=1,max=168"`
	UserID      uint   `json:"user_id" binding:"-"`
}

func (r *CreateRequest) ConvertToUrl(url *model.Url) {
	url.OriginalUrl = r.OriginalUrl
	if r.Duration == nil {
		url.ExpiryTime = time.Now().Add(config.Get().Url.DefaultDuration)
	} else {
		url.ExpiryTime = time.Now().Add(time.Hour * time.Duration(*r.Duration))
	}
	url.UserID = r.UserID
}

type CreateResponse struct {
	ShortCode string `json:"short_code"`
}

func Create(c *gin.Context) {
	var req CreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		errs.Fail(c, errs.InvalidRequest.WithTips(err.Error()))
		return
	}

	log.Info("Create Url", "request", req)

	payload, ok := c.Get("payload")
	if !ok {
		errs.Fail(c, errs.InvaildToken.WithOrigin(errors.New("payload not found")))
		return
	}
	req.UserID = payload.(*jwt.Claims).UserId

	var url model.Url
	req.ConvertToUrl(&url)

	// 生成短链接
	shortCode, err := GenerateShortCode(c)
	if err != nil {
		errs.Fail(c, errs.FailedGenShortCode.WithOrigin(err))
		return
	}
	url.ShortCode = string(shortCode)

	// 加入短代码到布隆过滤器
	if err := bloom.Add(c, url.ShortCode); err != nil {
		errs.Fail(c, errs.ErrBloomFilter.WithOrigin(err))
		return
	}

	// 保存到数据库
	if err := database.Query.Url.WithContext(c.Request.Context()).Create(&url); err != nil {
		errs.Fail(c, errs.ErrDatabase.WithOrigin(err))
		return
	}

	// 计算剩余过期时间
	ttl := time.Until(url.ExpiryTime)
	if ttl > cacheTTL {
		// 如果剩余时间大于默认缓存时间,使用默认缓存时间
		ttl = cacheTTL
	}

	// 更新缓存,使用较小的过期时间
	// 保存到缓存
	if err := cache.Set(c, constant.ShortCodeCacheKey+url.ShortCode, []byte(req.OriginalUrl), ttl); err != nil {
		log.Error("Failed to set cache", "error", err)
	}

	errs.Success(c, CreateResponse{
		ShortCode: url.ShortCode,
	})
}
