package cart

import (
	"encoding/json"
	"fmt"
	"go-api-ws/attribute"
	"go-api-ws/auth"
	"go-api-ws/counter"
	"go-api-ws/helpers"
	"go-api-ws/product"
	"net/http"
)

func createCart(w http.ResponseWriter, r *http.Request) {
	urlToken, err := helpers.GetTokenFromUrl(r)
	helpers.CheckErr(err)

	if len(urlToken) > 0 {
		token := auth.ParseToken(urlToken)
		claims, err := auth.GetTokenClaims(token)
		helpers.CheckErr(err)

		if err != nil {
			helpers.WriteResultWithStatusCode(w, err, http.StatusForbidden)
		} else {
			if auth.CheckIfTokenIsNotExpired(claims) {
				cartID, err := CheckDoesUserHasACart(claims["sub"].(string))
				if err != nil {
					fmt.Println(err)
					cartID = CreateCartInMongoDB(claims["sub"].(string))
				}
				response := helpers.Response{
					Code:   http.StatusOK,
					Result: cartID}
				response.SendResponse(w)
			} else {
				helpers.WriteResultWithStatusCode(w, "Token is expired", http.StatusForbidden)
			}
		}
	} else {
		// ToDo After placing order request to createCart comes in without token in url
		cartID := CreateCartInMongoDB("")
		response := helpers.Response{
			Code:   http.StatusOK,
			Result: cartID}
		response.SendResponse(w)
	}
}

func pullCart(w http.ResponseWriter, r *http.Request) {
	urlUserToken, err := helpers.GetTokenFromUrl(r)
	helpers.CheckErr(err)

	urlCartId, err := helpers.GetCartIdFromUrl(r)
	helpers.CheckErr(err)

	if len(urlUserToken) > 0 {
		token := auth.ParseToken(urlUserToken)
		claims, err := auth.GetTokenClaims(token)
		helpers.CheckErr(err)

		if err != nil {
			helpers.WriteResultWithStatusCode(w, err, http.StatusForbidden)
		} else {
			if auth.CheckIfTokenIsNotExpired(claims) {
				cart := GetUserCartFromMongoByID(urlCartId)
				response := helpers.Response{
					Code:   http.StatusOK,
					Result: cart.Items}
				response.SendResponse(w)
			} else {
				helpers.WriteResultWithStatusCode(w, "Token is expired", http.StatusForbidden)
			}
		}
	} else {
		cart := GetUserCartFromMongoByID(urlCartId)
		response := helpers.Response{
			Code:   http.StatusOK,
			Result: cart.Items}
		response.SendResponse(w)
	}
}

func updateCart(w http.ResponseWriter, r *http.Request) {
	urlCartId, err := helpers.GetCartIdFromUrl(r)
	helpers.CheckErr(err)

	var item CustomerCart
	err = json.NewDecoder(r.Body).Decode(&item)
	helpers.PanicErr(err)
	var attributes []attribute.ItemAttribute
	for _, itemOptions := range item.Item.ProductOption.ExtensionAttributes.ConfigurableItemOptions {
		attributes = append(attributes, attribute.GetAttributeNameFromSolr(itemOptions.OptionsID, itemOptions.OptionValue))
	}

	productFromSolr := product.GetProductFromSolrBySKU(item.Item.SKU)
	if len(item.Item.SKU) == 4 {
		item.Item.SKU = product.BuildSKUFromItemAttributes(attributes, item.Item.SKU)
	}
	item.Item.Price = productFromSolr.Price
	item.Item.ProductType = productFromSolr.TypeID
	item.Item.Name = productFromSolr.Name
	item.Item.ItemID = counter.GetAndIncreaseItemIdCounterInMongo()

	updateUserCartInMongo(urlCartId, item.Item)
	response := helpers.Response{
		Code:   http.StatusOK,
		Result: item.Item}
	response.SendResponse(w)
}

func deleteFromUserCart(w http.ResponseWriter, r *http.Request) {
	urlCartId, err := helpers.GetCartIdFromUrl(r)
	helpers.CheckErr(err)
	var item CustomerCart
	err = json.NewDecoder(r.Body).Decode(&item)
	helpers.PanicErr(err)
	deleteItemFromCartInMongo(urlCartId, item)
	response := helpers.Response{
		Code:   http.StatusOK,
		Result: true}
	response.SendResponse(w)
}
