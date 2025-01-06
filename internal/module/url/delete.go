package url

import (
	"errors"

	"github.com/Cattle0Horse/url-shortener/internal/global/constant"
	"github.com/Cattle0Horse/url-shortener/internal/global/database"
	"github.com/Cattle0Horse/url-shortener/internal/global/errs"
	"github.com/Cattle0Horse/url-shortener/internal/global/jwt"
	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context) {
	shortCode := c.Param("code")
	if shortCode == "" {
		errs.Fail(c, errs.InvalidRequest.WithTips("short code is empty"))
		return
	}

	// 获取用户ID
	payload, ok := c.Get("payload")
	if !ok {
		errs.Fail(c, errs.InvaildToken.WithOrigin(errors.New("payload not found")))
		return
	}
	userID := payload.(jwt.Claims).UserId

	// 预检测，合法性与布隆过滤器
	exists, err := CheckCode(c.Request.Context(), []byte(shortCode))
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
		errs.Fail(c, errs.Forbidden.WithTips("you don't have permission to delete this url"))
		return
	}

	// 删除数据库记录(软删除)
	if _, err = u.WithContext(c.Request.Context()).Where(u.ID.Eq(url.ID)).Delete(); err != nil {
		errs.Fail(c, errs.ErrDatabase.WithOrigin(err))
		return
	}

	// 删除缓存
	if err := cache.Del(c, constant.ShortCodeCacheKey+shortCode); err != nil {
		log.Error("Failed to delete cache", "error", err)
	}

	// 不删除布隆过滤器中的记录，这需要重建，可以选择定时重建
	// if err := bloom.Delete(c, shortCode); err != nil {
	// 	log.Error("Failed to delete bloom filter", "error", err)
	// }

	errs.Success(c, nil)
}
