package config

import (
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	App      AppConfig      `mapstructure:"app"`
	Database DatabaseConfig `mapstructure:"database"`
	Logger   LogConfig      `mapstructure:"logger"`
}

func NewConfig(configFilePath string) (*Config, error) {
	config := Config{}
	viper.SetConfigFile(configFilePath)

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	if config.App.Env == "development" {
		log.Println("The App is running in development env")
	}

	return &config, nil
}

type AppConfig struct {
	Env             string        `mapstructure:"env"`
	BaseUrl         string        `mapstructure:"base_url"`
	Port            string        `mapstructure:"port"`
	ContextTimeout  time.Duration `mapstructure:"context_timeout"`
	DefaultDuration time.Duration `mapstructure:"default_duration"`
}

type DatabaseConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DbName   string `mapstructure:"dbname"`
}

func (dc *DatabaseConfig) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dc.User, dc.Password, dc.Host, dc.Port, dc.DbName)
}

type LogConfig struct {
	Level string `mapstructure:"level"`
}
