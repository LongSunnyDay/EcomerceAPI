package discount

import (
	"net/http"
	"github.com/go-chi/chi"
)

func DiscountRouter() http.Handler {
	r := chi.NewRouter()
	r.Post("/", createDiscount)
	r.Get("/{discountID}", getDiscount)
	r.Get("/list", getDiscountList)
	r.Delete("/{discountID}", removeDiscount)
	r.Put("/{discountID}", updateDiscount)
	return r
}