package cart

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/kjk/betterguid"
	"go-api-ws/auth"
	"go-api-ws/helpers"
	"net/http"
	"time"
)

func createCart(w http.ResponseWriter, req *http.Request) {
	urlToken := req.URL.Query()["token"][0]
	if len(urlToken) > 0 {
		token, _ := auth.ParseToken(urlToken)

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if claims.VerifyExpiresAt(time.Now().Unix(), true) {
				cartID := getCartIDFromMongo(claims["sub"].(string), "Registered user")
				response := Response{
					Code:   http.StatusOK,
					Result: cartID}
				helpers.WriteResultWithStatusCode(w, response, response.Code)
			}
		}
	} else {
		id := betterguid.New()
		guestCartId := createGuestCartInMongo(id)
		response := Response{
			Code:   http.StatusOK,
			Result: guestCartId}
		helpers.WriteResultWithStatusCode(w, response, response.Code)
	}
}

func pullCart(w http.ResponseWriter, req *http.Request)  {
	urlCartId := req.URL.Query()["cartId"][0]
	getCartFromMongoByID(urlCartId)

}