package user

import (
	"context"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"go-api-ws/config"
	"go-api-ws/helpers"
	"net/http"
)

// Data models
type User struct {
	ID       string    `json:"id,omitempty"`
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
	Addresses              interface{} `json:"addresses" bson:"address"`
	CreatedAt              int64       `json:"created_at,omitempty" bson:"created_at"`
	CreatedIn              string      `json:"created_in,omitempty" bson:"created_in"`
	DisableAutoGroupChange int32       `json:"disable_auto_group_change,omitempty" bson:"disable_auto_group_change"`
	Email                  string      `json:"email,omitempty" bson:"email"`
	FirstName              string      `json:"firstname,omitempty" bson:"firstname"`
	GroupID                int32       `json:"group_id,omitempty" bson:"group_id"`
	ID                     string       `json:"id,omitempty" bson:"id"`
	LastName               string      `json:"lastname,omitempty" bson:"lastname"`
	StoreID                int32       `json:"store_id,omitempty" bson:"store_id"`
	UpdatedAt              int64       `json:"updated_at,omitempty" bson:"updated_at"`
	WebsiteID              int32       `json:"website_id,omitempty" bson:"website_id"`
}

type OrderHistory struct {
	Items          []Item `json:"items" bson:"items"`
	SearchCriteria string `json:"search_criteria" bson:"search_criteria"`
	TotalCount     int    `json:"total_count" bson:"omitempty"`
}

type Item struct {
	SKU string `json:"sku,omitempty" bson:"sku"`
}

// CONNECTIONSTRING DB connection string
const CONNECTIONSTRING = "mongodb://localhost:32768"

// DBNAME Database name
const DBNAME = "go-api-ws"

// COLLNAME Collection name
const COLLNAME = "users"

var db *mongo.Database

// Connect establish a connection to database
func init() {
	client, err := mongo.NewClient(CONNECTIONSTRING)
	helpers.PanicErr(err)

	err = client.Connect(context.Background())
	helpers.PanicErr(err)

	// Collection types can be used to access the database
	db = client.Database(DBNAME)
}

// Database operations

// MYSQL
func insertUserIntoDb(user User) {
	passwordHash, err := hashPassword(user.Password)
	helpers.PanicErr(err)
	user.Password = passwordHash
	db, err := config.Conf.GetDb()
	helpers.PanicErr(err)
	_, err = db.Exec("INSERT INTO users("+
		"Email, "+
		"Password)"+
		" VALUES(?, ?)",
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

func getUserIdFromDbByEmail(email string) (string) {
	db, err := config.Conf.GetDb()
	helpers.PanicErr(err)
	var id string
	err = db.QueryRow("SELECT ID FROM users u WHERE email = ?", email).Scan(&id)
	helpers.PanicErr(err)

	return id
}

func getUserFromDbById(id float64) (Customer) {
	db, err := config.Conf.GetDb()
	helpers.PanicErr(err)
	var userFromDb Customer
	err = db.QueryRow("SELECT First_name, Last_name, Email FROM users u WHERE ID = ?", id).Scan(&userFromDb.FirstName, &userFromDb.LastName, &userFromDb.Email)
	helpers.PanicErr(err)

	return userFromDb
}

func getGroupIdFromDbById(id int) (int) {
	db, err := config.Conf.GetDb()
	helpers.PanicErr(err)
	var groupId int
	err = db.QueryRow("SELECT Group_id FROM users u WHERE ID = ?", id).Scan(&groupId)
	helpers.PanicErr(err)

	return groupId
}

// MongoDB
func insertUserIntoMongo(userInfo Result) {
	_, err := db.Collection(COLLNAME).InsertOne(context.Background(),
		bson.NewDocument(
			bson.EC.String("type", "User info"),
			bson.EC.Int64("created_at", userInfo.CreatedAt),
			bson.EC.String("created_in", userInfo.CreatedIn),
			bson.EC.Int32("disable_auto_group_change", userInfo.DisableAutoGroupChange),
			bson.EC.String("email", userInfo.Email),
			bson.EC.String("firstname", userInfo.FirstName),
			bson.EC.String("lastname", userInfo.LastName),
			bson.EC.Int32("group_id", userInfo.GroupID),
			bson.EC.String("id", userInfo.ID),
			bson.EC.Int32("store_id", userInfo.StoreID),
			bson.EC.Int64("updated_at", userInfo.UpdatedAt),
			bson.EC.Int32("website_id", userInfo.WebsiteID),
			bson.EC.Interface("address", userInfo.Addresses)))
	helpers.PanicErr(err)
}

func getUserFromMongo(id string) (Result) {
	cur, err := db.Collection(COLLNAME).Find(context.Background(), bson.NewDocument(
		bson.EC.Interface("id", id),
		bson.EC.String("type", "User info")))
	helpers.PanicErr(err)
	var userInfo Result
	for cur.Next(context.Background()) {
		err := cur.Decode(&userInfo)
		helpers.PanicErr(err)
	}
	cur.Close(context.Background())
	return userInfo
}

func getUserOrderHistoryFromMongo(id string) (OrderHistory) {
	cur, err := db.Collection(COLLNAME).Find(context.Background(), bson.NewDocument(
		bson.EC.Interface("id", id),
		bson.EC.String("type", "Order history")))
	helpers.PanicErr(err)
	var orderHistory OrderHistory
	for cur.Next(context.Background()) {
		err := cur.Decode(&orderHistory)
		helpers.PanicErr(err)
	}
	cur.Close(context.Background())
	return orderHistory
}
