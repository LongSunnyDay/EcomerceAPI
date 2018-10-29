package todoMongo

import (
	"net/http"
	"encoding/json"
	"go-api-ws/todoMongo/models"

	"github.com/gorilla/mux"
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

//  gets a todo
func GetTodo(w http.ResponseWriter, r *http.Request) {


	params := mux.Vars(r)


	payload := GetAllTodos()
	for _, t := range payload {
		if t.ID == params["_id"] {
			json.NewEncoder(w).Encode(t)
			return
		}
	}
	json.NewEncoder(w).Encode("Todo not found")
}

func RemoveTodo(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	_ = json.NewDecoder(r.Body).Decode(&todo)
	DeleteTodo(todo)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	todoID := mux.Vars(r)["_id"]
	var todo models.Todo
	_ = json.NewDecoder(r.Body).Decode(&todo)
	UpdateTodoByID(todo, todoID)
}
