// Package redis provides a redis client
package redis

import (
	"github.com/Cattle0Horse/url-shortener/config"
	"github.com/redis/go-redis/v9"
)

// Nil is the redis.Nil, used to check if a key exists
var Nil = redis.Nil

// Client returns a redis client
func Client() redis.UniversalClient {
	c := config.Get().Cache.Redis
	return redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:          c.Addr,
		DialTimeout:    c.DialTimeout,
		MaxIdleConns:   c.MaxConn,
		MaxActiveConns: c.MaxConn,
	})
}
