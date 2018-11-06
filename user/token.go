package user

import (
	"../helpers"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"time"
)

const MySecret = "SenelisMegstaMociutesApvalumus"
const adminRole = "admin"
const userRole = "user"

func LoginEndpoint(w http.ResponseWriter, req *http.Request) {
	var userLogin LoginForm
	var response Response

	_ = json.NewDecoder(req.Body).Decode(&userLogin)
	validationResult := helpers.CheckJSONSchemaWithGoStruct("file://userRole/jsonSchemaModels/userLogin.schema.json", userLogin)

	pswd := userLogin.Password
	userLogin.Password = ""

	if validationResult.Valid() {
		userFromDb := getUserDataFromDbByEmail(userLogin.Username)

		if checkPasswordHash(pswd, userFromDb.Password) {
			userLogin.Password = ""
			var role string
			if userFromDb.GroupId < 1 {
				role = adminRole
			} else {
				role = userRole
			}
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"sub":  userFromDb.ID,
				"exp":  time.Now().Add(time.Hour * 1).Unix(),
				"role": role,
			})
			authToken, err := token.SignedString([]byte(MySecret))
			helpers.PanicErr(err)
			token = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"sub": userFromDb.ID,
				"exp": time.Now().Add(time.Hour * 4).Unix(),
			})
			refreshToken, err := token.SignedString([]byte(MySecret))
			helpers.PanicErr(err)
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
		response.Result = validationResult.Errors()
		json.NewEncoder(w).Encode(response)

		fmt.Printf("The document is not valid. See errors :\n")
		for _, desc := range validationResult.Errors() {
			fmt.Printf("- %s\n", desc)
		}
	}
}

func ProtectedEndpointMiddleware(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		urlToken := req.URL.Query()["token"][0]
		token, _ := jwt.Parse(urlToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("There is an error")
			}
			return []byte(MySecret), nil
		})
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if claims["role"] == "adminRole" && claims.VerifyExpiresAt(time.Now().Unix(), true) {
				handlerFunc.ServeHTTP(w, req)
			} else {
				response := Response{
					Code:   403,
					Result: "Token expired"}
				w.Header().Set("content-type", "application/json")
				w.WriteHeader(403)
				json.NewEncoder(w).Encode(response)
			}
		} else {
			response := Response{
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

			groupId := getGroupIdFromDbById(claims["sub"])

			var role string
			if groupId < 1 {
				role = adminRole
			} else {
				role = userRole
			}
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"sub":  claims["sub"],
				"exp":  time.Now().Add(time.Hour * 1).Unix(),
				"role": role,
			})
			newAuthToken, err := token.SignedString([]byte(MySecret))
			helpers.PanicErr(err)

			token = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"sub": claims["sub"],
				"exp": time.Now().Add(time.Hour * 4).Unix(),
			})
			newRefreshToken, err := token.SignedString([]byte(MySecret))
			helpers.PanicErr(err)

			response := Response{
				Code:   200,
				Result: newAuthToken,
				Meta: map[string]string{
					"refreshToken": newRefreshToken}}

			w.WriteHeader(200)
			w.Header().Set("content-type", "application/json")
			json.NewEncoder(w).Encode(response)
		} else {
			response := Response{
				Code:   403,
				Result: "Token expired"}
			w.Header().Set("content-type", "application/json")
			w.WriteHeader(403)
			json.NewEncoder(w).Encode(response)
		}
	} else {
		response := Response{
			Code:   400,
			Result: "Invalid token"}
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(response)
	}
}
