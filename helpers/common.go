package helpers

import (
	"encoding/json"
	fr "github.com/DATA-DOG/fastroute"
	"log"
	"net/http"
	"strconv"
)

const HeaderContentType = "Content-Type"
const MIMEApplicationJSON = "application/json"

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func HttpError(err error, w http.ResponseWriter) {
	if err != nil {
		log.Printf(err.Error())
		a := ""
		w.Write([]byte(a))
	}
}

func WriteJsonResult(w http.ResponseWriter, data interface{}) {
	jsonData, err := json.Marshal(data)
	CheckErr(err)
	w.Header().Set(HeaderContentType, MIMEApplicationJSON)
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func GetIntParameter(r *http.Request, name string) (int, error) {
	parameter := fr.Parameters(r).ByName(name)
	return strconv.Atoi(parameter)
}
