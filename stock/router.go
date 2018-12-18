package stock

import (
	"github.com/go-chi/chi"
	"net/http"
)

func RouterStock() http.Handler {
	r := chi.NewRouter()
	r.Get("/check", checkStock)
	r.Post("/", insertToStock)
	r.Put("/", updateStockItem)
	r.Delete("/", removeItemFromStock)
	return r
}
