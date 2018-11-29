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
					response := helpers.Response{
						Code:http.StatusOK,
						Result:cartID}
					response.SendResponse(w)
				} else {
					response := helpers.Response{
						Code:http.StatusOK,
						Result:cartID}
					response.SendResponse(w)
				}
			}
		}
	} else {
		cartID := CreateCartInMongoDB("")
		response := helpers.Response{
			Code:http.StatusOK,
			Result:cartID}
		response.SendResponse(w)
	}
}

func pullCart(w http.ResponseWriter, r *http.Request) {
	urlUserToken := r.URL.Query()["token"][0]
	urlCartId := r.URL.Query()["cartId"][0]
	if len(urlUserToken) > 0 {
		token, _ := auth.ParseToken(urlUserToken)
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if claims.VerifyExpiresAt(time.Now().Unix(), true) {
				cart := GetUserCartFromMongoByID(urlCartId)
				response := helpers.Response{
					Code:   http.StatusOK,
					Result: cart}
				response.SendResponse(w)
			}
		}
	} else {
		cart := getGuestCartFromMongoByID(urlCartId)
		response := helpers.Response{
			Code:   http.StatusOK,
			Result: cart}
		response.SendResponse(w)
	}
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

	productFromSolr := product.GetProductFromSolrBySKU(item.Item.SKU)

	item.Item.SKU = productFromSolr.Sku
	item.Item.Price = productFromSolr.Price
	item.Item.ProductType = productFromSolr.TypeID
	item.Item.Name = productFromSolr.Name
	item.Item.ItemID = counter.GetAndIncreaseItemIdCounterInMongo()

	if len(urlUserToken) > 0 {
		updateUserCartInMongo(urlCartId, item.Item)
		response := helpers.Response{
			Code:   http.StatusOK,
			Result: item.Item}
		response.SendResponse(w)
	} else {
		updateGuestCartInMongo(urlCartId, item.Item)
		response := helpers.Response{
			Code:   http.StatusOK,
			Result: item.Item}
		response.SendResponse(w)
	}
}

func deleteFromUserCart(w http.ResponseWriter, r *http.Request) {
	urlUserToken := r.URL.Query()["token"][0]
	urlCartId := r.URL.Query()["cartId"][0]
	var item CartItem
	_ = json.NewDecoder(r.Body).Decode(&item)

	if len(urlUserToken) > 0 {
		deleteItemFromCartInMongo(urlCartId, item)
		response := helpers.Response{
			Code:   http.StatusOK,
			Result: true}
		response.SendResponse(w)
	} else {
		deleteItemFromGuestCartInMongo(urlCartId, item)
		response := helpers.Response{
			Code:   http.StatusOK,
			Result: true}
		response.SendResponse(w)
	}
}
