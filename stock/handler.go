package stock

import (
	"encoding/json"
	"go-api-ws/helpers"
	"go-api-ws/user"
	"net/http"
)

func checkStock(w http.ResponseWriter, r *http.Request)  {
	itemSkuFromUrl := r.URL.Query()["sku"][0]
	itemData := getDataFromDbBySku(itemSkuFromUrl)
	response := user.Response{
		Code:http.StatusOK,
		Result:itemData}
	helpers.WriteResultWithStatusCode(w, response, response.Code)
}

func insertToStock(w http.ResponseWriter, r *http.Request)  {
	var stockData stockData
	_ = json.NewDecoder(r.Body).Decode(&stockData)
	insertDataToStock(stockData)
}