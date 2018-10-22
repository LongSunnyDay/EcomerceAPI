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
var user m.User

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
	r.Put("/userp/{userID}", updatePassword)

	r.Post("/login", LoginEndpoint)
	r.Get("/protected", ProtectedEndpoint(getAllUsers))
	return r
}

// RegisterUser function
func registerUser(w http.ResponseWriter, r *http.Request) {
	schemaLoader := gojsonschema.NewReferenceLoader("file://user/models/userRegister.schema.json")
	_ = json.NewDecoder(r.Body).Decode(&user)
	documentLoader := gojsonschema.NewGoLoader(user)
	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	helpers.CheckErr(err)

	if result.Valid() {
		id := betterguid.New()
		user.ID = id
		user.Password, err = hashPassword(user.Password)
		helpers.CheckErr(err)
		db, err := c.Conf.GetDb()
		helpers.CheckErr(err)
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
		//json.NewEncoder(w).Encode(result)

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
	db, err := c.Conf.GetDb()
	helpers.CheckErr(err)
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
	db, err := c.Conf.GetDb()
	helpers.CheckErr(err)
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
	users =[]m.User{}
	db, err := c.Conf.GetDb()
	helpers.CheckErr(err)
	rows, err := db.Query("SELECT * FROM users")
	helpers.CheckErr(err)
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(
			&user.ID,
			&user.Customer.FirstName,
			&user.Customer.LastName,
			&user.Customer.Email,
			&user.Password)
			helpers.CheckErr(err)
		users = append(users, user)
	}
	err = rows.Err()
	helpers.CheckErr(err)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// UpdateUser function
func updateUser(w http.ResponseWriter, r *http.Request) {
	var schemaLoader = gojsonschema.NewReferenceLoader("file://user/models/userUpdate.schema.json")
	var updatedUser m.User
	_ = json.NewDecoder(r.Body).Decode(&updatedUser)
	documentLoader := gojsonschema.NewGoLoader(updatedUser)
	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	helpers.CheckErr(err)

	if result.Valid() {
		db, err := c.Conf.GetDb()
		helpers.CheckErr(err)
		userID  := chi.URLParam(r, "userID")
		err = db.QueryRow("SELECT * FROM users u WHERE id=?", userID).
			Scan(&user.ID, &user.Customer.FirstName, &user.Customer.LastName, &user.Customer.Email, &user.Password)
		if err != nil {
			json.NewEncoder(w).Encode("Got an error: " + err.Error())
			return
		}
		res, err := db.Exec("UPDATE users u SET First_name = ?, Last_name = ?, Email = ? WHERE ID = ?", updatedUser.Customer.FirstName, updatedUser.Customer.LastName, updatedUser.Customer.Email, userID)
		fmt.Println(res)
		helpers.CheckErr(err)
		return

	}
	json.NewEncoder(w).Encode(&m.User{})
}

func updatePassword(w http.ResponseWriter, r *http.Request){
	schemaLoader := gojsonschema.NewReferenceLoader("file://user/models/userPassword.schema.json")
	var password m.UpdatePassword

	_ = json.NewDecoder(r.Body).Decode(&password)
	fmt.Println(password)
	documentLoader := gojsonschema.NewGoLoader(password)
	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	helpers.CheckErr(err)

	if result.Valid(){
		db, err := c.Conf.GetDb()
		helpers.CheckErr(err)
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
			helpers.CheckErr(err)

		}
	} else {
		json.NewEncoder(w).Encode("There is and error updating your password:")
		fmt.Printf("The document is not valid. See errors :\n")
		for _, desc := range result.Errors() {
			fmt.Printf("- %s\n", desc)
		}
	}


}
