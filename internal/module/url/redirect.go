package url

import (
	"errors"
	"net/http"
	"time"

	"github.com/Cattle0Horse/url-shortener/internal/global/constant"
	"github.com/Cattle0Horse/url-shortener/internal/global/database"
	"github.com/Cattle0Horse/url-shortener/internal/global/errs"
	pkgcache "github.com/Cattle0Horse/url-shortener/pkg/cache"
	"github.com/gin-gonic/gin"
)

// TODO: 注意缓存数据库一致性问题
func Redirect(c *gin.Context) {
	shortCode := c.Param("code")
	if shortCode == "" {
		errs.Fail(c, errs.InvalidRequest.WithTips("short code is empty"))
		return
	}

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

	// 先从缓存获取
	originalUrl, err := cache.Get(c, constant.ShortCodeCacheKey+shortCode)
	if err == nil {
		// 缓存命中,直接重定向，307临时重定向
		c.Redirect(http.StatusTemporaryRedirect, string(originalUrl))
		return
	}
	if !errors.Is(err, pkgcache.ErrCacheMiss) {
		log.Error("Failed to get cache", "error", err)
		errs.Fail(c, errs.ErrCache.WithOrigin(err))
		return
	}

	// 缓存未命中,从数据库获取
	u := database.Query.Url
	url, err := u.WithContext(c.Request.Context()).Where(u.DeletedAt.IsNull(), u.ShortCode.Eq(shortCode)).First()
	if err != nil {
		errs.Fail(c, errs.NotFound.WithOrigin(err))
		return
	}

	// 检查是否过期
	if time.Now().After(url.ExpiryTime) {
		errs.Fail(c, errs.NotFound.WithTips("url expired"))
		return
	}

	// 计算剩余过期时间
	ttl := time.Until(url.ExpiryTime)
	if ttl > cacheTTL {
		// 如果剩余时间大于默认缓存时间,使用默认缓存时间
		ttl = cacheTTL
	}

	// 更新缓存,使用较小的过期时间
	if err := cache.Set(c, constant.ShortCodeCacheKey+shortCode, []byte(url.OriginalUrl), ttl); err != nil {
		log.Error("Failed to set cache", "error", err)
	}

	// 307临时重定向
	c.Redirect(http.StatusTemporaryRedirect, url.OriginalUrl)
}
