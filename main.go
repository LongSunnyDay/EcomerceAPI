package main

import (
	"flag"
	"fmt"
	fr "github.com/DATA-DOG/fastroute"
	"go-api-ws/auth"
	"go-api-ws/config"
	"go-api-ws/core"
	"go-api-ws/db"
	"go-api-ws/helpers"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
)

var checkErr = helpers.CheckErr

var version string
var author string
var routes map[string]fr.Router
var router http.Handler

var module core.ApiModule

func init() {
	module = core.ApiModule{
		Author: "Marius",
		Version: "0.1",
		Name:"Sample modules",
		Description:"Sample module description.",
	}

	processFlags()

	initAuth()

	routes = initRoutes()

	fr.New("/api/db", db.DbHandler())
	router = fr.RouterFunc(func(req *http.Request) http.Handler {
		return routes[req.Method] // fastroute.Router is also http.Handler
	})
}

func initAuth() {
	auth.CreateUsersTableIfNotExists()
}

func initRoutes() map[string]fr.Router {

	routes = map[string]fr.Router{
		"GET": fr.Chain(
			fr.New("/about", AboutHandler),
		),
		"POST": fr.Chain(),
	}
	return routes
}

func processFlags() *config.FlagSettings {

	osUser := os.Getenv("USER")
	osGoPath := os.Getenv("PATH")

	flags := config.NewFlags()
	flag.StringVar(&flags.Config, "config", "config.yml", "config file (default is path/config.yaml|json|toml)")
	flag.StringVar(&flags.AssetsPath, "assets-path", "assets", "Path to assets dir")
	flag.StringVar(&flags.LogFile, "logFile", "/var/log/api-ws.log", "Log file")
	flag.StringVar(&flags.Port, "port", "9090", "http listen port")
	flag.StringVar(&flags.Host, "host", "localhost", "http service host name")

	fmt.Printf("User: %s Path: %s", osUser, osGoPath)
	flag.Parse()

	config.GetConfig(flags.Config)

	config.Conf.Port = flags.Port
	config.Conf.LogFile = flags.LogFile
	config.Conf.Port = flags.Port
	config.Conf.Port = flags.Port

	return flags
}

func main() {

	p := "MyPassword"
	password := []byte(p)
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {

	}

	err = bcrypt.CompareHashAndPassword(hash, password)
	if err == nil {
		fmt.Printf("Password %s and hash %s matched.", password, hash)
	} else {
		fmt.Printf("Password does not match.")
	}

	fmt.Printf("Server was started on http://localhost" + config.Conf.Port)
	http.ListenAndServe(config.Conf.Port, router)
}

func AboutHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "This is a Jiva Labs minimalistic Product Service implementation. "+
		"ver.: %s created by %s", version, author)
}
