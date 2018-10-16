//package main
//
//import (
//	"flag"
//	"fmt"
//	fr "github.com/DATA-DOG/fastroute"
//	"github.com/labstack/gommon/log"
//	"./auth"
//	"./config"
//	"./core"
//	"./db"
//	"./helpers"
//	"golang.org/x/crypto/bcrypt"
//	"net/http"
//	"os"
//)
//
//var checkErr = helpers.CheckErr
//
//var version string
//var author string
//var routes map[string]fr.Router
//var router http.Handler
//
//var module core.ApiModule
//
//func init() {
//	module = core.ApiModule{
//		Author:      "Marius",
//		Version:     "0.1",
//		Name:        "Sample modules",
//		Description: "Sample module description.",
//	}
//
//	processFlags()
//
//	initAuth()
//
//	routes = initRoutes()
//
//	router = fr.RouterFunc(func(req *http.Request) http.Handler {
//		return routes[req.Method] // fastroute.Router is also http.Handler
//	})
//}
//
//func initAuth() {
//	auth.CreateUsersTableIfNotExists()
//}
//
//func initRoutes() map[string]fr.Router {
//
//	routes = map[string]fr.Router{
//		"GET": fr.Chain(
//			fr.New("/about", AboutHandler),
//
//			fr.New("/api/db/version", db.ShowVersionHandler),
//			fr.New("/api/db/databases", db.ShowDatabasesHandler),
//			fr.New("/api/db/tables", db.ShowTablesHandler),
//		),
//		"POST": fr.Chain(),
//	}
//	return routes
//}
//
//func processFlags() *FlagSettings {
//
//	osUser := os.Getenv("USER")
//	osGoPath := os.Getenv("PATH")
//
//	flags := NewFlags()
//	flag.StringVar(&flags.Config, "config", "config.yml", "config file (default is path/config.yaml|json|toml)")
//	flag.StringVar(&flags.AssetsPath, "assets-path", "assets", "Path to assets dir")
//	flag.StringVar(&flags.LogFile, "logFile", "/var/log/api-ws.log", "Log file")
//	flag.StringVar(&flags.Port, "port", ":9090", "http listen port")
//	flag.StringVar(&flags.Host, "host", "localhost", "http service host name")
//
//	fmt.Printf("User: %s Path: %s", osUser, osGoPath)
//	flag.Parse()
//
//	GetConfig(flags.Config)
//
//	//config.Conf.Port = flags.Port
//	//config.Conf.LogFile = flags.LogFile
//	//config.Conf.Host = flags.Host
//	//config.Conf.LogFile = flags.LogFile
//
//	return flags
//}
//
//func main() {
//
//	var router = fr.RouterFunc(func(req *http.Request) http.Handler {
//		return routes[req.Method] // fastroute.Router is also http.Handler
//	})
//
//	p := "MyPassword"
//	password := []byte(p)
//	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
//	if err != nil {
//
//	}
//
//	err = bcrypt.CompareHashAndPassword(hash, password)
//	if err == nil {
//		fmt.Printf("Password %s and hash %s matched.", password, hash)
//	} else {
//		fmt.Printf("Password does not match.")
//	}
//
//	fmt.Printf("\n⇨ http server started on http://%s%s\n", "localhost", Conf.Port)
//	log.Fatal(http.ListenAndServe(Conf.Port, router))
//}
//
//func AboutHandler(w http.ResponseWriter, req *http.Request) {
//	fmt.Fprintf(w, "This is a Jiva Labs minimalistic Product Service implementation. "+
//		"ver.: %s created by %s", version, author)
//}
package main

import (
	"./user"
	c "./config"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

func init()  {
	c.GetConfig("config.yml")
}

func main()  {
	//db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/go-api-ws")
	//helpers.CheckErr(err)
	//defer db.Close()
	r := user.UserRouter()
	http.ListenAndServe(":8080", r)
}
