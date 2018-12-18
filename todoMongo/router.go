package todoMongo

import (
	"github.com/go-chi/chi"
	"net/http"
)

func TodoRouter() http.Handler {
	router := chi.NewRouter()
	router.Post("/create", CreateTodo)
	router.Get("/getTodo/{id}", GetTodo)
	router.Get("/getList", ListTodos)
	router.Delete("/remove/{id}", RemoveTodo)
	router.Put("/replace/{id}", ReplaceTodo)
	router.Put("/update/{id}", UpdateTodo)
	return router

}
