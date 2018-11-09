package todoMongo

import (
	"encoding/json"
	"go-api-ws/todoMongo/models"
	"net/http"
	"github.com/go-chi/chi"
)

var todos []models.Todo

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	_ = json.NewDecoder(r.Body).Decode(&todo)
	InsertTodo(todo)
	json.NewEncoder(w).Encode(todo)
}

func ListTodos(w http.ResponseWriter, r *http.Request) {
	payload := GetAllTodos()
	json.NewEncoder(w).Encode(payload)
}

//getTodo test
func GetTodo(w http.ResponseWriter, r *http.Request)  {
	urlTodoId := chi.URLParam(r, "id")
	json.NewEncoder(w).Encode(GetOneTodo(urlTodoId))
}

func RemoveTodo(w http.ResponseWriter, r *http.Request) {
	urlTodoId := r.URL.Query()["todoId"][0]
	DeleteTodo(urlTodoId)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	urlTodoId := r.URL.Query()["todoId"][0]
	var todo models.Todo
	_ = json.NewDecoder(r.Body).Decode(&todo)
	UpdateTodoByID(todo, urlTodoId)
}
