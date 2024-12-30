package shortcode

import (
	"math/rand"
)

type ShortCode struct {
	lenght int
}

func NewShortCode(length int) *ShortCode {
	return &ShortCode{
		lenght: length,
	}
}

const chars = "abcdefjhijklmnopqrstuvwsyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func (s *ShortCode) GenerateShortCode() string {
	length := len(chars)
	result := make([]byte, s.lenght)

	for i := 0; i < s.lenght; i++ {
		result[i] = chars[rand.Intn(length)]
	}
	return string(result)
}
