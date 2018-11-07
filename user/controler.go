package user

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

func GetNewAuthToken(sub string, role string) (*jwt.Token) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  sub,
		"exp":  time.Now().Add(time.Hour * 1).Unix(),
		"role": role,
	})
	return token
}

func GetNewRefreshToken(sub string) (*jwt.Token) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  sub,
		"exp":  time.Now().Add(time.Hour * 4).Unix(),
	})
	return token
}