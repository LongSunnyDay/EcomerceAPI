package total

import (
	"github.com/go-chi/chi"
	"net/http"
)

func RoutesTotal() http.Handler {
	r := chi.NewRouter()
	r.Post("/shipping-information", GetTotalsWithPaymentMethods)
	r.Get("/totals", GetTotals)
	return r
}
