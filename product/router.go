package product

import (
	"github.com/go-chi/chi"
	"net/http"
)

func RouterProduct() http.Handler {
	r := chi.NewRouter()
	r.Post("/simple", InsertSimpleProductToDb)
	r.Get("/simple", getSimpleProductFromDb)
	r.Delete("/", deleteProductFromDb)
	r.Put("/simple", updateSimpleProductInDb)
	return r
}
