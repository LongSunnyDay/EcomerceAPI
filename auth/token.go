package auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"go-api-ws/helpers"
	"time"
)

const MySecret = "SenelisMegstaMociutesApvalumus"

func ParseToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There is an error")
		}
		return []byte(MySecret), nil
	})
	helpers.PanicErr(err)
	return token, nil
}

func GetNewAuthToken(sub int64, role string) (*jwt.Token) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  sub,
		"exp":  time.Now().Add(time.Hour * 1).Unix(),
		"role": role,
	})
	return token
}

func GetNewRefreshToken(sub int64) (*jwt.Token) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": sub,
		"exp": time.Now().Add(time.Hour * 4).Unix(),
	})
	return token
}