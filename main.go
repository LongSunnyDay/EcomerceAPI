package main

import (
	"./cart"
	"./config"
	"./todoMongo"
	"./user"
	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

func init()  {
	config.GetConfig("config.yml")
}

func main()  {

	r := chi.NewRouter()
	r.Mount("/api/user", user.UserRouter())
	r.Mount("/api/cart", cart.CartRouter())
	r.Mount("/api/todo", todoMongo.TodoRouter())
	http.ListenAndServe(":8080", r)
}
