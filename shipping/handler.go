package shipping

import (
	"encoding/json"
	"go-api-ws/helpers"
	"net/http"
)

func AddShippingMethods(w http.ResponseWriter, r *http.Request) {
	var methods []Method
	err := json.NewDecoder(r.Body).Decode(&methods)
	helpers.PanicErr(err)
	validationResult := helpers.CheckJSONSchemaWithGoStruct("file://shipping/jsonSchemaModels/add-shipping-methods.json",
		methods)
	if validationResult.Valid() {
		for _, method := range methods {
			method.insertToDb()
		}
		helpers.WriteResultWithStatusCode(w, "ok", http.StatusOK)
	} else {
		helpers.WriteResultWithStatusCode(w, validationResult.Errors(), http.StatusBadRequest)
	}
}

func GetShippingMethods(w http.ResponseWriter, r *http.Request) {
	shippingMethods := getShippingMethodsFromDb()
	response := helpers.Response{
		Code:   http.StatusOK,
		Result: shippingMethods}
	response.SendResponse(w)
}

func updateShippingMethod(w http.ResponseWriter, r *http.Request) {
	var method Method
	err := json.NewDecoder(r.Body).Decode(&method)
	helpers.PanicErr(err)
	validationResult := helpers.CheckJSONSchemaWithGoStruct("file://shipping/jsonSchemaModels/update-shipping-method.json",
		method)
	if validationResult.Valid() {
		method.updatePaymentMethodInDb()
		helpers.WriteResultWithStatusCode(w, "ok", http.StatusOK)
	} else {
		helpers.WriteResultWithStatusCode(w, validationResult.Errors(), http.StatusBadRequest)
	}
}

func removePaymentMethod(w http.ResponseWriter, r *http.Request) {
	id, err := helpers.GetParameterFromUrl("id", r)
	helpers.PanicErr(err)
	removePaymentMethodFromDb(id)
	helpers.WriteResultWithStatusCode(w, "ok", http.StatusOK)
}
