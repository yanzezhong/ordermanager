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

func GenerateToken(c config.Config, username, password string) (string, string, error) {
	var jwtSecret = []byte(c.Auth.AccessSecret)

	nowTime := time.Now()
	acccessExpireTime := nowTime.Add(time.Duration(c.Auth.AccessExpire))
	refreshExpireTime := nowTime.Add(time.Duration(c.Auth.RefreshExpire))

	accessClaims := JwtClaims{
		username,
		password,
		jwt.StandardClaims{
			ExpiresAt: acccessExpireTime.Unix(),
			Issuer:    "orderManager",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)

	accessTokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", "", err
	}

	refreshClaims := JwtClaims{
		username,
		password,
		jwt.StandardClaims{
			ExpiresAt: refreshExpireTime.Unix(),
			Issuer:    "orderManager",
		},
	}

	token = jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshTokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", "", err
	}
	return accessTokenString, refreshTokenString, err
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
