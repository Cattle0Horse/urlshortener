package schema

import "time"

type CreateUrlRequest struct {
	OriginalUrl string `json:"original_url" binding:"required,url"`
	Duration    *int   `json:"duraion,omitempty" binding:"omitempty,min=1,max=100"`
}
type CreateUrlResponse struct {
	ShortUrl string `json:"short_url"`
}

type FetchAllRequest struct {
	Page int `form:"page" binding:"required"`
	Size int `form:"size" binding:"required"`
}

type UpdateByExpiryTimeRequest struct {
	ExpiryTime time.Time `json:"expiry_time"`
}
