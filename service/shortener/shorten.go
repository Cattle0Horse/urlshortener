package shortener

import (
	"math/big"
	"sync"

	"github.com/Cattle0Horse/url-shortener/internal/global/cache"
	"github.com/Cattle0Horse/url-shortener/tools"
)

const (
	base62Charset      = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	bloomFilterSize    = 1000000
	bloomFilterHashNum = 7
)

var (
	snowflake   *tools.Snowflake
	bloomFilter *cache.BloomFilter
	once        sync.Once
	bloomOnce   sync.Once
)

func initSnowflake() {
	once.Do(func() {
		var err error
		snowflake, err = tools.NewSnowflake(1) // workerID=1
		if err != nil {
			panic(err)
		}
	})
}

func initBloomFilter() {
	bloomOnce.Do(func() {
		bloomFilter = cache.NewBloomFilter(bloomFilterSize, bloomFilterHashNum)
	})
}

func numberToBase62(num int64) string {
	base := big.NewInt(62)
	result := make([]byte, 0)

	// Convert number to base62
	for num > 0 {
		quotient := new(big.Int)
		remainder := new(big.Int)
		quotient, remainder = quotient.DivMod(big.NewInt(num), base, remainder)
		num = quotient.Int64()
		result = append(result, base62Charset[remainder.Int64()])
	}

	// Reverse the result
	reversed := make([]byte, len(result))
	for i, j := 0, len(result)-1; j >= 0; i, j = i+1, j-1 {
		reversed[i] = result[j]
	}

	return string(reversed)
}

func ShortenURL(url string) string {
	initSnowflake()
	initBloomFilter()

	id := snowflake.NextID()
	shortURL := numberToBase62(id)

	// Add to bloom filter
	bloomFilter.Add(shortURL)

	return shortURL
}

func CheckURLExists(shortURL string) bool {
	initBloomFilter()
	return bloomFilter.Contains(shortURL)
}
