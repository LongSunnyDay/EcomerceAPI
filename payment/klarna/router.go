package klarna

import (
	"github.com/go-chi/chi"
	"net/http"
)

func RouterKlarna() http.Handler {
	r := chi.NewRouter()
	r.Post("/sessions", CreateSession)
	r.Post("/order", CreateOrder)
	return r
}
