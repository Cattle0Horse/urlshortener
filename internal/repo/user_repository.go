package repo

import (
	"context"

	"github.com/Cattle0Horse/url-shortener/internal/entity"
	"gorm.io/gorm"
)

type userRepository struct {
	database *gorm.DB
}

func NewUserRepository(database *gorm.DB) *userRepository {
	return &userRepository{
		database: database,
	}
}
func (r *userRepository) Create(c context.Context, user *entity.User) error {
	result := r.database.Create(user)
	return result.Error
}
func (r *userRepository) GetByEmail(c context.Context, email string) (*entity.User, error) {
	var user entity.User
	result := r.database.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
func (r *userRepository) IsEmailAvailable(ctx context.Context, email string) (bool, error) {
	var user entity.User
	result := r.database.Where("email = ?", email).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return true, nil
		}
		return false, result.Error
	}
	return false, nil
}
func (r *userRepository) UpdatePasswordByEmail(ctx context.Context, email string, password string) (*entity.User, error) {
	var user entity.User
	result := r.database.Model(&user).Where("email = ?", email).Update("password", password)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
