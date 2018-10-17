//Marius Skaisgiris FT Jiva Labs colab 2018
package todo

import (
	"go-api-ws/config"
	"go-api-ws/helpers"
)

func createTodom (todo *Todo) *Todo{
	db, err := config.Conf.GetDb()
	helpers.CheckErr(err)

	result, err := db.Exec("INSERT INTO Todo (Title, Categry, State) VALUES(?, ?, ?)",
		todo.Title, todo.Category, todo.State)
	helpers.CheckErr(err)
}
