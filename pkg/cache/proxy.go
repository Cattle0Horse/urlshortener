package cache

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/Cattle0Horse/url-shortener/config"
	"github.com/redis/go-redis/v9"
)

type proxy struct {
	distributedCache Interface
	localCache       Interface
	// localCacheTTL is the local cache ttl
	localCacheTTL time.Duration
	// remoteCacheTTL is the remote cache ttl
	remoteCacheTTL time.Duration
}

var _ Interface = (*proxy)(nil)

// NewProxy creates a new cache proxy, which contains a distributed cache and a local cache
func NewProxy(client redis.UniversalClient) (Interface, error) {
	return newProxy(client)
}

func newProxy(client redis.UniversalClient) (*proxy, error) {
	c := config.Get().Cache
	lc, err := NewLocalCache()
	if err != nil {
		return nil, err
	}

	return &proxy{
		distributedCache: NewRedisRemoteCache(client),
		localCache:       lc,
		remoteCacheTTL:   c.Redis.TTL,
		localCacheTTL:    c.LocalCache.TTL,
	}, nil
}

func (p *proxy) Set(ctx context.Context, k string, v []byte, ttl time.Duration) error {
	if err := p.distributedCache.Set(ctx, k, v, ttl); err != nil {
		return fmt.Errorf("failed to set distributed cache: %w", err)
	}

	if err := p.localCache.Set(ctx, k, v, ttl); err != nil {
		return fmt.Errorf("failed to set local cache: %w", err)
	}

	return nil
}

func (p *proxy) Get(ctx context.Context, k string) ([]byte, error) {
	// first, try to get from local cache
	long, err := p.localCache.Get(ctx, k)
	if err == nil {
		return long, nil
	}

	if !errors.Is(err, ErrCacheMiss) { // non cache miss error, return error
		return nil, err
	}

	defer func() { // try to set local cache if get from distributed cache
		if len(long) > 0 {
			if err = p.localCache.Set(ctx, k, long, p.localCacheTTL); err != nil {
				slog.ErrorContext(ctx, "failed to set local cache", slog.Any("error", err))
			}
		}
	}()

	// second, try to get from distributed cache
	long, err = p.distributedCache.Get(ctx, k) // need to fill the long variable

	return long, err
}

func (p *proxy) Del(ctx context.Context, k string) error {
	if err := p.distributedCache.Del(ctx, k); err != nil {
		return fmt.Errorf("failed to delete distributed cache: %w", err)
	}

	if err := p.localCache.Del(ctx, k); err != nil {
		return fmt.Errorf("failed to delete local cache: %w", err)
	}

	return nil
}

func (p *proxy) Close() error {
	if err := p.distributedCache.Close(); err != nil {
		return err
	}

	return p.localCache.Close()
}
