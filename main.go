package main

import (
	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
	"go-api-ws/cart"
	"go-api-ws/config"
	"go-api-ws/payment"
	"go-api-ws/shipping"
	"go-api-ws/stock"
	"go-api-ws/todoMongo"
	"go-api-ws/total"
	"go-api-ws/user"
	"net/http"
	"go-api-ws/currency"
	"go-api-ws/language"
)

func init()  {
	config.GetConfig("config.yml")

}

func main()  {

	r := chi.NewRouter()
	r.Mount("/api/user", user.RouterUser())
	r.Mount("/api/cart", cart.RouterCart())
	r.Mount("/api/currency", currency.CurrencyRouter())
	r.Mount("/api/language", language.LanguageRouter())
	r.Mount("/api/todo", todoMongo.TodoRouter())
	r.Mount("/api/stock", stock.RouterStock())
	r.Mount("/api/payment-methods", payment.RouterPayment())
	r.Mount("/api/shipping-methods", shipping.RoutesShippingMethods())
	r.Mount("/api/shipping-information", total.RoutesTotal())
	http.ListenAndServe(":8080", r)
}
