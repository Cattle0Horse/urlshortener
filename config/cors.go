package config

type Cors struct {
	Disabled     bool     `yaml:"disabled" mapstructure:"disabled"`
	AllowOrigins []string `yaml:"allow_origins" mapstructure:"allow_origins"`
}
