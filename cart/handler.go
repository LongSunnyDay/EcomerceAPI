package cart

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"go-api-ws/attribute"
	"go-api-ws/auth"
	"go-api-ws/counter"
	"go-api-ws/helpers"
	"go-api-ws/product"
	"net/http"
	"time"
)

func createCart(w http.ResponseWriter, req *http.Request) {
	urlToken := req.URL.Query()["token"][0]
	if len(urlToken) > 0 {
		token, _ := auth.ParseToken(urlToken)

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if claims.VerifyExpiresAt(time.Now().Unix(), true) {
				userId := getCartIDFromMongo(claims["sub"].(string), "Registered user")
				response := Response{
					Code:   http.StatusOK,
					Result: userId}
				helpers.WriteResultWithStatusCode(w, response, response.Code)
			}
		}
	} else {
		CreateCartInMongoDB("")
		response := Response{
			Code:   http.StatusOK,
			Result: ""}
		helpers.WriteResultWithStatusCode(w, response, response.Code)
	}
}

func pullCart(w http.ResponseWriter, req *http.Request) {
	fmt.Println("pullCart CALLED")
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

func addToUserCart(w http.ResponseWriter, r *http.Request) {
	urlCartId := r.URL.Query()["cartId"][0]
	var item CartItem
	_ = json.NewDecoder(r.Body).Decode(&item)
	var attributes []attribute.ItemAttribute
	for _, itemOptions := range item.Item.ProductOption.ExtensionAttributes.ConfigurableItemOptions  {
		attributes =  append(attributes, attribute.GetAttributeNameFromSolr(itemOptions.OptionsID, itemOptions.OptionValue))
	}
	sku := product.BuildSKUFromItemAttributes(attributes, item.Item.SKU)
	productFromSolr := product.GetProductFromSolrBySKU(sku)
	item.Item.SKU = productFromSolr.Sku
	item.Item.Price = productFromSolr.Price
	item.Item.ProductType = productFromSolr.TypeID
	item.Item.Name = productFromSolr.Name
	item.Item.ItemID = counter.GetAndIncreaseItemIdCounterInMongo()
	updateUserCartInMongo(urlCartId, item.Item)
}

func deleteFromUserCart(w http.ResponseWriter, r *http.Request)  {
	urlCartId := r.URL.Query()["cartId"][0]
	var item CartItem
	_ = json.NewDecoder(r.Body).Decode(&item)
	deleteItemFromUserCartInMongo(urlCartId, item)
}