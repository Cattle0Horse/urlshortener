package database

import (
	"time"

	"github.com/Cattle0Horse/url-shortener/config"
	"github.com/Cattle0Horse/url-shortener/internal/global/query"
	"github.com/Cattle0Horse/url-shortener/internal/model"
	"github.com/Cattle0Horse/url-shortener/pkg/tools"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var Query *query.Query
var DB *gorm.DB

func Init() {
	cfg := config.Get().MySQL

	gormConfig := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名（表名不加s）
		},
		DisableForeignKeyConstraintWhenMigrating: true,
		SkipDefaultTransaction:                   true,
		TranslateError:                           true,
	}

	switch config.Get().Server.Mode {
	case config.ModeDebug:
		gormConfig.Logger = logger.Default.LogMode(logger.Info)
	case config.ModeRelease:
		gormConfig.Logger = logger.Discard
	}
	var err error
	DB, err = gorm.Open(mysql.Open(cfg.DSN()), gormConfig)

	tools.PanicOnErr(err)

	sqlDB, err := DB.DB()
	tools.PanicOnErr(err)
	sqlDB.SetMaxIdleConns(cfg.MaxConn)
	sqlDB.SetMaxOpenConns(cfg.MaxConn)
	sqlDB.SetConnMaxLifetime(time.Hour)
	sqlDB.SetConnMaxIdleTime(time.Minute)

	// tools.PanicOnErr(db.Use(otel.GetGormPlugin()))
	tools.PanicOnErr(DB.AutoMigrate(model.User{}, model.Url{}, model.Sequence{}))
	// 修改字段排序规则为大小写敏感
	// TODO: 优雅的处理
	DB.Exec("ALTER TABLE url MODIFY COLUMN short_code VARCHAR(255) COLLATE utf8_bin;")
	Query = query.Use(DB)
}
