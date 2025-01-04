package config

import (
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	App      *AppConfig      `mapstructure:"app"`
	Database *DatabaseConfig `mapstructure:"database"`
	Redis    *RedisConfig    `mapstructure:"redis"`
	Logger   *LogConfig      `mapstructure:"logger"`
	JWT      *JWTConfig      `mapstructure:"jwt"`
	Email    *EmailConfig    `mapstructure:"email"`
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

type RedisConfig struct {
	Address           string        `mapstructure:"address"`
	Password          string        `mapstructure:"password"`
	DB                int           `mapstructure:"db"`
	UrlDuration       time.Duration `mapstructure:"url_duration"`
	EmailCodeDuration time.Duration `mapstructure:"email_code_duration"`
}

type LogConfig struct {
	Level string `mapstructure:"level"`
}

type JWTConfig struct {
	AccessTokenSecret   string        `mapstructure:"access_token_secret"`
	AccessTokenDuration time.Duration `mapstructure:"access_token_duration"`
	// RefreshTokenSecret   string        `mapstructure:"refresh_token_secret"`
	// RefreshTokenDuration time.Duration `mapstructure:"refresh_token_duration"`
}

type EmailConfig struct {
	Password    string `mapstructure:"password"`
	Username    string `mapstructure:"username"`
	HostAddress string `mapstructure:"host_address"`
	HostPort    string `mapstructure:"host_port"`
	Subject     string `mapstructure:"subject"`
	TestMail    string `mapstructure:"test_mail"`
}
