package utils

import (
	"OrderManagement/OrderManagement/internal/config"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type JwtClaims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func GenerateToken(c config.Config, username, password string) (string, error) {
	var jwtSecret = []byte(c.Auth.AccessSecret)

	nowTime := time.Now()
	expireTime := nowTime.Add(time.Duration(c.Auth.AccessExpire))

	claims := JwtClaims{
		username,
		password,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "orderManager",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	SignedString, err := token.SignedString(jwtSecret)

	return SignedString, err
}

func ParseToken(c config.Config, token string) (*JwtClaims, error) {
	var jwtSecret = []byte(c.Auth.AccessSecret)
	tokenClaims, err := jwt.ParseWithClaims(token, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*JwtClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
