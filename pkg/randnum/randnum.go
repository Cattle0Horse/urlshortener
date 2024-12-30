package randnum

import (
	"math/rand"
)

type RandNum struct {
	length int
}

func NewRandNum(length int) *RandNum {
	return &RandNum{
		length: length,
	}
}

const nums = "0123456789"

func (r *RandNum) Generate() string {
	result := make([]byte, r.length)

	length := len(nums)

	for i := range result {
		result[i] = nums[rand.Intn(length)]
	}

	return string(result)
}
