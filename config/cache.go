package config

import "time"

// Redis is the redis config
type Redis struct {
	// Addr is the redis address
	Addrs []string `validate:"required" yaml:"addrs" mapstructure:"addrs"`
	// DialTimeout Dial timeout for establishing new connections.
	// Default is 5 seconds.
	DialTimeout time.Duration `validate:"required" yaml:"dial_timeout" mapstructure:"dial_timeout"`
	// MaxIdleConn is the max open connections
	MaxConn int `validate:"required,min=1" yaml:"max_conn" mapstructure:"max_conn"`
	// TTL is the redis cache ttl
	TTL time.Duration `validate:"required" yaml:"ttl" mapstructure:"ttl"`
}

// LocalCache is the local cache config of turl server
type LocalCache struct {
	// TTL is the local cache ttl
	TTL time.Duration `validate:"required" yaml:"ttl" mapstructure:"ttl"`
	// Capacity is the local cache capacity
	Capacity int `validate:"required" yaml:"capacity" mapstructure:"capacity"`
	// MaxMemory max memory for value size in MB
	MaxMemory int `validate:"required" yaml:"max_memory" mapstructure:"max_memory"`
}

// Cache is the cache config of turl server
type Cache struct {
	// Redis is the redis config of turl server
	Redis *Redis `yaml:"redis" mapstructure:"redis"`
	// LocalCache is the local cache config
	LocalCache *LocalCache `validate:"required" yaml:"local_cache" mapstructure:"local_cache"`
}
