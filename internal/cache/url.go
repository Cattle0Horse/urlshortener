package cache

import (
	"context"

	"github.com/Cattle0Horse/url-shortener/internal/entity"
)

const urlPrifix = "url:"

func (c *RedisCache) SetURL(ctx context.Context, url *entity.Url) error {
	if err := c.client.Set(ctx, urlPrifix+url.ShortCode, url.OriginalUrl, c.urlDuration).Err(); err != nil {
		return err
	}

	return nil
}

func (c *RedisCache) GetURL(ctx context.Context, shortCode string) (originalURL string, err error) {
	originalURL = c.client.Get(ctx, urlPrifix+shortCode).Val()

	return originalURL, nil
}

func (c *RedisCache) DelURL(ctx context.Context, shortCode string) error {
	return c.client.Del(ctx, urlPrifix+shortCode).Err()
}
