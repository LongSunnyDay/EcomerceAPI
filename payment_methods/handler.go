package payment_methods

import (
	"encoding/json"
	"fmt"
	"go-api-ws/helpers"
	"net/http"
)

func AddPaymentMethods(w http.ResponseWriter, r *http.Request) {
	var methods []Method
	err := json.NewDecoder(r.Body).Decode(&methods)
	helpers.PanicErr(err)
	validationResult := helpers.CheckJSONSchemaWithGoStruct(
		"file://payment_methods/jsonSchemaModels/add-payment-methods-Methods.schema.json",
		methods)
	if validationResult.Valid() {
		for _, method := range methods {
			method.insertToDb()
		}
		helpers.WriteResultWithStatusCode(w, "ok", http.StatusOK)
	} else {
		fmt.Println(validationResult.Errors())
		helpers.WriteResultWithStatusCode(w, validationResult.Errors(), http.StatusBadRequest)
	}
}

func GetPaymentMethods(w http.ResponseWriter, r *http.Request) {
	paymentMethods := getPaymentMethodsFromDb()
	response := helpers.Response{
		Code:   http.StatusOK,
		Result: paymentMethods}
	response.SendResponse(w)
}

func updatePaymentMethod(w http.ResponseWriter, r *http.Request) {
	var method Method
	err := json.NewDecoder(r.Body).Decode(&method)
	helpers.PanicErr(err)
	validationResult := helpers.CheckJSONSchemaWithGoStruct(
		"file://payment_methods/jsonSchemaModels/update-payment-methods-Method.schema.json",
		method)
	if validationResult.Valid() {
		method.updatePaymentMethodInDb()
		helpers.WriteResultWithStatusCode(w, "ok", http.StatusOK)
	} else {
		fmt.Println(validationResult.Errors())
		helpers.WriteResultWithStatusCode(w, validationResult.Errors(), http.StatusBadRequest)
	}
}

func removePaymentMethod(w http.ResponseWriter, r *http.Request) {
	id, err := helpers.GetParameterFromUrl("id", r)
	helpers.CheckErr(err)
	removePaymentMethodFromDb(id)
	helpers.WriteResultWithStatusCode(w, "ok", http.StatusOK)
}
