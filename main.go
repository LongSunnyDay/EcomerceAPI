package main

import (
	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
	"go-api-ws/cart"
	"go-api-ws/config"
	"go-api-ws/todoMongo"
	"go-api-ws/user"
	"net/http"
)

func init()  {
	config.GetConfig("config.yml")
}

func main()  {

	r := chi.NewRouter()
	r.Mount("/api/user", user.RouterUser())
	r.Mount("/api/cart", cart.CartRouter())
	r.Mount("/api/todo", todoMongo.TodoRouter())
	http.ListenAndServe(":8080", r)
}
