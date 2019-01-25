package auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"go-api-ws/config"
	"go-api-ws/helpers"
	"time"
)

func ParseToken(tokenString string) *jwt.Token {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there is an error when parsing token")
		}
		return []byte(config.MySecret), nil
	})
	helpers.PanicErr(err)
	return token
}

func GetTokenClaims(token *jwt.Token) (jwt.MapClaims, error) {
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, fmt.Errorf("something whent wrong then mapping claims")
	}
}

func CheckIfTokenIsNotExpired(claims jwt.MapClaims) bool {
	if claims.VerifyExpiresAt(time.Now().Unix(), true) {
		return true
	}
	return false
}

func GetNewAuthToken(sub string, groupId int) *jwt.Token {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":     sub,
		"exp":     time.Now().Add(time.Hour * 4).Unix(),
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
