package user

import (
	"go-api-ws/core"
	"magento_api/structs"
	"github.com/xeipuuv/gojsonschema"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"encoding/json"
	"github.com/kjk/betterguid"
	"fmt"
	"github.com/gorilla/mux"
	"go-api-ws/helpers"

)


var userModule core.ApiModule

func init() {
	userModule = core.ApiModule{
		Name: "User module",
		Description: "User module. Supports username and email authentication. Categories are stored as a flat list.",
		Version: "0.1",
		Author: "Remigijus Bauzys @ JivaLabs",
	}


}

var users []structs.User
var schemaLoader = gojsonschema.NewReferenceLoader("file://models/userRegister.schema.json")

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// RegisterUser function
func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user structs.User
	_ = json.NewDecoder(r.Body).Decode(&user)
	documentLoader := gojsonschema.NewGoLoader(user)
	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	helpers.CheckErr(err)

	if result.Valid() {

		id := betterguid.New()
		user.ID = id
		pswd, err := hashPassword(user.Password)
		helpers.CheckErr(err)

		user.Password = pswd
		users = append(users, user)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("User: " + user.Customer.FirstName + " has been registered")
	} else {
		json.NewEncoder(w).Encode("There is and error creating an user:")
		fmt.Printf("The document is not valid. See errors :\n")
		for _, desc := range result.Errors() {
			fmt.Printf("- %s\n", desc)
		}
	}
}

// GetUser function
func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, user := range users {
		if user.ID == params["id"] {
			json.NewEncoder(w).Encode(user)
			return
		}
	}
	json.NewEncoder(w).Encode(&structs.User{})
}

// RemoveUser function
func RemoveUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, user := range users {
		if user.ID == params["id"] {
			users = append(users[:index], users[index+1:]...)
			json.NewEncoder(w).Encode(users)
			return
		}
	}
	json.NewEncoder(w).Encode(&structs.User{})
}

// GetAllUsers function
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// UpdateUser function
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, user := range users {
		if user.ID == params["id"] {
			var updatedUser structs.User
			_ = json.NewDecoder(r.Body).Decode(&updatedUser)
			users[index] = updatedUser
			json.NewEncoder(w).Encode(updatedUser)
			return
		}
	}
	json.NewEncoder(w).Encode(&structs.User{})
}
