package todoMongo

import (
	"net/http"
	"github.com/go-chi/chi"

			)



func TodoRouter() http.Handler{
	router := chi.NewRouter()
	router.Post("/todo", CreateTodo)
	router.Get("/todo/{todoID}", GetTodo)
	router.Get("/todo", ListTodos)
	router.Delete("/todo/{todoID}", RemoveTodo)
	router.Put("/todo/{todoID}", UpdateTodo)
	return router



}
