package cart

import (
	"github.com/go-chi/chi"
	"net/http"
)

func RouterCart() http.Handler {
	r := chi.NewRouter()
	r.Post("/create", createCart)
	r.Get("/pull", pullCart)

	return r
}