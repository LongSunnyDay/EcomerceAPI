package postNord

import (
	"github.com/go-chi/chi"
	"net/http"
)

func RouterPostNord() http.Handler {
	r := chi.NewRouter()
	r.Post("/getTransitTime", GetTransitTimeInformation)
	return r
}