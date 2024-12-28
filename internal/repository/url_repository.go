package repository

import (
	"context"
	"time"

	"github.com/Cattle0Horse/url-shortener/internal/domain"
	"gorm.io/gorm"
)

type urlRepository struct {
	database *gorm.DB
}

func NewUrlRepository(database *gorm.DB) domain.UrlRepository {
	return &urlRepository{
		database: database,
	}
}

func (ur *urlRepository) Create(c context.Context, url *domain.Url) error {
	result := ur.database.Create(url)
	return result.Error
}
func (ur *urlRepository) FetchOriginalUrlByShortCode(c context.Context, shortCode string) (string, error) {
	var url domain.Url
	result := ur.database.Where("short_code = ?", shortCode).First(&url)
	if result.Error != nil {
		return "", result.Error
	}
	return url.OriginalUrl, nil
}
func (ur *urlRepository) UpdateByExpiryTime(c context.Context, shortCode string, expiryTime time.Time, userID string) error {
	var url domain.Url
	result := ur.database.Model(&url).Where("short_code = ? AND user_id = ? AND deleted_at IS NULL", shortCode, userID).Update("expiry_time", expiryTime)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (ur *urlRepository) DeleteByShortCode(c context.Context, shortCode string, userID string) error {
	var url domain.Url
	result := ur.database.Model(&url).Where("short_code = ? AND user_id = ? AND deleted_at IS NULL", shortCode, userID).Update("deleted_at", time.Now())
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (ur *urlRepository) FetchAllByUserID(c context.Context, userID string, page int, size int) ([]domain.Url, error) {
	var urls []domain.Url
	result := ur.database.Where("user_id = ? AND deleted_at IS NULL", userID).Offset((page - 1) * size).Limit(size).Find(&urls)
	if result.Error != nil {
		return nil, result.Error
	}
	return urls, nil
}
func (ur *urlRepository) IsShortCodeAvailable(c context.Context, shortCode string) (bool, error) {
	// 查询短代码是否已经存在
	var url domain.Url
	result := ur.database.Where("short_code = ? AND deleted_at IS NULL", shortCode).First(&url)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return true, nil
		}
		return false, result.Error
	}
	return false, nil
}
