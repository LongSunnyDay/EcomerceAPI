package stock

import "net/http"

func checkStock(w http.ResponseWriter, r *http.Request)  {
	itemSkuFromUrl := r.URL.Query()["sku"][0]
	itemData := getDataFromDbBySku(itemSkuFromUrl)

}
