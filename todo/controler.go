//Marius Skaisgiris FT Jiva Labs colab 2018
package todo

import (
	"net/http"
	"github.com/xeipuuv/gojsonschema"
	"encoding/json"
	"go-api-ws/helpers"
	m "go-api-ws/todo/models"
	c "go-api-ws/config"
	"fmt"
)

var todo m.Todo
var todos[]m.Todo

func createTodo(w http.ResponseWriter, r *http.Request){
	var schemaLoader = gojsonschema.NewReferenceLoader("file://todo/models/todoCreate.json")
	_ = json.NewDecoder(r.Body).Decode(&todo)
	documentLoader := gojsonschema.NewGoLoader(todo)
	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	helpers.CheckErr(err)

	if result.Valid(){
		db, err := c.Conf.GetDb()
		helpers.CheckErr(err)

		result, err := db.Exec("INSERT INTO todos(" +
			"ID, " +
			"Title, " +
			"Category, " +
			"Content, " +
			"Creation_time, " +
			"Modification_time, " +
			"State)" +
			" VALUES(?, ?, ?, ?, ?, ?, ?)",
			todo.ID,
			todo.Title,
			todo.Category,
			todo.Content,
			todo.Created,
			todo.Modified,
			todo.State)
		fmt.Println(result)
		helpers.CheckErr(err)

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