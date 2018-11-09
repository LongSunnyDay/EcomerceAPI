package todoMongo

import (
	"encoding/json"
	"go-api-ws/todoMongo/models"
	"net/http"
						)

var todos []models.Todo

//func ConnectionHandler(){
//	client, err := mongo.NewClient("mongodb://foo:bar@localhost:27017")
//	helpers.CheckErr(err)
//
//	err = client.Connect(context.TODO())
//	helpers.CheckErr(err)
//
//	collection := client.Database("go-api-ws").Collection("todos")
//	fmt.Println(collection)
//
//}

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


//func GetListTodos(w http.ResponseWriter, r *http.Request) {
//	payload := ListAllTodos()
//	json.NewEncoder(w).Encode(payload)
//}



//  gets a todo
//func GetTodo(w http.ResponseWriter, r *http.Request) {
//	urlTodoId := r.URL.Query()["todoId"][0]
//
//	payload := GetAllTodos()
//	for _, t := range payload {
//		if t.ID == urlTodoId {
//			json.NewEncoder(w).Encode(t)
//			return
//		}
//	}
//	json.NewEncoder(w).Encode("Todo not found")
//}

//getTodo test
func GetTodo(w http.ResponseWriter, r *http.Request)  {
	GetOneTodo()


}


func RemoveTodo(w http.ResponseWriter, r *http.Request) {
	urlTodoId := r.URL.Query()["todoId"][0]
	//_ = json.NewDecoder(r.Body).Decode(&todo)
	DeleteTodo(urlTodoId)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	urlTodoId := r.URL.Query()["todoId"][0]
	var todo models.Todo
	_ = json.NewDecoder(r.Body).Decode(&todo)
	UpdateTodoByID(todo, urlTodoId)
}
