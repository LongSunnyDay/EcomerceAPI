package todo

import (
	"github.com/go-chi/chi"
	"net/http"
)

func TodoRouter() http.Handler{
	r := chi.NewRouter()
	r.Post("/todo/create", createTodo)
	return r
}
