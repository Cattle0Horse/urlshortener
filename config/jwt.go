package config

import "time"

type JWT struct {
	AccessSecret string        `yaml:"access_secret" mapstructure:"access_secret"`
	AccessExpire time.Duration `yaml:"access_expire" mapstructure:"access_expire"`
}
