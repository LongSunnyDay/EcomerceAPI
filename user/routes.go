package user

import (
	"github.com/go-chi/chi"
	"net/http"
)

func UserRouter() http.Handler {
	r := chi.NewRouter()
	r.Post("/create", registerUser)
	r.Post("/login", LoginEndpoint)
	r.Get("/me", meEndpoint)
	r.Get("/order-history", getOrderHistory)
	r.Post("/refresh", RefreshToken)
	return r
}
