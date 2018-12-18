package currency

import (
	"github.com/go-chi/chi"
	"net/http"
)

func CurrencyRouter() http.Handler {
	r := chi.NewRouter()
	r.Post("/", createCurrency)
	r.Get("/{currencyID}", getCurrency)
	r.Get("/list", getCurrencyList)
	r.Delete("/{currencyID}", removeCurrency)
	r.Put("/{currencyID}", updateCurrency)
	return r
}
