// Provide base structure and common function for url module
package url

import (
	"context"
	"time"

	"github.com/Cattle0Horse/url-shortener/internal/model"
	"github.com/Cattle0Horse/url-shortener/pkg/base62"
)

type Url struct {
	ShortCode   string    `json:"short_code"`
	OriginalUrl string    `json:"original_url"`
	ExpiryTime  time.Time `json:"expiry_time"`
	CreatedAt   time.Time `json:"created_at"`
}

func (u *Url) ConvertFromModel(url *model.Url) {
	u.ShortCode = url.ShortCode
	u.OriginalUrl = url.OriginalUrl
	u.ExpiryTime = url.ExpiryTime
	u.CreatedAt = url.CreatedAt
}

func GenerateShortCode(c context.Context) ([]byte, error) {
	// 生成唯一ID
	id, err := tddlGen.Next(c)
	if err != nil {
		log.Error("Failed to generate ID", "error", err)
		return nil, err
	}

	// base62编码
	shortCode := base62.Encode(id)
	return shortCode, nil
}

// 检测短代码的合法性
// 1. 长度检测
// 2. 字符检测
func CheckCodeLegality(c context.Context, code []byte) bool {
	// TODO: implement me
	panic("implement me")
}

// 检查短代码是否存在
// 1. 布隆过滤器检测(结果是 "可能存在" 与 "一定不存在")
// 2. 数据库查询
func CheckCodeExists(c context.Context, code []byte) (bool, error) {
	// TODO: implement me
	panic("implement me")
}

// 合法性与存在性检测
func CheckCode(c context.Context, code []byte) (bool, error) {
	if !CheckCodeLegality(c, code) {
		return false, nil
	}
	exists, err := CheckCodeExists(c, code)
	return exists, err
}
