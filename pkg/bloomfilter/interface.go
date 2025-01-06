package bloomfilter

import (
	"context"
	"errors"
)

var (
	ErrBloomFilterAlreadyExists = errors.New("bloom filter already exists")
)

// Interface 布隆过滤器接口
type Interface interface {
	SelfCheck(ctx context.Context) (bool, error)
	// Create 创建布隆过滤器
	Create(ctx context.Context) error
	// MayExists 检查元素是否存在
	MayExists(ctx context.Context, item string) (bool, error)
	// Add 添加元素到布隆过滤器
	Add(ctx context.Context, item string) error
	// Reset 重置布隆过滤器(调整容量和错误率)
	Reset(ctx context.Context, itemCount uint64, errorRate float64) error
	// Delete 删除布隆过滤器
	Delete(ctx context.Context) error
}
