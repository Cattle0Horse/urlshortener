package url

import (
	"errors"
	"time"

	"github.com/Cattle0Horse/url-shortener/internal/global/constant"
	"github.com/Cattle0Horse/url-shortener/internal/global/database"
	"github.com/Cattle0Horse/url-shortener/internal/global/errs"
	"github.com/Cattle0Horse/url-shortener/internal/global/jwt"
	"github.com/gin-gonic/gin"
)

type UpdateRequest struct {
	Duration int `json:"duration" binding:"required,min=1,max=168"` // 168小时 = 7天
}

func Update(c *gin.Context) {
	shortCode := c.Param("code")
	if shortCode == "" {
		errs.Fail(c, errs.InvalidRequest.WithTips("short code is empty"))
		return
	}

	var req UpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		errs.Fail(c, errs.InvalidRequest.WithTips(err.Error()))
		return
	}

	// 获取用户ID
	payload, ok := c.Get("payload")
	if !ok {
		errs.Fail(c, errs.InvaildToken.WithOrigin(errors.New("payload not found")))
		return
	}
	userID := payload.(*jwt.Claims).UserId

	// 预检测，合法性与布隆过滤器
	exists, err := PreCheckCode(c.Request.Context(), []byte(shortCode))
	if err != nil {
		log.Error("Failed to check bloom filter", "error", err)
		errs.Fail(c, errs.ErrBloomFilter.WithOrigin(err))
		return
	}
	if !exists {
		// 布隆过滤器显示一定不存在
		errs.Fail(c, errs.NotFound.WithTips("url not found"))
		return
	}

	// 查询短链接
	u := database.Query.Url
	url, err := u.WithContext(c.Request.Context()).Where(u.DeletedAt.IsNull(), u.ShortCode.Eq(shortCode)).First()
	if err != nil {
		errs.Fail(c, errs.NotFound.WithOrigin(err))
		return
	}

	// 验证用户权限
	if url.UserID != userID {
		errs.Fail(c, errs.Forbidden.WithTips("you don't have permission to update this url"))
		return
	}

	// 更新过期时间
	newExpiryTime := time.Now().Add(time.Hour * time.Duration(req.Duration))
	_, err = u.WithContext(c.Request.Context()).Where(u.ID.Eq(url.ID)).
		Updates(map[string]interface{}{
			"expiry_time": newExpiryTime,
		})
	if err != nil {
		errs.Fail(c, errs.ErrDatabase.WithOrigin(err))
		return
	}

	// 计算剩余过期时间
	ttl := time.Until(newExpiryTime)
	if ttl > cacheTTL {
		// 如果剩余时间大于默认缓存时间,使用默认缓存时间
		ttl = cacheTTL
	}

	// 更新缓存
	if err := cache.Set(c, constant.ShortCodeCacheKey+shortCode, []byte(url.OriginalUrl), ttl); err != nil {
		log.Error("Failed to set cache", "error", err)
	}

	errs.Success(c, nil)
}
