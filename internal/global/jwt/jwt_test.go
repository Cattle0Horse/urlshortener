package jwt

import (
	"testing"
	"time"

	"github.com/Cattle0Horse/url-shortener/config"
	"github.com/stretchr/testify/assert"
)

func TestJWT(t *testing.T) {
	// Setup test config
	testConfig := &config.JWT{
		AccessSecret: "test-secret",
		AccessExpire: time.Hour,
	}
	Init(testConfig)

	t.Run("should initialize jwt config correctly", func(t *testing.T) {
		assert.Equal(t, []byte("test-secret"), accessSecret)
		assert.Equal(t, time.Hour, accessExpire)
	})

	t.Run("should create and parse valid token", func(t *testing.T) {
		payload := Payload{UserId: 123}

		// Create token
		token, err := CreateToken(payload)
		assert.NoError(t, err)
		assert.NotEmpty(t, token)

		// Parse token
		claims, ok := ParseToken(token)
		assert.True(t, ok)
		assert.Equal(t, uint(123), claims.UserId)
		assert.WithinDuration(t, time.Now().Add(time.Hour), claims.ExpiresAt.Time, time.Second)
	})

	t.Run("should fail to parse invalid token", func(t *testing.T) {
		testCases := []struct {
			name  string
			token string
		}{
			{"empty token", ""},
			{"malformed token", "invalid.token.string"},
			{"expired token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxMjMsImV4cCI6MTY4MDAwMDAwMH0.invalid-signature"},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				claims, ok := ParseToken(tc.token)
				assert.False(t, ok)
				assert.Nil(t, claims)
			})
		}
	})

	t.Run("should fail to parse token with wrong secret", func(t *testing.T) {
		// Create token with different secret
		Init(&config.JWT{
			AccessSecret: "different-secret",
			AccessExpire: time.Hour,
		})
		token, err := CreateToken(Payload{UserId: 123})
		assert.NoError(t, err)

		// Revert to original secret
		Init(testConfig)

		// Try to parse
		claims, ok := ParseToken(token)
		assert.False(t, ok)
		assert.Nil(t, claims)
	})
}
