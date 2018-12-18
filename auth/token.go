package auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"go-api-ws/config"
	"go-api-ws/helpers"
	"time"
)

func ParseToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there is an error")
		}
		return []byte(config.MySecret), nil
	})
	helpers.PanicErr(err)
	return token, nil
}

func GetNewAuthToken(sub string, groupId int) *jwt.Token {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":     sub,
		"exp":     time.Now().Add(time.Hour * 1).Unix(),
		"groupId": groupId,
	})
	return token
}

func GetNewRefreshToken(sub string) *jwt.Token {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": sub,
		"exp": time.Now().Add(time.Hour * 4).Unix(),
	})
	return token
}
