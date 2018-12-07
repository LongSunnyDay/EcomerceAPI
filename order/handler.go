package order

import (
	"encoding/json"
	"go-api-ws/helpers"
	"net/http"
)

func PlaceOrder(w http.ResponseWriter, r *http.Request)  {
	var orderData PlaceOrderData
	err := json.NewDecoder(r.Body).Decode(&orderData)
	helpers.PanicErr(err)

}
