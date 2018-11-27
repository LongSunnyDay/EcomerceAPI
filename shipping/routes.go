package shipping

import (
	"github.com/go-chi/chi"
	"net/http"
)

func RoutesShippingMethods() http.Handler {
	r := chi.NewRouter()
	r.Post("/", AddShippingMethods)
	r.Get("/", getShippingMethods)
	r.Put("/", updateShippingMethod)
	r.Delete("/", removePaymentMethod)
	return r
}
