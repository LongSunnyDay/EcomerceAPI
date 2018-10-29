package main

import (
	"go-api-ws/config"
	//"./user"
	_ "github.com/go-sql-driver/mysql"
	"net/http"

	"go-api-ws/todoMongo"
)

func init()  {
	config.GetConfig("config.yml")
}

func main()  {
	//r := user.UserRouter()
	r := todoMongo.TodoRouter()


	http.ListenAndServe(":8080", r)
}
