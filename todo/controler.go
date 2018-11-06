//Marius Skaisgiris FT Jiva Labs colab 2018
package todo

import (
	c "../config"
	"../helpers"
	m "../todo/models"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/xeipuuv/gojsonschema"
	"net/http"
)

var todo m.Todo
var todos[]m.Todo

func createTodo(w http.ResponseWriter, r *http.Request){
	var schemaLoader = gojsonschema.NewReferenceLoader("file://todo/jsonSchemaModels/todoSchema.json")
	_ = json.NewDecoder(r.Body).Decode(&todo)
	documentLoader := gojsonschema.NewGoLoader(todo)
	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	helpers.PanicErr(err)

	if result.Valid(){
		db, err := c.Conf.GetDb()
		helpers.PanicErr(err)

		result, err := db.Exec("INSERT INTO todos(" +
			"ID, " +
			"Title, " +
			"Category, " +
			"Content, " +
			"Modification_time, " +
			"State)" +
			" VALUES(?, ?, ?, ?, ?, ?)",
			todo.ID,
			todo.Title,
			todo.Category,
			todo.Content,
			todo.Modified,
			todo.State)
		fmt.Println(result)
		helpers.PanicErr(err)

		todos = append(todos, todo)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("Todo: " + todo.Title + " has been registered. " )
	}	else {
		json.NewEncoder(w).Encode("There is and error creating todo:")
		fmt.Printf("The document is not valid. See errors :\n")
		for _, desc := range result.Errors() {
			fmt.Printf("- %s\n", desc)
		}
	}

}

func getTodo(w http.ResponseWriter, r *http.Request){
todoID := chi.URLParam(r, "todoID")
	db, err := c.Conf.GetDb()
	helpers.PanicErr(err)

	queryErr := db.QueryRow("SELECT * FROM todos t WHERE ID=?", todoID).
		Scan(&todo.ID, &todo.Title, &todo.Category, &todo.Content, &todo.Created, &todo.Modified, &todo.State)
	if queryErr != nil {
		json.NewEncoder(w).Encode("Got an error: " + queryErr.Error())
		return
	}
	json.NewEncoder(w).Encode(todo)

}

func removeTodo(w http.ResponseWriter, r *http.Request) {
	todoID := chi.URLParam(r, "todoID")
	db, err := c.Conf.GetDb()
	helpers.PanicErr(err)

	queryErr := db.QueryRow("SELECT * FROM todos t WHERE id=?", todoID).
		Scan(&todo.ID, &todo.Title, &todo.Category, &todo.Content, &todo.Created, &todo.Modified, &todo.State)

	if queryErr != nil {
		json.NewEncoder(w).Encode("Got an error: " + queryErr.Error())
		return
	}
	db.Exec("DELETE u FROM todos t WHERE t.id=?", todoID)
	json.NewEncoder(w).Encode("Todo " + todo.Title + " has been deleted")
}


func editTodo(w http.ResponseWriter, r *http.Request) {
	//var schemaLoader = gojsonschema.NewReferenceLoader("file://user/jsonSchemaModels/userUpdate.schema.json")
	//var updatedUser m.Todo
	//_ = json.NewDecoder(r.Body).Decode(&updatedUser)
	//documentLoader := gojsonschema.NewGoLoader(updatedUser)
	//result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	//helpers.PanicErr(err)

	var schemaLoader= gojsonschema.NewReferenceLoader("file://todo/jsonSchemaModels/todoSchema.json")
	_ = json.NewDecoder(r.Body).Decode(&todo)
	documentLoader := gojsonschema.NewGoLoader(todo)
	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	helpers.PanicErr(err)

	if result.Valid() {

	todoID := chi.URLParam(r, "todoID")
	db, err := c.Conf.GetDb()
	helpers.PanicErr(err)

	queryErr := db.QueryRow("SELECT * FROM todos t WHERE id=?", todoID).
		Scan(&todo.ID, &todo.Title, &todo.Category, &todo.Content, &todo.Created, &todo.Modified, &todo.State)

	if queryErr != nil {
		json.NewEncoder(w).Encode("Got an error: " + queryErr.Error())
		return
	}

	res, err := db.Exec("UPDATE todos t SET Title = ?, Category = ?, Content = ?, State = ?,")
		fmt.Println(res)
		helpers.PanicErr(err)
		return
}

}

func getAllTodos(){

}