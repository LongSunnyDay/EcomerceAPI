package cart

import (
	"github.com/go-chi/chi"
	"net/http"
)

func RouterCart() http.Handler {
	r := chi.NewRouter()
	r.Post("/create", createCart)
	r.Get("/pull", pullCart)
	r.Post("/payment-methods", addPaymentMethod)
	r.Get("/payment-methods", getPaymentMethods)
	r.Post("/update", updateCart)
	r.Post("/delete", deleteFromUserCart)

	return r
}