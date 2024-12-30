package jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWT struct {
	accessTokenSecret   []byte
	accessTokenDuration time.Duration
}

func NewJWT(ats []byte, atd time.Duration) *JWT {
	return &JWT{
		accessTokenSecret:   ats,
		accessTokenDuration: atd,
	}
}

type UserCliams struct {
	Email  string `json:"email"`
	UserID uint   `json:"user_id"`
	jwt.RegisteredClaims
}

func (j *JWT) Generate(email string, userID uint) (string, error) {
	claims := UserCliams{
		Email:  email,
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.accessTokenDuration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.accessTokenSecret)
}

func (j *JWT) ParseToken(tokenString string) (*UserCliams, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserCliams{}, func(t *jwt.Token) (interface{}, error) {
		return j.accessTokenSecret, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*UserCliams); ok {
		return claims, nil
	}

	return nil, fmt.Errorf("failed to parseToken")
}
