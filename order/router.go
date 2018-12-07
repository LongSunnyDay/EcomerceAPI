package order

import (
	"github.com/go-chi/chi"
	"net/http"
)

func RouterOrder() http.Handler {
	r := chi.NewRouter()
	r.Post("/", PlaceOrder)
	return r
}
