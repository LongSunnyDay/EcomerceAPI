package klarna

import (
	"github.com/go-chi/chi"
	"net/http"
)

func RouterKlarna() http.Handler {
	r := chi.NewRouter()
	r.Post("/sessions", CreateSession)
	return r
}
