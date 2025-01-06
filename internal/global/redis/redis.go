// Package redis provides a redis client
package redis

import (
	"time"

	"github.com/Cattle0Horse/url-shortener/config"
	"github.com/redis/go-redis/v9"
)

var (
	// Nil is the redis.Nil, used to check if a key exists
	Nil = redis.Nil
	// 全局客户端
	Client redis.UniversalClient
)

func Init() {
	cfg := config.Get().Cache.Redis
	Client = NewClient(cfg.Addrs, cfg.DialTimeout, cfg.MaxConn)
}

// Client returns a redis client
func NewClient(addrs []string, dialTimeout time.Duration, maxConn int) redis.UniversalClient {
	return redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:          addrs,
		DialTimeout:    dialTimeout,
		MaxIdleConns:   maxConn,
		MaxActiveConns: maxConn,
	})
}
