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
	// 检查长度
	if len(code) == 0 || len(code) > 8 {
		// 短链接长度应该在1-8之间
		// 因为我们使用uint64生成ID,base62编码后最长8位
		return false
	}

	// 检查是否包含非法字符
	// 尝试解码,如果失败说明包含非法字符
	_, err := base62.Decode(code)
	return err == nil
}

// 检查短代码是否可能存在
// 1. 布隆过滤器检测(结果是 "可能存在" 与 "一定不存在")
func CodeMayExists(c context.Context, code []byte) (bool, error) {
	// 布隆过滤器检测
	exists, err := bloom.MayExists(c, string(code))
	if err != nil {
		return false, err
	}
	return exists, nil
}

// 合法性与可能存在性检测（借助布隆过滤器）
func PreCheckCode(c context.Context, code []byte) (bool, error) {
	if !CheckCodeLegality(c, code) {
		return false, nil
	}
	return CodeMayExists(c, code)
}
