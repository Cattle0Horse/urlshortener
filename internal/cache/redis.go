package cache

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)


type RedisCache struct {
	client            *redis.Client
	urlDuration       time.Duration
	emailCodeDuration time.Duration
}

func NewRedisCache(address string, password string, db int, urlDuration time.Duration, emailCodeDuration time.Duration) (*RedisCache, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       db,
	})

	if err := client.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}

	return &RedisCache{
		client:            client,
		urlDuration:       urlDuration,
		emailCodeDuration: emailCodeDuration,
	}, nil
}


func (c *RedisCache) Close() error {
	return c.client.Close()
}
