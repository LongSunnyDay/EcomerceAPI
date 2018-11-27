package payment

import (
	"github.com/go-chi/chi"
	"net/http"
)

func RouterPayment() http.Handler {
	r := chi.NewRouter()
	r.Post("/", AddPaymentMethods)
	r.Get("/", GetPaymentMethods)
	r.Put("/", updatePaymentMethod)
	r.Delete("/", removePaymentMethod)
	return r
}
