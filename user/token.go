package user

import (
	"../config"
	"../helpers"
	m "./models"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/xeipuuv/gojsonschema"
	"net/http"
	"time"
)

var mySecret = "senelismegstamociutesapvalumus"

func LoginEndpoint(w http.ResponseWriter, req *http.Request) {
	var user m.LoginForm
	var dbData m.User
	_ = json.NewDecoder(req.Body).Decode(&user)
	schemaLoader := gojsonschema.NewReferenceLoader("file://user/models/userLogin.schema.json")
	documentLoader := gojsonschema.NewGoLoader(user)
	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	pswd := user.Password
	user.Password = ""
	helpers.CheckErr(err)
	if result.Valid() {
		db, err := config.Conf.GetDb()
		helpers.CheckErr(err)
		err = db.QueryRow("SELECT ID, Password FROM users u WHERE email = ?", user.Email).Scan(&dbData.ID, &dbData.Password)
		if err != nil {
			json.NewEncoder(w).Encode("Got an error: " + err.Error())
			return
		}

		if checkPasswordHash(pswd, dbData.Password) {
			user.Password = ""
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"id":  dbData.ID,
				"exp": time.Now().Add(time.Hour * 1).Unix(),
				"user_type": "admin",
			})
			tokenString, err := token.SignedString([]byte(mySecret))
			helpers.CheckErr(err)
			json.NewEncoder(w).Encode(tokenString)
		} else {
			json.NewEncoder(w).Encode("Invalid password")
		}
	} else {
		json.NewEncoder(w).Encode("There is and error:")
		fmt.Printf("The document is not valid. See errors :\n")
		for _, desc := range result.Errors() {
			fmt.Printf("- %s\n", desc)
		}
	}
}

func ProtectedEndpoint(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		params := request.Header.Get("Authorization")
		paramsFmt := params[7:]
		fmt.Println(paramsFmt)
		token, _ := jwt.Parse(paramsFmt, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("There is an error")
			}
			return []byte(mySecret), nil
		})
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if claims["user_type"] == "admin" && claims.VerifyExpiresAt(time.Now().Unix(), true){
				handlerFunc.ServeHTTP(writer, request)
			} else {
				json.NewEncoder(writer).Encode("Not authorized")
			}
		} else {
			json.NewEncoder(writer).Encode("Invalid authorization token")
		}
	}

}
