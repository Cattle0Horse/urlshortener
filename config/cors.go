package config

type Cors struct {
	AllowOrigins []string `yaml:"allow_origins" mapstructure:"allow_origins"`
}
