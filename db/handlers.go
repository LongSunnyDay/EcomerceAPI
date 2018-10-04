package db

import (
	fr "github.com/DATA-DOG/fastroute"
	"go-api-ws/helpers"
	"net/http"
)

func InitHandler(handler http.Handler) http.HandlerFunc {
	return nil
}

type ErrorResponse struct {
	Attributes map[string]string `json:"@attributes"`
	Code       int               `json:"code,omitempty"`
	Message    string            `json:"message,omitempty"`
}

// /api/db/version
func ShowVersionHandler(w http.ResponseWriter, r *http.Request) {
	// user.CreatedAt = time.New().Local()
	version, err := getVersion()
	helpers.HttpError(err, w)
	helpers.WriteJsonResult(w, version)
}

// /api/db/databases
func ShowDatabasesHandler(w http.ResponseWriter, r *http.Request) {
	databases, err := showDatabases()
	helpers.HttpError(err, w)
	helpers.WriteJsonResult(w, databases)
}

// /api/db/tables?database=bibubabu
func ShowTablesHandler(w http.ResponseWriter, r *http.Request) {
	tables, err := showTables()
	helpers.HttpError(err, w)
	helpers.WriteJsonResult(w, tables)
}

func appendHandlers(router *fr.Router) fr.Router {
	return nil
}
