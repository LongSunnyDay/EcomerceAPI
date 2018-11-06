package user

import (
	"github.com/go-chi/chi"
	"net/http"
)

//var userRoutes map[string]fr.Router

//func initUserRoutes() map[string]fr.Router {
////var routes = map[string]fr.Router{
////	"POST": fr.Chain(
////		fr.New("/user/register", RegisterUser)),
////}
////
////	return routes
//
//userRoutes = map[string]fr.Router{
//
//"POST":	fr.Chain(
//
//		fr.New("/api/product/list", RegisterUser),
//	),
//}
//return userRoutes
//}

func UserRouter() http.Handler {
	r := chi.NewRouter()
	r.Post("/create", registerUser)
	r.Post("/login", LoginEndpoint)
	r.Get("/me", meEndpoint)
	r.Get("/order-history", getOrderHistory)
	r.Post("/refresh", RefreshToken)

	return r
}
