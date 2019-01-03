package order

import (
	"github.com/go-chi/chi"
	"net/http"
)

func RouterOrder() http.Handler {
	r := chi.NewRouter()
	r.Post("/", PlaceOrder)
	r.Get("/order-history", GetCustomerOrderHistory)
	return r
}
