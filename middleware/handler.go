package middleware

import (
	"go-api-ws/auth"
	"go-api-ws/helpers"
	"net/http"
)

const (
	role = "user"
)

func protectedEndpoint(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		urlToken, err := helpers.GetTokenFromUrl(req)
		helpers.PanicErr(err)
		token := auth.ParseToken(urlToken)
		claims, err := auth.GetTokenClaims(token)
		helpers.CheckErr(err)
		if err != nil {
			helpers.WriteResultWithStatusCode(w, "Invalid token", http.StatusBadRequest)
		} else {
			if claims["role"] == role && auth.CheckIfTokenIsNotExpired(claims) {
				handlerFunc.ServeHTTP(w, req)
			} else {
				helpers.WriteResultWithStatusCode(w, "Token expired", http.StatusForbidden)
			}
		}
	}

}
