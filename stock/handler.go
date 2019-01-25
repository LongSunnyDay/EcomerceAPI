package stock

import (
	"encoding/json"
	"go-api-ws/helpers"
	"net/http"
)

func checkStock(w http.ResponseWriter, r *http.Request) {
	itemSkuFromUrl, err := helpers.GetParameterFromUrl("sku", r)
	helpers.PanicErr(err)
	var item DataStock
	item.GetDataFromDbBySku(itemSkuFromUrl)
	response := helpers.Response{
		Code:   http.StatusOK,
		Result: item}
	response.SendResponse(w)
}

func insertToStock(w http.ResponseWriter, r *http.Request) {
	var stockData DataStock
	err := json.NewDecoder(r.Body).Decode(&stockData)
	helpers.PanicErr(err)
	stockData.insertDataToStock()
	helpers.WriteResultWithStatusCode(w, "ok", http.StatusOK)
}

func updateStockItem(w http.ResponseWriter, r *http.Request) {
	var stockData DataStock
	err := json.NewDecoder(r.Body).Decode(&stockData)
	helpers.PanicErr(err)
	stockData.updateDataInDb()
	helpers.WriteResultWithStatusCode(w, "ok", http.StatusOK)
}

func removeItemFromStock(w http.ResponseWriter, r *http.Request) {
	itemSkuFromUrl, err := helpers.GetParameterFromUrl("sku", r)
	helpers.PanicErr(err)
	removeItemFromDb(itemSkuFromUrl)
	helpers.WriteResultWithStatusCode(w, "ok", http.StatusOK)
}
