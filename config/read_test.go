package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInit(t *testing.T) {
	t.Run("should panic when config file not exist", func(t *testing.T) {
		assert.Panics(t, func() {
			Init("non_exist_config.yaml")
		})
	})

	t.Run("should load config with default path", func(t *testing.T) {
		// Test
		assert.NotPanics(t, func() {
			Init()
		})

		// Verify
		assert.NotNil(t, Get())
	})

	t.Run("should load config with custom path", func(t *testing.T) {
		// Setup
		originalConfig := c
		defer func() { c = originalConfig }()

		// Test
		assert.NotPanics(t, func() {
			Init("config.yaml")
		})

		// Verify
		assert.NotNil(t, c)
	})
}

func TestSetAndGet(t *testing.T) {
	t.Run("should set and get config correctly", func(t *testing.T) {
		// Setup
		originalConfig := c
		defer func() { c = originalConfig }()

		testConfig := &Config{
			Server: &Server{
				Mode: ModeDebug,
			},
		}

		// Test
		Set(testConfig)
		actual := Get()

		// Verify
		require.Equal(t, testConfig, actual)
	})
}

func TestIsReleaseAndIsDebug(t *testing.T) {
	t.Run("should return correct mode status", func(t *testing.T) {
		// Setup
		originalConfig := c
		defer func() { c = originalConfig }()

		testCases := []struct {
			name      string
			mode      Mode
			isDebug   bool
			isRelease bool
		}{
			{"debug mode", ModeDebug, true, false},
			{"release mode", ModeRelease, false, true},
			{"unknown mode", "unknown", false, false},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				Set(&Config{
					Server: &Server{
						Mode: tc.mode,
					},
				})

				assert.Equal(t, tc.isDebug, IsDebug())
				assert.Equal(t, tc.isRelease, IsRelease())
			})
		}
	})
}
