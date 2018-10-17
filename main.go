package main

import (
	"./user"
	"./config"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

func init()  {
	config.GetConfig("config.yml")
}

func main()  {
	r := user.UserRouter()
	http.ListenAndServe(":8080", r)
}
