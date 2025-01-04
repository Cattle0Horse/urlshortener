package shortener

import (
	"github.com/spaolacci/murmur3"
	"math/big"
)

const base62Charset = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func hashToBase62(hashValue uint32) string {
	base := big.NewInt(62)
	result := make([]byte, 0)

	// Convert hashValue to base62
	for hashValue > 0 {
		quotient := new(big.Int)
		remainder := new(big.Int)
		quotient, remainder = quotient.DivMod(big.NewInt(int64(hashValue)), base, remainder)
		hashValue = uint32(quotient.Int64())
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
	// Step 1: Create a MurmurHash3 hasher
	hasher := murmur3.New32()

	// Step 2: Hash the URL
	hasher.Write([]byte(url))
	hashValue := hasher.Sum32()

	// Step 3: Convert the hash value to base62
	shortenedURL := hashToBase62(hashValue)

	return shortenedURL
}
