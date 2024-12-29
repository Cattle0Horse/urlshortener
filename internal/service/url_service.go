package service

import (
	"context"
	"errors"
	"math/rand"
	"time"

	"github.com/Cattle0Horse/url-shortener/internal/domain"
)

type UrlRepository interface {
	Create(c context.Context, url *domain.Url) error
	FetchOriginalUrlByShortCode(c context.Context, shortCode string) (string, error)
	UpdateByExpiryTime(c context.Context, shortCode string, expiryTime time.Time, userID string) error
	DeleteByShortCode(c context.Context, shortCode string, userID string) error
	FetchAllByUserID(c context.Context, userID string, page int, size int) ([]domain.Url, error)
	IsShortCodeAvailable(c context.Context, shortCode string) (bool, error)
}
type urlService struct {
	urlRepository   UrlRepository
	contextTimeout  time.Duration
	defaultDuration time.Duration
	serverAddress   string
}

func NewUrlService(urlRepository UrlRepository, timeout time.Duration, defaultDuration time.Duration, serverAddress string) *urlService {
	return &urlService{
		urlRepository:   urlRepository,
		contextTimeout:  timeout,
		defaultDuration: defaultDuration,
		serverAddress:   serverAddress,
	}
}

func (us *urlService) Create(c context.Context, originalUrl string, duration *int, userID string) (string, error) {
	ctx, cancel := context.WithTimeout(c, us.contextTimeout)
	defer cancel()

	shortCode, err := us.generatorShortCode(ctx)
	if err != nil {
		return "", err
	}

	var expiryTime time.Time
	if duration == nil {
		expiryTime = time.Now().Add(us.defaultDuration)
	} else {
		expiryTime = time.Now().Add(time.Hour * time.Duration(*duration))
	}
	url := &domain.Url{
		OriginalUrl: originalUrl,
		ShortCode:   shortCode,
		ExpiryTime:  expiryTime,
		UserID:      userID,
	}

	// 插入数据库
	err = us.urlRepository.Create(ctx, url)
	if err != nil {
		return "", err
	}

	shortUrl := us.serverAddress + "/" + url.ShortCode
	return shortUrl, nil
}

func (us *urlService) FetchAllByUserID(c context.Context, userID string, page int, size int) ([]domain.Url, error) {
	ctx, cancel := context.WithTimeout(c, us.contextTimeout)
	defer cancel()

	urls, err := us.urlRepository.FetchAllByUserID(ctx, userID, page, size)
	if err != nil {
		return nil, err
	}
	return urls, nil
}

func (us *urlService) FetchOriginalUrl(c context.Context, shortCode string) (string, error) {
	ctx, cancel := context.WithTimeout(c, us.contextTimeout)
	defer cancel()

	// 访问数据库
	originalUrl, err := us.urlRepository.FetchOriginalUrlByShortCode(ctx, shortCode)
	if err != nil {
		return "", err
	}
	return originalUrl, nil
}

func (us *urlService) Delete(c context.Context, shortCode string, userID string) error {
	ctx, cancel := context.WithTimeout(c, us.contextTimeout)
	defer cancel()

	err := us.urlRepository.DeleteByShortCode(ctx, shortCode, userID)
	if err != nil {
		return err
	}
	return nil
}

func (us *urlService) UpdateByExpiryTime(c context.Context, shortCode string, expiryTime time.Time, userID string) error {
	ctx, cancel := context.WithTimeout(c, us.contextTimeout)
	defer cancel()

	err := us.urlRepository.UpdateByExpiryTime(ctx, shortCode, expiryTime, userID)
	if err != nil {
		return err
	}
	return nil
}

func (us *urlService) generatorShortCode(c context.Context) (string, error) {
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
