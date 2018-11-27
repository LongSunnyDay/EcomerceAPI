package user

import (
	"github.com/go-chi/chi"
	"net/http"
)

func RouterUser() http.Handler {
	r := chi.NewRouter()
	r.Post("/create", registerUser)
	r.Post("/login", loginEndpoint)
	r.Get("/me", meEndpoint)
	r.Get("/order-history", getOrderHistory)
	r.Post("/refresh", refreshToken)
	r.Post("/me", updateUser)

	return r
}
