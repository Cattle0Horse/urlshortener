package config

import "time"

// Server is the config of turl server
type Server struct {
	Host   string `yaml:"host" mapstructure:"host"`
	Port   string `yaml:"port" mapstructure:"port"`
	Prefix string `yaml:"prefix" mapstructure:"prefix"`
	Mode   Mode   `yaml:"mode" mapstructure:"mode"`
	// Readonly is the read-only mode of turl server
	Readonly bool `yaml:"readonly" mapstructure:"readonly"`
	// RequestTimeout is the http server request timeout of turl server
	RequestTimeout time.Duration `validate:"required" yaml:"request_timeout" mapstructure:"request_timeout"`
	// GlobalRateLimitKey is the key of global rate limiter
	GlobalRateLimitKey string `validate:"required" yaml:"global_rate_limit_key" mapstructure:"global_rate_limit_key"`
	// GlobalWriteRate is the token bucket rate of write api rate limiter
	GlobalWriteRate int `validate:"required,gt=0" yaml:"global_write_rate" mapstructure:"global_write_rate"`
	// GlobalWriteBurst is the token bucket burst of write api rate limiter
	GlobalWriteBurst int `validate:"required,min=1" yaml:"global_write_burst" mapstructure:"global_write_burst"`
	// StandAloneReadRate is the token bucket rate of read api rate limiter
	StandAloneReadRate int `validate:"required,gt=0" yaml:"stand_alone_read_rate" mapstructure:"stand_alone_read_rate"`
	// StandAloneReadBurst is the token bucket burst of read api rate limiter
	StandAloneReadBurst int `validate:"required,min=1" yaml:"stand_alone_read_burst" mapstructure:"stand_alone_read_burst"`
}
