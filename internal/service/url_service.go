package service

import (
	"context"
	"errors"
	"math/rand"
	"time"

	"github.com/Cattle0Horse/url-shortener/internal/entity"
)

type UrlRepository interface {
	Create(c context.Context, url *entity.Url) error
	FetchOriginalUrlByShortCode(c context.Context, shortCode string) (string, error)
	UpdateByExpiryTime(c context.Context, shortCode string, expiryTime time.Time, userID uint) error
	DeleteByShortCode(c context.Context, shortCode string, userID uint) error
	FetchAllByUserID(c context.Context, userID uint, page int, size int) ([]entity.Url, error)
	IsShortCodeAvailable(c context.Context, shortCode string) (bool, error)
}

type UrlService struct {
	urlRepository   UrlRepository
	defaultDuration time.Duration
	serverAddress   string
}

func NewUrlService(urlRepository UrlRepository, defaultDuration time.Duration, serverAddress string) *UrlService {
	return &UrlService{
		urlRepository:   urlRepository,
		defaultDuration: defaultDuration,
		serverAddress:   serverAddress,
	}
}

func (us *UrlService) Create(c context.Context, originalUrl string, duration *int, userID uint) (string, error) {
	shortCode, err := us.generatorShortCode(c)
	if err != nil {
		return "", err
	}

	var expiryTime time.Time
	if duration == nil {
		expiryTime = time.Now().Add(us.defaultDuration)
	} else {
		expiryTime = time.Now().Add(time.Hour * time.Duration(*duration))
	}
	url := &entity.Url{
		OriginalUrl: originalUrl,
		ShortCode:   shortCode,
		ExpiryTime:  expiryTime,
		UserID:      userID,
	}

	// 插入数据库
	err = us.urlRepository.Create(c, url)
	if err != nil {
		return "", err
	}

	shortUrl := us.serverAddress + "/" + url.ShortCode
	return shortUrl, nil
}

func (us *UrlService) FetchAllByUserID(c context.Context, userID uint, page int, size int) ([]entity.Url, error) {
	urls, err := us.urlRepository.FetchAllByUserID(c, userID, page, size)
	if err != nil {
		return nil, err
	}
	return urls, nil
}

func (us *UrlService) FetchOriginalUrl(c context.Context, shortCode string) (string, error) {
	// 访问数据库
	originalUrl, err := us.urlRepository.FetchOriginalUrlByShortCode(c, shortCode)
	if err != nil {
		return "", err
	}
	return originalUrl, nil
}

func (us *UrlService) Delete(c context.Context, shortCode string, userID uint) error {
	err := us.urlRepository.DeleteByShortCode(c, shortCode, userID)
	if err != nil {
		return err
	}
	return nil
}

func (us *UrlService) UpdateByExpiryTime(c context.Context, shortCode string, expiryTime time.Time, userID uint) error {

	err := us.urlRepository.UpdateByExpiryTime(c, shortCode, expiryTime, userID)
	if err != nil {
		return err
	}
	return nil
}

func (us *UrlService) generatorShortCode(c context.Context) (string, error) {
	// todo: change the function
	var recursive func(int) (string, error)
	recursive = func(n int) (string, error) {
		if n > 5 {
			return "", errors.New("重试过多")
		}
		shortCode := generatorShortCodeDetail()

		isAvailable, err := us.urlRepository.IsShortCodeAvailable(c, shortCode)
		if err != nil {
			return "", err
		}

		if isAvailable {
			return shortCode, nil
		}

		return recursive(n + 1)
	}
	return recursive(0)
}

const chars = "abcdefjhijklmnopqrstuvwsyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func generatorShortCodeDetail() string {
	length := len(chars)
	result := make([]byte, 6)

	for i := 0; i < 6; i++ {
		result[i] = chars[rand.Intn(length)]
	}
	return string(result)
}
