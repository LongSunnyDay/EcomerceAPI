package payment

import (
	"encoding/json"
	"fmt"
	"go-api-ws/helpers"
	"net/http"
)

func AddPaymentMethods(w http.ResponseWriter, r *http.Request) {
	var methods []method
	_ = json.NewDecoder(r.Body).Decode(&methods)
	validationResult := helpers.CheckJSONSchemaWithGoStruct(
		"file://payment/jsonSchemaModels/add-payment-methods.schema.json",
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
	var method method
	_ = json.NewDecoder(r.Body).Decode(&method)
	validationResult := helpers.CheckJSONSchemaWithGoStruct(
		"file://payment/jsonSchemaModels/update-payment-method.schema.json",
		method)
	if validationResult.Valid(){
		method.updatePaymentMethodInDb()
		helpers.WriteResultWithStatusCode(w, "ok", http.StatusOK)
	} else {
		fmt.Println(validationResult.Errors())
		helpers.WriteResultWithStatusCode(w, validationResult.Errors(), http.StatusBadRequest)
	}
}

func removePaymentMethod(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query()["id"][0]
	removePaymentMethodFromDb(id)
	helpers.WriteResultWithStatusCode(w, "ok", http.StatusOK)
}
