package discount

import (
	"github.com/go-chi/chi"
	"net/http"
)

func RouterDiscount() http.Handler {
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
	r.Post("/", applyCoupon)
	r.Get("/{couponID}", getCoupon)
	r.Get("/list", getCouponList)
	r.Delete("/{couponID}", removeCoupon)
	r.Put("/{couponID}", updateCoupon)
	return r
}