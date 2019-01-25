package product

import (
	"encoding/json"
	"go-api-ws/helpers"
	"net/http"
)

func InsertSimpleProductToDb(w http.ResponseWriter, r *http.Request) {
	var simpleProduct SimpleProductStruct
	err := json.NewDecoder(r.Body).Decode(&simpleProduct)
	helpers.PanicErr(err)
	simpleProduct.insertSimpleProductToDb()
}

func getSimpleProductFromDb(w http.ResponseWriter, r *http.Request) {
	productSKU, err := helpers.GetParameterFromUrl("sku", r)
	helpers.PanicErr(err)
	product := getSimpleProductFromDbBySku(productSKU)
	response := helpers.Response{
		Code:   http.StatusOK,
		Result: product}
	response.SendResponse(w)
}

func deleteProductFromDb(w http.ResponseWriter, r *http.Request) {
	productSKU, err := helpers.GetParameterFromUrl("sku", r)
	helpers.PanicErr(err)
	removeProductFromDbBySku(productSKU)
	response := helpers.Response{
		Code: http.StatusOK}
	response.SendResponse(w)
}

func updateSimpleProductInDb(w http.ResponseWriter, r *http.Request) {
	var simpleProduct SimpleProductStruct
	err := json.NewDecoder(r.Body).Decode(&simpleProduct)
	helpers.PanicErr(err)
	rowsAffected := simpleProduct.updateProductInDb()
	response := helpers.Response{
		Code:   http.StatusOK,
		Result: rowsAffected}
	response.SendResponse(w)
}
