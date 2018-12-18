package cart

import (
	"github.com/go-chi/chi"
	"go-api-ws/payment"
	"go-api-ws/shipping"
	"net/http"
)

func RouterCart() http.Handler {
	r := chi.NewRouter()
	r.Post("/create", createCart)
	r.Get("/pull", pullCart)
	r.Post("/update", updateCart)
	r.Post("/delete", deleteFromUserCart)
	r.Post("/payment-methods", payment.AddPaymentMethods)
	r.Get("/payment-methods", payment.GetPaymentMethods)
	r.Post("/shipping-methods", shipping.GetShippingMethods)
	return r
}
