package user

import (
	"github.com/kjk/betterguid"
	"go-api-ws/config"
	"go-api-ws/helpers"
	"net/http"
)

// Data models
type User struct {
	ID       string   `json:"id,omitempty"`
	Customer Customer `json:"customer,omitempty"`
	Password string   `json:"password,omitempty"`
	GroupId  int      `json:"group_id,omitempty"`
}

type Customer struct {
	FirstName string `json:"firstname,omitempty"`
	LastName  string `json:"lastname,omitempty"`
	Email     string `json:"email,omitempty"`
}

type UpdatePassword struct {
	Password    string `json:"password,omitempty"`
	NewPassword string `json:"newPassword,omitempty"`
}

type LoginForm struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

type Response struct {
	Acknowledged bool          `json:"acknowledged,omitempty"`
	Code         int           `json:"code,omitempty"`
	CreatedAt    string        `json:"created_at,omitempty"`
	Payload      *http.Request `json:"payload,omitempty"`
	Result       interface{}   `json:"result,omitempty"`
	ResultCode   int           `json:"result_code,omitempty"`
	TaskID       string        `json:"task_id,omitempty"`
	Transmited   bool          `json:"transmited,omitempty"`
	TransmitedAt string        `json:"transmited_at,omitempty"`
	UpdatedAt    string        `json:"updated_at,omitempty"`
	Url          string        `json:"url,omitempty"`
	Meta         interface{}   `json:"meta,omitempty"`
}

type MeUser struct {
	Code   int    `json:"code"`
	Result Result `json:"result,omitempty"`
}

type Result struct {
	Addresses              interface{} `json:"addresses"`
	CreatedAt              string      `json:"created_at,omitempty"`
	CreatedIn              string      `json:"created_in,omitempty"`
	DisableAutoGroupChange int         `json:"disable_auto_group_change,omitempty"`
	Email                  string      `json:"email,omitempty"`
	FirstName              string      `json:"firstname,omitempty"`
	GroupID                int         `json:"group_id,omitempty"`
	ID                     int         `json:"id,omitempty"`
	LastName               string      `json:"lastname,omitempty"`
	StoreID                int         `json:"store_id,omitempty"`
	UpdatedAt              string      `json:"updated_at,omitempty"`
	WebsiteID              int         `json:"website_id,omitempty"`
}

// Database operations
func insertUserIntoDb(user User) {
	id := betterguid.New()
	user.ID = id
	passwordHash, err := hashPassword(user.Password)
	helpers.PanicErr(err)
	user.Password = passwordHash
	db, err := config.Conf.GetDb()
	helpers.PanicErr(err)
	_, err = db.Exec("INSERT INTO users("+
		"ID, "+
		"First_name, "+
		"Last_name, "+
		"Email, "+
		"Password)"+
		" VALUES(?, ?, ?, ?, ?)",
		user.ID,
		user.Customer.FirstName,
		user.Customer.LastName,
		user.Customer.Email,
		user.Password)
	helpers.PanicErr(err)
}

func getUserFromDbByEmail(email string) (User) {
	db, err := config.Conf.GetDb()
	helpers.PanicErr(err)
	var userFromDb User
	err = db.QueryRow("SELECT ID, Password, Group_id FROM users u WHERE email = ?", email).Scan(&userFromDb.ID, &userFromDb.Password, &userFromDb.GroupId)
	helpers.PanicErr(err)

	return userFromDb
}

func getUserFromDbById(id string) (Customer) {
	db, err := config.Conf.GetDb()
	helpers.PanicErr(err)
	var userFromDb Customer
	err = db.QueryRow("SELECT First_name, Last_name, Email FROM users u WHERE ID = ?", id).Scan(&userFromDb.FirstName, &userFromDb.LastName, &userFromDb.Email)
	helpers.PanicErr(err)

	return userFromDb
}

func getGroupIdFromDbById(id string) (int) {
	db, err := config.Conf.GetDb()
	helpers.PanicErr(err)
	var groupId int
	err = db.QueryRow("SELECT Group_id FROM users u WHERE ID = ?", id).Scan(&groupId)
	helpers.PanicErr(err)

	return groupId
}
