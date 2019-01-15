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

func CouponRouter() http.Handler {
	r := chi.NewRouter()
	r.Post("/", createCoupon)
	//r.Get("/{couponID}", getCoupon)
	//r.Get("/list", getCouponList)
	//r.Delete("/{couponID}", removeCoupon)
	//r.Put("/{couponID}", updateCoupon)
	return r
}