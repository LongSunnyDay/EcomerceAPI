package paypal

import (
	"github.com/go-chi/chi"
	"net/http"
)

func RoutesPaypal() http.Handler {
	r := chi.NewRouter()
	r.Post("/create", Create)
	r.Post("/execute", Execute)
	return r
}
