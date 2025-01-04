package jwt

import (
	"time"

	"github.com/Cattle0Horse/url-shortener/config"
	"github.com/golang-jwt/jwt/v5"
)

type Payload struct {
	UserId uint `json:"user_id"`
}

type Claims struct {
	Payload
	jwt.RegisteredClaims
}

// CreateToken 签发用户Token
func CreateToken(payload Payload) (string, error) {
	claims := Claims{
		Payload: payload,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(config.Get().JWT.AccessExpire))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(config.Get().JWT.AccessSecret)
}

// ParseToken 解析用户Token
func ParseToken(token string) (claims *Claims, ok bool) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{},
		func(token *jwt.Token) (any, error) {
			return []byte(config.Get().JWT.AccessSecret), nil
		},
	)
	if err != nil || tokenClaims == nil {
		return nil, false
	}
	if claims, ok = tokenClaims.Claims.(*Claims); !ok || !tokenClaims.Valid {
		return nil, false
	}
	return claims, true
}
