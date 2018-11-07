package user

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/chi"
	"github.com/xeipuuv/gojsonschema"
	"go-api-ws/config"
	"go-api-ws/core"
	"go-api-ws/helpers"
	"net/http"
	"time"
)

var userModule core.ApiModule

func init() {
	userModule = core.ApiModule{
		Name:        "User module",
		Description: "User module. Supports username and email authentication. Categories are stored as a flat list.",
		Version:     "0.1",
		Author:      "Matas Cereskevicius @ JivaLabs",
	}

}

var users []User
var user User

//Get Order History
func getOrderHistory(w http.ResponseWriter, r *http.Request) {
	urlToken, err := helpers.GetTokenFromUrl(r)
	helpers.PanicErr(err)
	token, _ := jwt.Parse(urlToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There is an error")
		}
		return []byte(MySecret), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if claims.VerifyExpiresAt(time.Now().Unix(), true) {
			response := Response{
				Code:200,
				Result: map[string]interface{}{
					"items":[]string{},
					"search_criteria": "",
					"total_count": 0}}

			w.Header().Set("content-type", "application/json")
			w.WriteHeader(200)
			json.NewEncoder(w).Encode(response)
		}
	} else {
		json.NewEncoder(w).Encode("Invalid authorization token")
		fmt.Println("Invalid token")
	}
}

// Me endpoint
func meEndpoint(w http.ResponseWriter, r *http.Request) {
	urlToken, err := helpers.GetTokenFromUrl(r)
	helpers.PanicErr(err)
	token, _ := jwt.Parse(urlToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There is an error")
		}
		return []byte(MySecret), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if claims.VerifyExpiresAt(time.Now().Unix(), true) {
			//Some hard coded data to send back to client
			addresses := [0]string{}
			me := &MeUser{Code: 200, Result: Result{
				Addresses:              addresses,
				CreatedAt:              "2018-10-25 12:55:40",
				CreatedIn:              "Default Store View",
				DisableAutoGroupChange: 0,
				GroupID:                1,
				ID:                     1646,
				WebsiteID:              1,
				UpdatedAt:              "2018-10-29 08:44:13",
				StoreID:                1}}

			userFromDb := getUserDataFromDbById(claims["sub"])

			me.Result.FirstName = userFromDb.FirstName
			me.Result.LastName = userFromDb.LastName
			me.Result.Email = userFromDb.Email

			w.Header().Set("content-type", "application/json")
			w.WriteHeader(200)
			json.NewEncoder(w).Encode(me)
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

// RegisterUser function
func registerUser(w http.ResponseWriter, r *http.Request) {
	_ = json.NewDecoder(r.Body).Decode(&user)
	validationResult := helpers.CheckJSONSchemaWithGoStruct("file://user/jsonSchemaModels/userRegister.schema.json", user)

	if validationResult.Valid() {
		sendNewUserToDb(user)

		var response Response
		w.Header().Set("content-type", "application/json")
		response.Code = 200
		response.Result = "ok"
		json.NewEncoder(w).Encode(response)
	} else {

		w.Header().Set("content-type", "application/json")
		w.WriteHeader(400)

		var response Response
		response.Code = 400
		response.Result = validationResult.Errors()
		json.NewEncoder(w).Encode(response)

		fmt.Printf("The document is not valid. See errors :\n")
		for _, desc := range validationResult.Errors() {
			fmt.Printf("- %s\n", desc)
		}
	}
}

// GetUser function
func getUser(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")
	db, err := config.Conf.GetDb()
	helpers.PanicErr(err)
	queryErr := db.QueryRow("SELECT * FROM users u WHERE id=?", userID).
		Scan(&user.ID, &user.Customer.FirstName, &user.Customer.LastName, &user.Customer.Email, &user.Password)
	if queryErr != nil {
		json.NewEncoder(w).Encode("Got an error: " + queryErr.Error())
		return
	}
	json.NewEncoder(w).Encode(user)
}

// RemoveUser function
func removeUser(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")
	db, err := config.Conf.GetDb()
	helpers.PanicErr(err)
	queryErr := db.QueryRow("SELECT * FROM users u WHERE id=?", userID).
		Scan(&user.ID, &user.Customer.FirstName, &user.Customer.LastName, &user.Customer.Email, &user.Password)
	if queryErr != nil {
		json.NewEncoder(w).Encode("Got an error: " + queryErr.Error())
		return
	}
	db.Exec("DELETE u FROM users u WHERE u.id=?", userID)
	json.NewEncoder(w).Encode("User " + user.Customer.FirstName + " deleted")
}

// GetAllUsers function
func getAllUsers(w http.ResponseWriter, r *http.Request) {
	users = []User{}
	db, err := config.Conf.GetDb()
	helpers.PanicErr(err)
	rows, err := db.Query("SELECT * FROM users")
	helpers.PanicErr(err)
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(
			&user.ID,
			&user.Customer.FirstName,
			&user.Customer.LastName,
			&user.Customer.Email,
			&user.Password)
		helpers.PanicErr(err)
		users = append(users, user)
	}
	err = rows.Err()
	helpers.PanicErr(err)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// UpdateUser function
func updateUser(w http.ResponseWriter, r *http.Request) {
	var schemaLoader = gojsonschema.NewReferenceLoader("file://userRole/jsonSchemaModels/userUpdate.schema.json")
	var updatedUser User
	_ = json.NewDecoder(r.Body).Decode(&updatedUser)
	documentLoader := gojsonschema.NewGoLoader(updatedUser)
	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	helpers.PanicErr(err)

	if result.Valid() {
		db, err := config.Conf.GetDb()
		helpers.PanicErr(err)
		userID := chi.URLParam(r, "userID")
		err = db.QueryRow("SELECT * FROM users u WHERE id=?", userID).
			Scan(&user.ID, &user.Customer.FirstName, &user.Customer.LastName, &user.Customer.Email, &user.Password)
		if err != nil {
			json.NewEncoder(w).Encode("Got an error: " + err.Error())
			return
		}
		res, err := db.Exec("UPDATE users u SET First_name = ?, Last_name = ?, Email = ? WHERE ID = ?", updatedUser.Customer.FirstName, updatedUser.Customer.LastName, updatedUser.Customer.Email, userID)
		fmt.Println(res)
		helpers.PanicErr(err)
		return

	}
	json.NewEncoder(w).Encode(&User{})
}

// Update password
func updatePassword(w http.ResponseWriter, r *http.Request) {
	schemaLoader := gojsonschema.NewReferenceLoader("file://userRole/jsonSchemaModels/userPassword.schema.json")
	var password UpdatePassword

	_ = json.NewDecoder(r.Body).Decode(&password)
	fmt.Println(password)
	documentLoader := gojsonschema.NewGoLoader(password)
	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	helpers.PanicErr(err)

	if result.Valid() {
		db, err := config.Conf.GetDb()
		helpers.PanicErr(err)
		userID := chi.URLParam(r, "userID")
		err = db.QueryRow("SELECT Password FROM users u WHERE id=?", userID).
			Scan(&user.Password)
		if err != nil {
			json.NewEncoder(w).Encode("Got an error: " + err.Error())
			return
		}
		if checkPasswordHash(password.Password, user.Password) {

			password.NewPassword, err = hashPassword(password.NewPassword)
			_, err := db.Exec("UPDATE users u SET Password = ? WHERE ID = ?", password.NewPassword, userID)
			helpers.PanicErr(err)

		}
	} else {
		json.NewEncoder(w).Encode("There is and error updating your password:")
		fmt.Printf("The document is not valid. See errors :\n")
		for _, desc := range result.Errors() {
			fmt.Printf("- %s\n", desc)
		}
	}

}

// Order-history endpoint
func orderHistory(w http.ResponseWriter, r *http.Request) {
	urlToken := r.URL.Query()["token"][0]

	token, _ := jwt.Parse(urlToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There is an error")
		}
		return []byte(MySecret), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if claims["user_type"] == "adminRole" && claims.VerifyExpiresAt(time.Now().Unix(), true) {

		} else {
			json.NewEncoder(w).Encode("Not authorized")
		}
	} else {
		json.NewEncoder(w).Encode("Invalid authorization token")
	}

}
