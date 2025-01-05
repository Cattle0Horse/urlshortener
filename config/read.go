package config

import (
	"github.com/Cattle0Horse/url-shortener/tools"
	"github.com/spf13/viper"
)

const defaultFilePath = "config.yaml"

var c *Config

func Init(path ...string) {
	filePath := defaultFilePath
	if len(path) == 1 {
		filePath = path[0]
	}
	viper.SetConfigFile(filePath)
	if tools.FileExist(filePath) {
		tools.PanicOnErr(viper.ReadInConfig())
		tools.PanicOnErr(viper.Unmarshal(&c))
	} else {
		panic("Config file not exist in " + filePath)
		// fmt.Println("Config file not exist in ", filePath, ". Using environment variables.")
		// tools.PanicOnErr(envconfig.Process("", c))
	}
}

func Set(config *Config) {
	c = config
}

func Get() *Config {
	return c
}

func IsRelease() bool {
	return c.Server.Mode == ModeRelease
}

func IsDebug() bool {
	return c.Server.Mode == ModeDebug
}
