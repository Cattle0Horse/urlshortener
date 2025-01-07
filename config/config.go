package config

type Mode string

const (
	ModeDebug   Mode = "debug"
	ModeRelease Mode = "release"
)

type Config struct {
	Server *Server `yaml:"server" mapstructure:"server"`
	MySQL  *MySQL  `yaml:"mysql" mapstructure:"mysql"`
	JWT    *JWT    `yaml:"jwt" mapstructure:"jwt"`
	Url    *Url    `yaml:"url" mapstructure:"url"`
	TDDL   *TDDL   `yaml:"tddl" mapstructure:"tddl"`
	Cache  *Cache  `yaml:"cache" mapstructure:"cache"`
	Cors   *Cors   `yaml:"cors" mapstructure:"cors"`
}
