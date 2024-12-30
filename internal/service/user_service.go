package service

import (
	"context"
	"fmt"

	"github.com/Cattle0Horse/url-shortener/internal/entity"
	"github.com/Cattle0Horse/url-shortener/internal/schema"
	"github.com/gin-gonic/gin"
)

type UserRepository interface {
	Create(c context.Context, user *entity.User) error
	GetByEmail(c context.Context, email string) (*entity.User, error)
	IsEmailAvailable(ctx context.Context, email string) (bool, error)
	UpdatePasswordByEmail(ctx context.Context, email string, password string) (*entity.User, error)
}

type PasswordHasher interface {
	HashPassword(password string) (string, error)
	ComparePassword(hashedPassword string, password string) bool
}

type UserCacher interface {
	GetEmailCode(ctx context.Context, email string) (string, error)
	SetEmailCode(ctx context.Context, email, emailCode string) error
}

type EmailSender interface {
	Send(email, emailCode string) error
}

type NumberRandomer interface {
	Generate() string
}

type JWTer interface {
	Generate(email string, userID uint) (string, error)
}

type UserService struct {
	userRepository UserRepository
	passwordHasher PasswordHasher
	jwter          JWTer
	userCacher     UserCacher
	emailSender    EmailSender
	numberRandomer NumberRandomer
}

func NewUserService(ur UserRepository, ph PasswordHasher, jwter JWTer, uc UserCacher, es EmailSender, nr NumberRandomer) *UserService {
	return &UserService{
		userRepository: ur,
		passwordHasher: ph,
		jwter:          jwter,
		userCacher:     uc,
		emailSender:    es,
		numberRandomer: nr,
	}
}
func (us *UserService) Login(c context.Context, req schema.LoginRequest) (*schema.LoginResponse, error) {
	user, err := us.userRepository.GetByEmail(c, req.Email)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by email: %v", err)
	}
	if !us.passwordHasher.ComparePassword(user.Password, req.Password) {
		return nil, schema.ErrEmailOrPasswordFailed
	}
	accessToken, err := us.jwter.Generate(user.Email, user.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %v", err)
	}
	return &schema.LoginResponse{
		AccessToken: accessToken,
		Email:       user.Email,
		UserID:      user.ID,
	}, nil
}
func (us *UserService) IsEmailAvailable(c context.Context, email string) (bool, error) {
	isAvaliable, err := us.userRepository.IsEmailAvailable(c, email)
	if err != nil {
		return false, err
	}

	return isAvaliable, nil
}

func (us *UserService) Register(c *gin.Context, req schema.RegisterReqeust) (*schema.LoginResponse, error) {
	// 判断emailCode是否正确
	emailCode, err := us.userCacher.GetEmailCode(c, req.Email)
	if err != nil {
		return nil, fmt.Errorf("failed to get emailCode from cache: %v", err)
	}
	if emailCode != req.EmailCode {
		return nil, schema.ErrEmailCodeNotEqual
	}

	// hash密码
	hash, err := us.passwordHasher.HashPassword(req.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %v", err)
	}

	// 插入数据库
	user := &entity.User{
		Password: hash,
		Email:    req.Email,
	}
	if err := us.userRepository.Create(c, user); err != nil {
		return nil, fmt.Errorf("failed to create user: %v", err)
	}

	// 生成access token
	accessToken, err := us.jwter.Generate(user.Email, user.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to generate access token: %v", err)
	}

	return &schema.LoginResponse{
		AccessToken: accessToken,
		Email:       req.Email,
		UserID:      user.ID,
	}, nil
}

func (us *UserService) SendEmailCode(ctx context.Context, email string) error {
	emailCode := us.numberRandomer.Generate()

	if err := us.emailSender.Send(email, emailCode); err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}

	if err := us.userCacher.SetEmailCode(ctx, email, emailCode); err != nil {
		return fmt.Errorf("failed to set emailcode in cache: %v", err)
	}

	return nil
}

func (us *UserService) ResetPassword(ctx context.Context, req *schema.ResetPasswordReqeust) (*schema.LoginResponse, error) {
	emailCode, err := us.userCacher.GetEmailCode(ctx, req.Email)
	if err != nil {
		return nil, err
	}
	if emailCode != req.EmailCode {
		return nil, schema.ErrEmailCodeNotEqual
	}

	hash, err := us.passwordHasher.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}
	// 更新数据库
	user, err := us.userRepository.UpdatePasswordByEmail(ctx, req.Email, hash)
	if err != nil {
		return nil, err
	}

	accessToken, err := us.jwter.Generate(user.Email, user.ID)
	if err != nil {
		return nil, err
	}

	return &schema.LoginResponse{
		AccessToken: accessToken,
		Email:       user.Email,
		UserID:      user.ID,
	}, nil
}
