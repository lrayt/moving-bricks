package auth

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var jwtSecret = []byte("your-256-bit-secret")

type JWTUserClaims struct {
	UserInfo
	jwt.RegisteredClaims
}

func GenToken(user *UserInfo) (string, error) {
	claims := JWTUserClaims{
		UserInfo: *user,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "your-application-name",
			Subject:   "user-authentication",
		},
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(jwtSecret)
}

func ParseToken(token string) (*UserInfo, error) {
	tokenClaims, parseErr := jwt.ParseWithClaims(token, &JWTUserClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwtSecret, nil
	})

	if parseErr != nil || tokenClaims == nil {
		return nil, errors.New(fmt.Sprintf("parse token err:%v", parseErr))
	}

	if claims, ok := tokenClaims.Claims.(*JWTUserClaims); !ok || !tokenClaims.Valid || claims == nil {
		return nil, errors.New("token is invalid")
	} else {
		return &claims.UserInfo, nil
	}
}
