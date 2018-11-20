package cart

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"go-api-ws/attribute"
	"go-api-ws/auth"
	"go-api-ws/counter"
	"go-api-ws/helpers"
	"go-api-ws/product"
	"net/http"
	"time"
)

func createCart(w http.ResponseWriter, r *http.Request) {
	urlToken := r.URL.Query()["token"][0]
	if len(urlToken) > 0 {
		token, _ := auth.ParseToken(urlToken)

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if claims.VerifyExpiresAt(time.Now().Unix(), true) {
				cartID, err := getUserCartIDFromMongo(claims["sub"].(string))
				helpers.CheckErr(err)
				if err != nil {
					cartID := CreateCartInMongoDB(claims["sub"].(string))
					response := Response{
						Code:   http.StatusOK,
						Result: cartID}
					helpers.WriteResultWithStatusCode(w, response, response.Code)
				} else {
					response := Response{
						Code:   http.StatusOK,
						Result: cartID}
					helpers.WriteResultWithStatusCode(w, response, response.Code)
				}
			}
		}
	} else {
		cartID := CreateCartInMongoDB("")
		response := Response{
			Code:   http.StatusOK,
			Result: cartID}
		helpers.WriteResultWithStatusCode(w, response, response.Code)
	}
}

func pullCart(w http.ResponseWriter, r *http.Request) {
	urlUserToken := r.URL.Query()["token"][0]
	urlCartId := r.URL.Query()["cartId"][0]
	if len(urlUserToken) > 0 {
		token, _ := auth.ParseToken(urlUserToken)
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if claims.VerifyExpiresAt(time.Now().Unix(), true) {
				cart := getUserCartFromMongoByID(urlCartId)
				response := Response{
					Code:   http.StatusOK,
					Result: cart}
				helpers.WriteResultWithStatusCode(w, response, response.Code)
			}
		}
	} else {
		cart := getGuestCartFromMongoByID(urlCartId)
		response := Response{
			Code:   http.StatusOK,
			Result: cart}
		helpers.WriteResultWithStatusCode(w, response, response.Code)
	}
}

func addPaymentMethod(w http.ResponseWriter, r *http.Request) {
	var methods []interface{}
	_ = json.NewDecoder(r.Body).Decode(&methods)
	insertPaymentMethodsToMongo(methods)
	helpers.WriteResultWithStatusCode(w, "ok", 200)
}

func getPaymentMethods(w http.ResponseWriter, r *http.Request) {
	paymentMethods := getPaymentMethodsFromMongo()
	response := Response{
		Code:   http.StatusOK,
		Result: paymentMethods}
	helpers.WriteResultWithStatusCode(w, response, response.Code)
}

func updateCart(w http.ResponseWriter, r *http.Request) {
	urlUserToken := r.URL.Query()["token"][0]
	urlCartId := r.URL.Query()["cartId"][0]

	var item CartItem
	_ = json.NewDecoder(r.Body).Decode(&item)
	var attributes []attribute.ItemAttribute
	for _, itemOptions := range item.Item.ProductOption.ExtensionAttributes.ConfigurableItemOptions {
		attributes = append(attributes, attribute.GetAttributeNameFromSolr(itemOptions.OptionsID, itemOptions.OptionValue))
	}

	sku := product.BuildSKUFromItemAttributes(attributes, item.Item.SKU)
	productFromSolr := product.GetProductFromSolrBySKU(sku)

	item.Item.SKU = productFromSolr.Sku
	item.Item.Price = productFromSolr.Price
	item.Item.ProductType = productFromSolr.TypeID
	item.Item.Name = productFromSolr.Name
	item.Item.ItemID = counter.GetAndIncreaseItemIdCounterInMongo()

	if len(urlUserToken) > 0 {
		updateUserCartInMongo(urlCartId, item.Item)
		response := Response{
			Code:   http.StatusOK,
			Result: item.Item}
		helpers.WriteResultWithStatusCode(w, response, response.Code)
	} else {
		updateGuestCartInMongo(urlCartId, item.Item)
		response := Response{
			Code:   http.StatusOK,
			Result: item.Item}
		helpers.WriteResultWithStatusCode(w, response, response.Code)
	}
}

func deleteFromUserCart(w http.ResponseWriter, r *http.Request) {
	urlUserToken := r.URL.Query()["token"][0]
	urlCartId := r.URL.Query()["cartId"][0]
	var item CartItem
	_ = json.NewDecoder(r.Body).Decode(&item)

	if len(urlUserToken) > 0 {
		deleteItemFromCartInMongo(urlCartId, item)
		response := Response{
			Code:   http.StatusOK,
			Result: true}
		helpers.WriteResultWithStatusCode(w, response, response.Code)
	} else {
		deleteItemFromGuestCartInMongo(urlCartId, item)
		response := Response{
			Code:   http.StatusOK,
			Result: true}
		helpers.WriteResultWithStatusCode(w, response, response.Code)
	}
}
