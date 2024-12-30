package entity

import (
	"time"

	"gorm.io/gorm"
)

type Url struct {
	ID          uint           `json:"id" gorm:"primarykey"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	OriginalUrl string         `json:"original_url"`
	ShortCode   string         `json:"short_code" gorm:"index"`
	ExpiryTime  time.Time      `json:"expiry_time"`
	UserID      uint           `json:"user_id" gorm:"index"`
}
