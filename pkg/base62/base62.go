package base62

import (
	"errors"
	"fmt"
)

const (
	// 62^8 < math.MaxUint64 < 62^9
	// so we can safely encode any uint64 to base62
	// if the input is larger than 8 bytes, we return an error
	carry    = 62
	maxBytes = 8
)

var (
	// ErrInvalidInput base error for invalid input
	ErrInvalidInput = errors.New("base62 invalid input")
	// ErrBase62Overflow is returned when the number to decode is too large, greater than eight bytes
	ErrBase62Overflow = fmt.Errorf("%w: number is too large", ErrInvalidInput)
	// ErrorInvalidCharacter is returned when an invalid character is found in the input
	ErrorInvalidCharacter = fmt.Errorf("%w: invalid character in input", ErrInvalidInput)
	// chars is the base62 alphabet
	chars = []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	// indices is an array of indices of characters in the base62 alphabet
	indices = [256]int{}
	// pow62 returns 62^n
	pow62 = [9]int{}
)

func init() {
	for i := range indices {
		indices[i] = -1
	}

	for i, char := range chars {
		indices[char] = i
	}

	pow62[0] = 1
	for i, n := 1, len(pow62); i < n; i++ {
		pow62[i] = pow62[i-1] * carry
	}
}

// Base62Encode encodes a number to base62
func Base62Encode(num uint64) []byte {
	b := make([]byte, 0, maxBytes)

	for ; num > 0; num /= carry {
		b = append(b, chars[num%carry])
	}
	reverse(b)

	return b
}

// Base62Decode decodes a base62 encoded string to a number
func Base62Decode(b []byte) (uint64, error) {
	n := len(b)
	if n > maxBytes {
		return 0, ErrBase62Overflow
	}

	var num uint64

	for i := range n {
		pos := indices[b[i]]
		if pos == -1 {
			return 0, ErrorInvalidCharacter
		}

		num += uint64(pow62[n-i-1] * pos)
	}

	return num, nil
}

func reverse(a []byte) {
	for left, right := 0, len(a)-1; left < right; left, right = left+1, right-1 {
		a[left], a[right] = a[right], a[left]
	}
}
