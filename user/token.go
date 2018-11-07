package user

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"go-api-ws/helpers"
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
	validationResult := helpers.CheckJSONSchemaWithGoStruct("file://user/jsonSchemaModels/userLogin.schema.json", userLogin)

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

			authToken := GetNewAuthToken(userFromDb.ID, role)
			authTokenString, err := authToken.SignedString([]byte(MySecret))
			helpers.PanicErr(err)

			refreshToken := GetNewRefreshToken(userFromDb.ID)
			refreshTokenString, err := refreshToken.SignedString([]byte(MySecret))
			helpers.PanicErr(err)

			response.Code = 200
			response.Result = authTokenString
			response.Meta = map[string]string{
				"refreshToken": refreshTokenString,
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

			authToken := GetNewAuthToken(claims["sub"].(string), role)
			authTokenString, err := authToken.SignedString([]byte(MySecret))
			helpers.PanicErr(err)

			refreshToken := GetNewRefreshToken(claims["sub"].(string))
			refreshTokenString, err := refreshToken.SignedString([]byte(MySecret))
			helpers.PanicErr(err)

			response := Response{
				Code:   200,
				Result: authTokenString,
				Meta: map[string]string{
					"refreshToken": refreshTokenString}}

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
