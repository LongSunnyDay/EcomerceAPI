package main

import (
	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
	"go-api-ws/cart"
	"go-api-ws/config"
	"go-api-ws/currency"
	"go-api-ws/language"
	"go-api-ws/order"
	"go-api-ws/payment/klarna"
	"go-api-ws/payment/paypal"
	"go-api-ws/payment_methods"
	"go-api-ws/postNord"
	"go-api-ws/product"
	"go-api-ws/shipping"
	"go-api-ws/stock"
	"go-api-ws/todoMongo"
	"go-api-ws/total"
	"go-api-ws/user"
	"net/http"
	"go-api-ws/discount"
)

func init() {
	config.GetConfig("config.yml")

}

func main() {

	klarna.GetClient()

	r := chi.NewRouter()
	r.Mount("/api/user", user.RouterUser())
	r.Mount("/api/cart", cart.RouterCart())
	r.Mount("/api/currency", currency.CurrencyRouter())
	r.Mount("/api/language", language.LanguageRouter())
	r.Mount("/api/todo", todoMongo.TodoRouter())
	r.Mount("/api/stock", stock.RouterStock())
	r.Mount("/api/payment-methods", payment_methods.RouterPayment())
	r.Mount("/api/shipping-methods", shipping.RoutesShippingMethods())
	r.Mount("/api/totals", total.RoutesTotal())
	r.Mount("/api/order", order.RouterOrder())
	r.Mount("/api/product", product.RouterProduct())
	r.Mount("/api/discount", discount.DiscountRouter())
	r.Mount("/api/coupon", discount.CouponRouter())
	r.Mount("/api/payment/paypal", paypal.RoutesPaypal())
	r.Mount("/api/payment/klarna", klarna.RouterKlarna())
	r.Mount("/api/postnord", postNord.RouterPostNord())
	http.ListenAndServe(":8080", r)
}
