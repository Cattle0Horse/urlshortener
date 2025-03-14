package cache

import (
	"context"
	"errors"
	"time"

	"github.com/Cattle0Horse/url-shortener/config"
	"github.com/allegro/bigcache/v3"
)

var (
	_ Interface = (*localCache)(nil)

	errInvalidCap = errors.New("cache: invalid capacity")
)

type localCache struct {
	cache *bigcache.BigCache
}

// NewLocalCache create a local cache
// capacity is the cache capacity
// ttl is the time to live
func NewLocalCache() (Interface, error) {
	return newLocalCache()
}

func newLocalCache() (*localCache, error) {
	c := config.Get().Cache.LocalCache
	if c.Capacity <= 0 {
		return nil, errInvalidCap
	}

	config := bigcache.DefaultConfig(c.TTL)
	config.MaxEntriesInWindow = c.Capacity
	config.HardMaxCacheSize = c.MaxMemory

	b, err := bigcache.New(context.Background(), config)
	if err != nil {
		return nil, err
	}

	return &localCache{cache: b}, err
}

// Set the k v pair to the cache
// Note that the duration is not used
func (l *localCache) Set(_ context.Context, k string, v []byte, _ time.Duration) error {
	return l.cache.Set(k, v)
}

// Get the value by key
func (l *localCache) Get(_ context.Context, k string) ([]byte, error) {
	v, err := l.cache.Get(k)
	if err != nil && errors.Is(err, bigcache.ErrEntryNotFound) {
		return nil, ErrCacheMiss
	}

	return v, err
}

func (l *localCache) Del(_ context.Context, k string) error {
	err := l.cache.Delete(k)
	if err != nil {
		if errors.Is(err, bigcache.ErrEntryNotFound) {
			return nil
		}

		return err
	}

	return nil
}

// Close the cache
func (l *localCache) Close() error {
	return l.cache.Close()
}
