package domain

import (
	"time"

	"gorm.io/gorm"
)

// --database entity
type Url struct {
	ID          uint           `json:"id" gorm:"primarykey"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	OriginalUrl string         `json:"original_url"`
	ShortCode   string         `json:"short_code" gorm:"index"`
	ExpiryTime  time.Time      `json:"expiry_time"`
	UserID      string         `json:"user_id" gorm:"index"`
}

// --repository

// --service

// --controller

// Create
type CreateUrlRequest struct {
	OriginalUrl string `json:"original_url" binding:"required,url"`
	Duration    *int   `json:"duraion,omitempty" binding:"omitempty,min=1,max=100"`
}
type CreateUrlResponse struct {
	ShortUrl string `json:"short_url"`
}

// Redirect

// FetchAll
type FetchAllRequest struct {
	Page int `form:"page" binding:"required,min=1"`
	Size int `form:"size" binding:"required,min=1"`
}

// Delete

// UpdateByExpiryTime
type UpdateByExpiryTimeRequest struct {
	ExpiryTime time.Time `json:"expiry_time"`
}
