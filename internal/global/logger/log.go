package logger

import (
	"github.com/Cattle0Horse/url-shortener/config"
	"log/slog"
	"os"
	"sync"
)

var (
	instance *slog.Logger
	once     sync.Once
)

// Get 获取全局 Logger 实例
func Get() *slog.Logger {
	once.Do(func() {
		var handler slog.Handler
		switch config.Get().Server.Mode {
		case config.ModeDebug:
			handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{AddSource: false})
		case config.ModeRelease:
			handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{AddSource: true})
		}
		instance = slog.New(handler)
	})
	return instance
}

func New(name, value string) *slog.Logger {
	return Get().With(name, value)
}

// NewModule 创建一个新的 Logger 实例
func NewModule(module string) *slog.Logger {
	return New("module", module)
}

// NewService 创建一个新的 Logger 实例
func NewService(service string) *slog.Logger {
	return New("service", service)
}
