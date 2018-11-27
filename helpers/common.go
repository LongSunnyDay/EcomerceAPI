package helpers

import (
	"encoding/json"
	"fmt"
	fr "github.com/DATA-DOG/fastroute"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/pkg/errors"
	"github.com/xeipuuv/gojsonschema"
	"log"
	"net/http"
	"strconv"
)

type Response struct {
	Code   int         `json:"code,omitempty"`
	Result interface{} `json:"result,omitempty"`
	Meta   interface{} `json:"meta,omitempty"`
}

const HeaderContentType = "Content-Type"
const MIMEApplicationJSON = "application/json"

func PanicErr(err error) {
	if err != nil {
		panic(err)
	}
}

func CheckErr(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
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

func WriteResultWithStatusCode(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set(HeaderContentType, MIMEApplicationJSON)
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func StructToBson(v interface{}) (doc *bson.Document, err error) {
	data, err := bson.Marshal(v)
	if err != nil {
		return
	}
	err = bson.Unmarshal(data, &doc)
	return
}

func (r Response) SendResponse(w http.ResponseWriter)  {
	WriteResultWithStatusCode(w, r, r.Code)
}