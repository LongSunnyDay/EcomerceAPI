package helpers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	fr "github.com/DATA-DOG/fastroute"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/pkg/errors"
	"github.com/xeipuuv/gojsonschema"
	_ "go/scanner"
	"log"
	"net/http"
	"strconv"
)

const (
	HeaderContentType   = "Content-Type"
	MIMEApplicationJSON = "application/json"
)

type (
	Response struct {
		Code   int         `json:"code,omitempty"`
		Result interface{} `json:"result,omitempty"`
		Meta   interface{} `json:"meta,omitempty"`
	}
)

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
	_, err = w.Write(jsonData)
	PanicErr(err)
}

func GetIntParameter(r *http.Request, name string) (int, error) {
	parameter := fr.Parameters(r).ByName(name)
	return strconv.Atoi(parameter)
}

func CheckJSONSchemaWithGoStruct(jsonSchemaPath string, structForValidation interface{}) *gojsonschema.Result {
	schemaForValidation := gojsonschema.NewReferenceLoader(jsonSchemaPath)
	documentToValidate := gojsonschema.NewGoLoader(structForValidation)
	validationResult, err := gojsonschema.Validate(schemaForValidation, documentToValidate)
	PanicErr(err)
	return validationResult
}

func GetTokenFromUrl(r *http.Request) (string, error) {
	token, err := GetParameterFromUrl("token", r)
	return token, err
}

func GetCartIdFromUrl(r *http.Request) (string, error) {
	cartId, err := GetParameterFromUrl("cartId", r)
	return cartId, err
}

func GetParameterFromUrl(paramName string, r *http.Request) (string, error) {
	param := r.URL.Query()[paramName][0]
	if len(param) > 1 {
		return param, nil
	}
	return "", errors.New(paramName + " not found in URl")
}

func WriteResultWithStatusCode(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set(HeaderContentType, MIMEApplicationJSON)
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	PanicErr(err)
}

func StructToBson(v interface{}) (doc *bson.Document, err error) {
	data, err := bson.Marshal(v)
	if err != nil {
		return
	}
	err = bson.Unmarshal(data, &doc)
	return
}

func (r Response) SendResponse(w http.ResponseWriter) {
	WriteResultWithStatusCode(w, r, r.Code)
}

func CheckIfRowExistsInMysql(dataBaseName *sql.DB, tableName string, fieldName string, fieldValue string) string {
	var value string
	res := dataBaseName.QueryRow("SELECT "+fieldName+" FROM "+tableName+" WHERE "+fieldName+" = ?", fieldValue)

	err := res.Scan(&value)

	PanicErr(err)

	return value
}
