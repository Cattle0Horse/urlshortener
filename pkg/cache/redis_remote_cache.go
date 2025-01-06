// Package cache implements a distributed cache Interface
// redis_remote_cache.go implements a distributed cache with redis
package cache

import (
	"context"
	"errors"
	"math/rand/v2"
	"time"

	"github.com/redis/go-redis/v9"
)

var _ Interface = (*redisCache)(nil)

type redisCache struct {
	rdb redis.UniversalClient
	// bucketSize int
}

// NewRedisRemoteCache returns a new redis cache
func NewRedisRemoteCache(client redis.UniversalClient) Interface {
	return newRedisCache(client)
}

func newRedisCache(client redis.UniversalClient) *redisCache {
	return &redisCache{
		rdb: client,
	}
}

// Set the k v pair to the cache
func (c *redisCache) Set(ctx context.Context, k string, v []byte, ttl time.Duration) error {
	//nolint:gosec,mnd
	ttl += time.Duration(rand.IntN(int(ttl / 10))) // add some jitter
	return c.rdb.SetEx(ctx, k, v, ttl).Err()
}

// Get the value by key
func (c *redisCache) Get(ctx context.Context, k string) ([]byte, error) {
	value, err := c.rdb.Get(ctx, k).Bytes()
	if err != nil && errors.Is(err, redis.Nil) {
		return nil, ErrCacheMiss
	}

	return value, err
}

func (c *redisCache) Del(ctx context.Context, k string) error {
	return c.rdb.Del(ctx, k).Err()
}

// Close the cache
func (c *redisCache) Close() error {
	return c.rdb.Close()
}
