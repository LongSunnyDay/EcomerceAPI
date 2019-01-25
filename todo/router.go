package todo

import (
	"github.com/go-chi/chi"
	"net/http"
)

func RouterTodo() http.Handler {
	r := chi.NewRouter()
	r.Post("/todo", createTodo)
	r.Get("/todo/{todoID}", getTodo)
	r.Delete("/todo/{todoID}", removeTodo)
	return r
}
