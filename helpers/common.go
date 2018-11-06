package helpers

import (
	"encoding/json"
	fr "github.com/DATA-DOG/fastroute"
	"github.com/pkg/errors"
	"github.com/xeipuuv/gojsonschema"
	"log"
	"net/http"
	"strconv"
)

const HeaderContentType = "Content-Type"
const MIMEApplicationJSON = "application/json"

func PanicErr(err error) {
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
	PanicErr(err)
	w.Header().Set(HeaderContentType, MIMEApplicationJSON)
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func GetIntParameter(r *http.Request, name string) (int, error) {
	parameter := fr.Parameters(r).ByName(name)
	return strconv.Atoi(parameter)
}

func CheckJSONSchemaWithGoStruct(jsonSchemaPath string, structForValidation interface{}) (*gojsonschema.Result) {
	schemaForValidation := gojsonschema.NewReferenceLoader(jsonSchemaPath)
	documentToValidate := gojsonschema.NewGoLoader(structForValidation)
	validationResult, err := gojsonschema.Validate(schemaForValidation, documentToValidate)
	PanicErr(err)
	return validationResult
}

func GetTokenFromUrl(r *http.Request) (string, error) {
	tokenFromUrl := r.URL.Query()["token"][0]
	if len(tokenFromUrl) > 1 {
		return tokenFromUrl, nil
	}

	return "", errors.New("Token not found")
}


