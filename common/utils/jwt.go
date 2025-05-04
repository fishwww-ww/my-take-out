package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type CustomPayload struct {
	UserId     uint64
	GrantScope string
	jwt.RegisteredClaims
}

func GenerateToken(uid uint64, subject string, secret string) (string, error) {
	claims := CustomPayload{
		UserId:     uid,
		GrantScope: subject,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "Auth_Server",                                 //签发者
			Subject:   subject,                                       //签发对象
			Audience:  jwt.ClaimStrings{"PC", "Wechat_Program"},      //签发受众
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)), //过期时间
			NotBefore: jwt.NewNumericDate(time.Now()),                //生效时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),                //签发时间
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secret))
	return token, err
}
