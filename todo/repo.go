package todo

import (
	"fmt"
	"go-api-ws/config"
	"go-api-ws/helpers"
)

func createTodo(todo *Todo) *Todo{
	db, err := config.Conf.GetDb()
	helpers.CheckErr(err)

	result, err := db.Exec("INSERT INTO Todo(Title, Category, State) VALUES(?, ?, ?)",
		todo.Title, todo.Category, todo.State)
	if err != nil {
		fmt.Println("ERROR saving to db - ", err)
	}

	Id64, err := result.LastInsertId()
	Id := int(Id64)
	todo = &Todo{Id: Id}

	db.QueryRow("SELECT State, Title, Category FROM Todo WHERE Id=?", todo.Id).
		Scan(&todo.State, &todo.Title, &todo.Category)
	return todo
}

func updateTodo(todoParams Todo) *Todo{

	db, err := config.Conf.GetDb()
	helpers.CheckErr(err)
	_, err = db.Exec("UPDATE Todo SET Title = ?, Category = ?, State = ? WHERE Id = ?", todoParams.Title, todoParams.Category, todoParams.State, todoParams.Id)

	if err != nil {
		fmt.Println("ERROR saving to db - ", err)
	}

	todo := &Todo{Id: todoParams.Id}
	err = db.QueryRow("SELECT State, Title, Category FROM Todo WHERE Id=?", todo.Id).Scan(&todo.State, &todo.Title, &todo.Category)
	if err != nil {
		fmt.Println("ERROR reading from db - ", err)
	}

	return &todoParams
}
