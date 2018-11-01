package user

import (
	"../config"
	"../helpers"
	m "../models"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/xeipuuv/gojsonschema"
	"net/http"
	"time"
)

var MySecret = "SenelisMegstaMociutesApvalumus"

func LoginEndpoint(w http.ResponseWriter, req *http.Request) {
	var userLogin m.LoginForm
	var dbData m.User
	var response m.Response
	_ = json.NewDecoder(req.Body).Decode(&userLogin)
	schemaLoader := gojsonschema.NewReferenceLoader("file://models/userLogin.schema.json")
	documentLoader := gojsonschema.NewGoLoader(userLogin)
	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	pswd := userLogin.Password
	userLogin.Password = ""
	helpers.CheckErr(err)
	if result.Valid() {
		db, err := config.Conf.GetDb()
		helpers.CheckErr(err)
		err = db.QueryRow("SELECT ID, Password, Group_id FROM users u WHERE email = ?", userLogin.Username).Scan(&dbData.ID, &dbData.Password, &dbData.GroupId)
		if err != nil {
			w.Header().Set("content-type", "application/json")
			w.WriteHeader(404)
			response.Code = 404
			response.Result = err.Error()
			json.NewEncoder(w).Encode(response)
			return
		}

		if checkPasswordHash(pswd, dbData.Password) {
			userLogin.Password = ""
			var role string
			if dbData.GroupId < 1 {
				role = "admin"
			} else {
				role = "user"
			}
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"sub":  dbData.ID,
				"exp":  time.Now().Add(time.Hour * 1).Unix(),
				"role": role,
			})
			authToken, err := token.SignedString([]byte(MySecret))
			helpers.CheckErr(err)
			token = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"sub": dbData.ID,
				"exp": time.Now().Add(time.Hour * 4).Unix(),
			})
			refreshToken, err := token.SignedString([]byte(MySecret))
			helpers.CheckErr(err)
			response.Code = 200
			response.Result = authToken
			response.Meta = map[string]string{
				"refreshToken": refreshToken,
			}
			json.NewEncoder(w).Encode(response)
		} else {
			w.Header().Set("content-type", "application/json")
			w.WriteHeader(401)
			response.Code = 401
			response.Result = "Password is invalid"
			json.NewEncoder(w).Encode(response)

			fmt.Println("Password is invalid")
		}
	} else {

		w.Header().Set("content-type", "application/json")
		w.WriteHeader(400)

		response.Code = 400
		response.Result = result.Errors()
		json.NewEncoder(w).Encode(response)

		fmt.Printf("The document is not valid. See errors :\n")
		for _, desc := range result.Errors() {
			fmt.Printf("- %s\n", desc)
		}
	}
}

func ProtectedEndpoint(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		urlToken := req.URL.Query()["token"][0]
		token, _ := jwt.Parse(urlToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("There is an error")
			}
			return []byte(MySecret), nil
		})
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if claims["role"] == "admin" && claims.VerifyExpiresAt(time.Now().Unix(), true) {
				handlerFunc.ServeHTTP(w, req)
			} else {
				response := m.Response{
					Code:   403,
					Result: "Token expired"}
				w.Header().Set("content-type", "application/json")
				w.WriteHeader(403)
				json.NewEncoder(w).Encode(response)
			}
		} else {
			response := m.Response{
				Code:   400,
				Result: "Invalid token"}
			w.Header().Set("content-type", "application/json")
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(response)
		}
	}

}

func RefreshToken(w http.ResponseWriter, req *http.Request) {
	var jsonBody map[string]string
	_ = json.NewDecoder(req.Body).Decode(&jsonBody)
	token, _ := jwt.Parse(jsonBody["refreshToken"], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There is an error")
		}
		return []byte(MySecret), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if claims.VerifyExpiresAt(time.Now().Unix(), true) {
			db, err := config.Conf.GetDb()
			helpers.CheckErr(err)
			var groupId int
			err = db.QueryRow("SELECT Group_id FROM users u WHERE ID = ?", claims["sub"]).Scan(&groupId)
			if err != nil {
				w.Header().Set("content-type", "application/json")
				w.WriteHeader(404)
				response := m.Response{
					Code:   404,
					Result: err.Error()}
				json.NewEncoder(w).Encode(response)
			}
			var role string
			if groupId < 1 {
				role = "admin"
			} else {
				role = "user"
			}
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"sub":  claims["sub"],
				"exp":  time.Now().Add(time.Hour * 1).Unix(),
				"role": role,
			})
			newAuthToken, err := token.SignedString([]byte(MySecret))
			helpers.CheckErr(err)

			token = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"sub": claims["sub"],
				"exp": time.Now().Add(time.Hour * 4).Unix(),
			})
			newRefreshToken, err := token.SignedString([]byte(MySecret))
			helpers.CheckErr(err)

			response := m.Response{
				Code:   200,
				Result: newAuthToken,
				Meta: map[string]string{
					"refreshToken": newRefreshToken}}

			w.WriteHeader(200)
			w.Header().Set("content-type", "application/json")
			json.NewEncoder(w).Encode(response)
		} else {
			response := m.Response{
				Code:   403,
				Result: "Token expired"}
			w.Header().Set("content-type", "application/json")
			w.WriteHeader(403)
			json.NewEncoder(w).Encode(response)
		}
	} else {
		response := m.Response{
			Code:   400,
			Result: "Invalid token"}
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(response)
	}
}
