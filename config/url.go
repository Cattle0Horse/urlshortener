package config

import "time"

type Url struct {
	DefaultDuration time.Duration `yaml:"default_duration" mapstructure:"default_duration"`
}
