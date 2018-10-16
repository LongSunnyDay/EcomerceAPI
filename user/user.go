package user

import (
	c "../config"
	"../core"
	"../helpers"
	m "./models"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/kjk/betterguid"
	"github.com/xeipuuv/gojsonschema"
	"golang.org/x/crypto/bcrypt"
	"net/http"
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

var users []m.User
var schemaLoader = gojsonschema.NewReferenceLoader("file://user/models/userRegister.schema.json")
var db, _ = c.Conf.GetDb()

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func UserRouter() http.Handler {
	r := chi.NewRouter()
	r.Post("/user/register", registerUser)
	r.Get("/user/list", getAllUsers)
	r.Get("/user/{userID}", getUser)
	r.Delete("/user/{userID}", removeUser)
	r.Put("/user/{userID}", updateUser)
	return r
}

// RegisterUser function
func registerUser(w http.ResponseWriter, r *http.Request) {
	var user m.User
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



		result, err := db.Exec("INSERT INTO users(" +
			"ID, " +
			"First_name, " +
			"Last_name, " +
			"Email, " +
			"Password)" +
			" VALUES(?, ?, ?, ?, ?)",
			user.ID,
			user.Customer.FirstName,
			user.Customer.LastName,
			user.Customer.Email,
			user.Password)
		fmt.Println(result)

		users = append(users, user)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("User: " + user.Customer.FirstName + " has been registered. ID: " + user.ID)
	} else {
		json.NewEncoder(w).Encode("There is and error creating an user:")
		fmt.Printf("The document is not valid. See errors :\n")
		for _, desc := range result.Errors() {
			fmt.Printf("- %s\n", desc)
		}
	}
}

// GetUser function
func getUser(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")
	for _, user := range users {
		if user.ID == userID {
			json.NewEncoder(w).Encode(user)
			return
		}
	}
	json.NewEncoder(w).Encode(&m.User{})
}

// RemoveUser function
func removeUser(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")
	for index, user := range users {
		if user.ID == userID {
			users = append(users[:index], users[index+1:]...)
			json.NewEncoder(w).Encode(users)
			return
		}
	}
	json.NewEncoder(w).Encode(&m.User{})
}

// GetAllUsers function
func getAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// UpdateUser function
func updateUser(w http.ResponseWriter, r *http.Request) {
	var updatedUser m.User
	_ = json.NewDecoder(r.Body).Decode(&updatedUser)
	documentLoader := gojsonschema.NewGoLoader(updatedUser)
	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	helpers.CheckErr(err)

	if result.Valid() {
		userID  := chi.URLParam(r, "userID")
		for index, user := range users {
			if user.ID == userID {
				users[index] = updatedUser
				json.NewEncoder(w).Encode(updatedUser)
				return
			}
		}
	}
	json.NewEncoder(w).Encode(&m.User{})
}
