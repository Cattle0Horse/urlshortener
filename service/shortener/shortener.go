package shortener

import (
	"sync"

	"github.com/Cattle0Horse/url-shortener/pkg/base62"
	"github.com/Cattle0Horse/url-shortener/pkg/bloom"
)

const (
	itemCount         = uint64(1000000) // 预计存储的URL数量
	falsePositiveRate = 0.01            // 期望的误判率
)

var (
	bloomFilter *bloom.Filter
	bloomOnce   sync.Once
)

func initBloomFilter() {
	bloomOnce.Do(func() {
		bloomFilter = bloom.New(itemCount, falsePositiveRate)
	})
}

func ShortenURL(url string) string {
	initBloomFilter()

	id := snowflake.NextID()
	shortURL := base62.Encode(id)

	// Add to bloom filter
	bloomFilter.Add(shortURL)

	return shortURL
}

func CheckCodeExists(shortCode string) bool {
	initBloomFilter()
	return bloomFilter.MayExist([]byte(shortCode))
}
