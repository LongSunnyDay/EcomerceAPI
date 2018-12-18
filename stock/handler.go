package stock

import (
	"encoding/json"
	"go-api-ws/helpers"
	"net/http"
)

func checkStock(w http.ResponseWriter, r *http.Request) {
	itemSkuFromUrl := r.URL.Query()["sku"][0]
	var item DataStock
	item.GetDataFromDbBySku(itemSkuFromUrl)
	response := helpers.Response{
		Code:   http.StatusOK,
		Result: item}
	response.SendResponse(w)
}

func insertToStock(w http.ResponseWriter, r *http.Request) {
	var stockData DataStock
	_ = json.NewDecoder(r.Body).Decode(&stockData)
	stockData.insertDataToStock()
	helpers.WriteResultWithStatusCode(w, "ok", http.StatusOK)
}

func updateStockItem(w http.ResponseWriter, r *http.Request) {
	var stockData DataStock
	_ = json.NewDecoder(r.Body).Decode(&stockData)
	stockData.updateDataInDb()
	helpers.WriteResultWithStatusCode(w, "ok", http.StatusOK)
}

func removeItemFromStock(w http.ResponseWriter, r *http.Request) {
	itemSkuFromUrl := r.URL.Query()["sku"][0]
	removeItemFromDb(itemSkuFromUrl)
	helpers.WriteResultWithStatusCode(w, "ok", http.StatusOK)
}
