package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"go-api-ws/auth"
	"go-api-ws/helpers"
	"net/http"
	"time"
)

const role = "user"

func protectedEndpoint(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		urlToken := req.URL.Query()["token"][0]
		token := auth.ParseToken(urlToken)
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if claims["role"] == role && claims.VerifyExpiresAt(time.Now().Unix(), true) {
				handlerFunc.ServeHTTP(w, req)
			} else {
				helpers.WriteResultWithStatusCode(w, "Token expired", http.StatusForbidden)
			}
		} else {
			helpers.WriteResultWithStatusCode(w, "Invalid token", http.StatusBadRequest)
		}
	}

}
