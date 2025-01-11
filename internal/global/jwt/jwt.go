package jwt

import (
	"time"

	"github.com/Cattle0Horse/url-shortener/config"
	"github.com/golang-jwt/jwt/v5"
)

// TODO: 添加签名算法验证
// TODO: 添加Issuer和Audience验证
// TODO: 实现Token刷新机制
// TODO: 添加Token黑名单功能
var (
	accessSecret []byte
	accessExpire time.Duration
)

func Init(cfg *config.JWT) {
	accessSecret = []byte(cfg.AccessSecret)
	accessExpire = cfg.AccessExpire
}

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
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(accessExpire)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(accessSecret)
}

// ParseToken 解析用户Token
func ParseToken(token string) (claims *Claims, ok bool) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{},
		func(token *jwt.Token) (any, error) {
			return accessSecret, nil
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
