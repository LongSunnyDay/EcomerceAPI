package todoMongo

import (
	"github.com/go-chi/chi"
	"net/http"
)



func TodoRouter() http.Handler{
	router := chi.NewRouter()
	router.Post("/create", CreateTodo)
	router.Get("/getTodo", GetTodo)
	router.Get("/getList", ListTodos)
	router.Delete("/remove", RemoveTodo)
	router.Put("/update", UpdateTodo)
	return router

}
