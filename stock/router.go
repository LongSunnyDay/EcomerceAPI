package stock

import (
	"github.com/go-chi/chi"
	"net/http"
)

func RouterStock() http.Handler {
	r := chi.NewRouter()
	r.Get("/check", checkStock)
	r.Post("/insert", insertToStock)
	return r
}