//Marius Skaisgiris FT Jiva Labs colab 2018
package todo

import (
	"go-api-ws/config"
	"go-api-ws/helpers"
)

func writeToDoToDB() {
	db, err := config.Conf.GetDb()
	helpers.PanicErr(err)

	_, err = db.Exec("INSERT INTO todos("+
		"ID, "+
		"Title, "+
		"Category, "+
		"Content, "+
		"Modification_time, "+
		"State)"+
		" VALUES(?, ?, ?, ?, ?, ?)",
		todo.ID,
		todo.Title,
		todo.Category,
		todo.Content,
		todo.Modified,
		todo.State)
	helpers.PanicErr(err)
}

func getToDoFromDB(todoID string) error {
	db, err := config.Conf.GetDb()
	helpers.PanicErr(err)
	queryErr := db.QueryRow("SELECT * FROM todos t WHERE ID=?", todoID).
		Scan(&todo.ID, &todo.Title, &todo.Category, &todo.Content, &todo.Created, &todo.Modified, &todo.State)
	if queryErr != nil {
		return queryErr
	}
	return nil
}
