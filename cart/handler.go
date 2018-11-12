package cart

import (
	"encoding/json"
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

func pullCart(w http.ResponseWriter, req *http.Request) {
	//urlUserToken := req.URL.Query()["token"][0]
	urlCartId := req.URL.Query()["cartId"][0]
	cart := getCartFromMongoByID(urlCartId)
	response := Response{
		Code:   http.StatusOK,
		Result: cart.Items}
	helpers.WriteResultWithStatusCode(w, response, response.Code)
}

func addPaymentMethod(w http.ResponseWriter, req *http.Request) {
	var methods []interface{}
	_ = json.NewDecoder(req.Body).Decode(&methods)
	insertPaymentMethodsToMongo(methods)
	helpers.WriteResultWithStatusCode(w, "GG", 200)
}

func getPaymentMethods(w http.ResponseWriter, r *http.Request) {
	paymentMethods := getPaymentMethodsFromMongo()
	response := Response{
		Code:   http.StatusOK,
		Result: paymentMethods}
	helpers.WriteResultWithStatusCode(w, response, response.Code)
}

func updateCart(w http.ResponseWriter, r *http.Request)  {
	urlCartId := r.URL.Query()["cartId"][0]
	fmt.Println("updateCart - ", urlCartId)
	var item CartItem
	_ = json.NewDecoder(r.Body).Decode(&item)
	fmt.Println(item)
	updateUserCartInMongo(urlCartId, item)
}