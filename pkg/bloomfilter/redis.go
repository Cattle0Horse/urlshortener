package bloomfilter

import (
	"context"
	"fmt"
	"log/slog"

	goredis "github.com/redis/go-redis/v9"
)

var (
	_ Interface = (*redisBloomFilter)(nil)
)

// redisBloomFilter Redis实现的布隆过滤器
type redisBloomFilter struct {
	// Redis客户端
	client goredis.UniversalClient
	// 布隆过滤器的key
	key string
	// 预计元素数量
	itemCount uint64
	// 期望错误率
	errorRate float64
}

// New 创建一个新的布隆过滤器
func NewRedisBloomFilter(client goredis.UniversalClient, key string, itemCount uint64, errorRate float64) Interface {
	return newRedisBloomFilter(client, key, itemCount, errorRate)
}

func newRedisBloomFilter(client goredis.UniversalClient, key string, itemCount uint64, errorRate float64) *redisBloomFilter {
	return &redisBloomFilter{
		client:    client,
		key:       key,
		itemCount: itemCount,
		errorRate: errorRate,
	}
}

// 检查布隆过滤器是否已存在
func (bf *redisBloomFilter) SelfCheck(ctx context.Context) (bool, error) {
	exists, err := bf.client.Exists(ctx, bf.key).Result()
	if err != nil {
		return false, fmt.Errorf("failed to check bloom filter existence: %w", err)
	}
	if exists == 1 {
		return true, nil
	}
	return false, nil
}

// Create 创建布隆过滤器
func (bf *redisBloomFilter) Create(ctx context.Context) error {
	// 检查布隆过滤器是否已存在
	exists, err := bf.SelfCheck(ctx)
	if err != nil {
		return err
	}
	if exists {
		return ErrBloomFilterAlreadyExists
	}

	// 创建布隆过滤器
	err = bf.client.Do(ctx, "BF.RESERVE", bf.key, bf.errorRate, bf.itemCount).Err()
	if err != nil {
		return fmt.Errorf("failed to create bloom filter: %w", err)
	}

	slog.Info("Created bloom filter",
		"key", bf.key,
		"itemCount", bf.itemCount,
		"errorRate", bf.errorRate,
	)
	return nil
}

// Add 添加元素到布隆过滤器
func (bf *redisBloomFilter) Add(ctx context.Context, item string) error {
	err := bf.client.Do(ctx, "BF.ADD", bf.key, item).Err()
	if err != nil {
		return fmt.Errorf("failed to add item to bloom filter: %w", err)
	}
	return nil
}

// MayExists 检查元素是否存在
func (bf *redisBloomFilter) MayExists(ctx context.Context, item string) (bool, error) {
	exists, err := bf.client.Do(ctx, "BF.EXISTS", bf.key, item).Bool()
	if err != nil {
		return false, fmt.Errorf("failed to check item existence in bloom filter: %w", err)
	}
	return exists, nil
}

// Reset 重置布隆过滤器
func (bf *redisBloomFilter) Reset(ctx context.Context, itemCount uint64, errorRate float64) error {
	// 删除旧的布隆过滤器
	err := bf.Delete(ctx)
	if err != nil {
		return fmt.Errorf("failed to delete old bloom filter: %w", err)
	}

	// 更新参数
	bf.itemCount = itemCount
	bf.errorRate = errorRate

	// 创建新的布隆过滤器
	return bf.Create(ctx)
}

// Delete 删除布隆过滤器
func (bf *redisBloomFilter) Delete(ctx context.Context) error {
	err := bf.client.Del(ctx, bf.key).Err()
	if err != nil {
		return fmt.Errorf("failed to delete bloom filter: %w", err)
	}
	return nil
}
