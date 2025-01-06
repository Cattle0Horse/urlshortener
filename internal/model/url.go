package model

import (
	"time"
)

type Url struct {
	Model
	ShortCode   string    `json:"short_code" gorm:"type:varchar(255);not null;uniqueIndex;comment:短链接"`
	OriginalUrl string    `json:"original_url" gorm:"type:varchar(255);not null;comment:原始链接"`
	ExpiryTime  time.Time `json:"expiry_time" gorm:"type:time;not null;comment:过期时间"`
	UserID      uint      `json:"user_id" gorm:"not null;comment:创建用户ID"`
}
