package cart

import (
	"fmt"
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
				cart := getUserCartFromMongo(claims["sub"].(float64))
				response := Response{
					Code:   http.StatusOK,
					Result: cart}
				helpers.WriteResultWithStatusCode(w, response, response.Code)
			}
		}
	} else {
		id := betterguid.New()
		CreateUserCartInMongo(id)
		response := Response{
			Code:   http.StatusOK,
			Result: id}
		helpers.WriteResultWithStatusCode(w, response, response.Code)
	}
}

func pullCart(w http.ResponseWriter, req *http.Request)  {
	urlToken := req.URL.Query()["token"][0]
	urlCartId := req.URL.Query()["cartId"][0]
	fmt.Println("pullCart", urlToken, urlCartId)


}