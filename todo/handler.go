package todo

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/xeipuuv/gojsonschema"
	"go-api-ws/config"
	"go-api-ws/helpers"
	"net/http"
)

var (
	todo  Todo
	todos []Todo
)

func createTodo(w http.ResponseWriter, r *http.Request) {
	var schemaLoader = gojsonschema.NewReferenceLoader("file://todo/jsonSchemaModels/todoSchema.json")
	err := json.NewDecoder(r.Body).Decode(&todo)
	helpers.PanicErr(err)
	documentLoader := gojsonschema.NewGoLoader(todo)
	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	helpers.PanicErr(err)

	if result.Valid() {
		writeToDoToDB()
		todos = append(todos, todo)
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode("Todo: " + todo.Title + " has been registered. ")
		helpers.PanicErr(err)
	} else {
		err = json.NewEncoder(w).Encode("There is and error creating todo:")
		helpers.PanicErr(err)
		fmt.Printf("The document is not valid. See errors :\n")
		for _, desc := range result.Errors() {
			fmt.Printf("- %s\n", desc)
		}
	}

}

func getTodo(w http.ResponseWriter, r *http.Request) {
	todoID, err := helpers.GetParameterFromUrl("todoID", r)
	helpers.PanicErr(err)

	queryErr := getToDoFromDB(todoID)

	if queryErr != nil {
		err = json.NewEncoder(w).Encode("Got an error: " + queryErr.Error())
		helpers.PanicErr(err)
		return
	}
	err = json.NewEncoder(w).Encode(todo)
	helpers.PanicErr(err)
}

func removeTodo(w http.ResponseWriter, r *http.Request) {
	todoID := chi.URLParam(r, "todoID")
	db, err := config.Conf.GetDb()
	helpers.PanicErr(err)

	queryErr := db.QueryRow("SELECT * FROM todos t WHERE id=?", todoID).
		Scan(&todo.ID, &todo.Title, &todo.Category, &todo.Content, &todo.Created, &todo.Modified, &todo.State)

	if queryErr != nil {
		err = json.NewEncoder(w).Encode("Got an error: " + queryErr.Error())
		helpers.PanicErr(err)
		return
	}
	_, err = db.Exec("DELETE u FROM todos t WHERE t.id=?", todoID)
	helpers.PanicErr(err)
	err = json.NewEncoder(w).Encode("Todo " + todo.Title + " has been deleted")
	helpers.PanicErr(err)
}

func editTodo(w http.ResponseWriter, r *http.Request) {

	var schemaLoader = gojsonschema.NewReferenceLoader("file://todo/jsonSchemaModels/todoSchema.json")
	_ = json.NewDecoder(r.Body).Decode(&todo)
	documentLoader := gojsonschema.NewGoLoader(todo)
	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	helpers.PanicErr(err)

	if result.Valid() {

		todoID := chi.URLParam(r, "todoID")
		db, err := config.Conf.GetDb()
		helpers.PanicErr(err)

		queryErr := db.QueryRow("SELECT * FROM todos t WHERE id=?", todoID).
			Scan(&todo.ID, &todo.Title, &todo.Category, &todo.Content, &todo.Created, &todo.Modified, &todo.State)

		if queryErr != nil {
			err = json.NewEncoder(w).Encode("Got an error: " + queryErr.Error())
			helpers.PanicErr(err)
			return
		}

		res, err := db.Exec("UPDATE todos t SET Title = ?, Category = ?, Content = ?, State = ?")
		fmt.Println(res)
		helpers.PanicErr(err)
		return
	}

}
