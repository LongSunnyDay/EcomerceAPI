package user

import (
	"context"
	"fmt"
	"github.com/mongodb/mongo-go-driver/bson"
	"go-api-ws/config"
	"go-api-ws/helpers"
	"time"
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

type UpdatedCustomer struct {
	UpdateUser CustomerData `json:"customer,omitempty" bson:"customer,omitempty"`
}

type CustomerData struct {
	Addresses              []Address `json:"addresses" bson:"address"`
	CreatedAt              time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	CreatedIn              string    `json:"created_in,omitempty" bson:"created_in,omitempty"`
	DisableAutoGroupChange int32     `json:"disable_auto_group_change,omitempty" bson:"disable_auto_group_change,omitempty"`
	Email                  string    `json:"email,omitempty" bson:"email,omitempty"`
	FirstName              string    `json:"firstname,omitempty" bson:"firstname,omitempty"`
	GroupID                int32     `json:"group_id,omitempty" bson:"group_id,omitempty"`
	ID                     string    `json:"id,omitempty" bson:"id,omitempty"`
	LastName               string    `json:"lastname,omitempty" bson:"lastname,omitempty"`
	StoreID                int32     `json:"store_id,omitempty" bson:"store_id,omitempty"`
	UpdatedAt              time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	WebsiteID              int32     `json:"website_id,omitempty" bson:"website_id,omitempty"`
	DefaultShipping        string    `json:"default_shipping,omitempty" bson:"default_shipping,omitempty"`
}

type Address struct {
	ID              int64    `json:"id,omitempty" bson:"id,omitempty"`
	CustomerID      int      `json:"customer_id,omitempty" bson:"customer_id,omitempty"`
	Region          Region   `json:"region" bson:"region,omitempty"`
	RegionID        int      `json:"region_id,omitempty" bson:"region_id,omitempty"`
	CountryID       string   `json:"country_id,omitempty" bson:"country_id,omitempty"`
	Street          []string `json:"street,omitempty" bson:"street,omitempty"`
	Telephone       string   `json:"telephone,omitempty" bson:"telephone,omitempty"`
	Postcode        string   `json:"postcode,omitempty" bson:"postcode,omitempty"`
	City            string   `json:"city,omitempty" bson:"city,omitempty"`
	Firstname       string   `json:"firstname,omitempty" bson:"firstname,omitempty"`
	Lastname        string   `json:"lastname,omitempty" bson:"lastname,omitempty"`
	DefaultShipping bool     `json:"default_shipping,omitempty" bson:"default_shipping,omitempty"`
	Company         string   `json:"company,omitempty" bson:"company,omitempty"`
	VatID           string   `json:"vat_id,omitempty" bson:"vat_id,omitempty"`
	DefaultBilling  bool     `json:"default_billing,omitempty" bson:"default_billing,omitempty"`
	StreetLine0     string   `json:"street_line_0,omitempty"`
	StreetLine1     string   `json:"street_line_1,omitempty"`
}

type Region struct {
	RegionCode string `json:"region_code" bson:"region_code"`
	Region     string `json:"region" bson:"region"`
	RegionID   int    `json:"region_id" bson:"region_id"`
}

type OrderHistory struct {
	Items          []Item `json:"items" bson:"items"`
	SearchCriteria string `json:"search_criteria" bson:"search_criteria"`
	TotalCount     int    `json:"total_count" bson:"omitempty"`
}

type Item struct {
	SKU string `json:"sku,omitempty" bson:"sku"`
}

// collectionName Collection name
const collectionName = "users"

// Database operations
// MYSQL
func insertUserIntoMySQL(user User) int64 {
	passwordHash, err := hashPassword(user.Password)
	helpers.PanicErr(err)
	user.Password = passwordHash
	db, err := config.Conf.GetDb()
	helpers.PanicErr(err)
	result, err := db.Exec("INSERT INTO users("+
		"email, "+
		"password, "+
		"firstname, "+
		"lastname,"+
		"created_at, "+
		"updated_at)"+
		" VALUES(?, ?, ?, ?, ?, ?)",
		user.Customer.Email,
		user.Password,
		user.Customer.FirstName,
		user.Customer.LastName,
		time.Now().UTC(),
		time.Now().UTC())
	helpers.PanicErr(err)
	lastInsertId, err := result.LastInsertId()
	helpers.PanicErr(err)
	return lastInsertId
}

func (user CustomerData) UpdateUserByIdMySQL() {
	dataBase, err := config.Conf.GetDb()
	helpers.PanicErr(err)

	query, err := dataBase.Prepare("UPDATE users SET " +
		"email=?, " +
		"firstname=?, " +
		"lastname=?, " +
		"updated_at=? ," +
		"default_shipping=? " +
		"WHERE id=?")
	helpers.PanicErr(err)

	_, er := query.Exec(user.Email, user.FirstName, user.LastName, time.Now().UTC(), user.DefaultShipping, user.ID)
	helpers.PanicErr(er)
	fmt.Println(user.Email + " updated in mysql")
}

func getUserFromMySQLByEmail(email string) User {
	db, err := config.Conf.GetDb()
	helpers.PanicErr(err)
	var userFromDb User
	err = db.QueryRow("SELECT ID, Password, Group_id FROM users u WHERE email = ?", email).Scan(&userFromDb.ID, &userFromDb.Password, &userFromDb.GroupId)
	helpers.PanicErr(err)

	return userFromDb
}

func getUserIdFromMySQLByEmail(email string) string {
	db, err := config.Conf.GetDb()
	helpers.PanicErr(err)
	var id string
	err = db.QueryRow("SELECT ID FROM users u WHERE email = ?", email).Scan(&id)
	helpers.PanicErr(err)

	return id
}

func getUserFromMySQLById(id int64) CustomerData {
	db, err := config.Conf.GetDb()
	helpers.PanicErr(err)
	var userFromDb CustomerData
	err = db.QueryRow("SELECT "+
		"id, "+
		"email, "+
		"group_id, "+
		"default_shipping, "+
		"created_at, "+
		"updated_at, "+
		"created_in, "+
		"firstname, "+
		"lastname, "+
		"store_id, "+
		"website_id, "+
		"disable_auto_group_change FROM users u WHERE ID = ?", id).Scan(
		&userFromDb.ID,
		&userFromDb.Email,
		&userFromDb.GroupID,
		&userFromDb.DefaultShipping,
		&userFromDb.CreatedAt,
		&userFromDb.UpdatedAt,
		&userFromDb.CreatedIn,
		&userFromDb.FirstName,
		&userFromDb.LastName,
		&userFromDb.StoreID,
		&userFromDb.WebsiteID,
		&userFromDb.DisableAutoGroupChange)
	helpers.PanicErr(err)

	return userFromDb
}

func getAddressesFromMySQL(userId int64) []Address {
	db, err := config.Conf.GetDb()
	helpers.PanicErr(err)
	var addresses []Address
	rows, err := db.Query("SELECT "+
		"id, "+
		"customer_id, "+
		"region_id, "+
		"country_id, "+
		"telephone, "+
		"postcode, "+
		"city, "+
		"firstname, "+
		"lastname, "+
		"default_shipping, "+
		"street_line_0, "+
		"street_line_1 FROM addresses WHERE customer_id = ?", userId)
	for rows.Next() {
		var address Address
		if err := rows.Scan(&address.ID, &address.CustomerID, &address.RegionID,
			&address.CountryID, &address.Telephone, &address.Postcode, &address.City,
			&address.Firstname, &address.Lastname, &address.DefaultShipping,
			&address.StreetLine0, &address.StreetLine1); err != nil {
			helpers.PanicErr(err)
		}
		address.formatStreetArray()
		address.getRegion()
		addresses = append(addresses, address)
	}
	if len(addresses) > 0 {
		return addresses
	}
	return []Address{}
}

func (address *Address) insertOrUpdateAddressIntoMySQL(customerId string) {
	db, err := config.Conf.GetDb()
	helpers.PanicErr(err)
	res, err := db.Exec("INSERT INTO addresses("+
		"id, "+
		"customer_id, "+
		"region_id, "+
		"country_id, "+
		"telephone, "+
		"postcode, "+
		"city, "+
		"firstname, "+
		"lastname, "+
		"default_shipping, "+
		"street_line_0, "+
		"street_line_1)"+
		" VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?) ON DUPLICATE KEY UPDATE "+
		"region_id=VALUES(region_id), "+
		"country_id=VALUES(country_id), "+
		"telephone=VALUES(telephone), "+
		"postcode=VALUES(postcode), "+
		"city=VALUES(city), "+
		"firstname=VALUES(firstname), "+
		"lastname=VALUES(lastname), "+
		"default_shipping=VALUES(default_shipping),"+
		"street_line_0=VALUES(street_line_0),"+
		"street_line_1=VALUES(street_line_1)",
		address.ID,
		customerId,
		address.RegionID,
		address.CountryID,
		address.Telephone,
		address.Postcode,
		address.City,
		address.Firstname,
		address.Lastname,
		address.DefaultShipping,
		address.Street[0],
		address.Street[1])
	helpers.PanicErr(err)
	addressId, err := res.LastInsertId()
	address.ID = addressId
}

func (address *Address) getRegion() {
	db, err := config.Conf.GetDb()
	helpers.PanicErr(err)
	var region Region
	err = db.QueryRow("SELECT region_id, region_code, region "+
		"FROM region WHERE region_id=?", address.RegionID).Scan(&region.RegionID, &region.RegionCode, &region.Region)
	helpers.PanicErr(err)
	address.Region = region
}

func (address *Address) formatStreetArray() {
	address.Street = []string{address.StreetLine0, address.StreetLine1}
	address.StreetLine0 = ""
	address.StreetLine1 = ""
}

func updateAddressInMySQL(address Address) {
	db, err := config.Conf.GetDb()
	helpers.PanicErr(err)
	_, err = db.Exec("UPDATE addresses SET "+
		"region_id=?, "+
		"country_id=?, "+
		"telephone=?, "+
		"postcode=?, "+
		"city=?, "+
		"firstname=?, "+
		"lastname=?, "+
		"default_shipping=? "+
		"WHERE id=?",
		address.RegionID,
		address.CountryID,
		address.Telephone,
		address.Postcode,
		address.City,
		address.Firstname,
		address.Lastname,
		address.DefaultShipping)
	helpers.PanicErr(err)
}

func getGroupIdFromMySQLById(id int) int {
	db, err := config.Conf.GetDb()
	helpers.PanicErr(err)
	var groupId int
	err = db.QueryRow("SELECT group_id FROM users u WHERE id = ?", id).Scan(&groupId)
	helpers.PanicErr(err)

	return groupId
}

// MongoDB
func insertUserIntoMongo(userInfo CustomerData) {
	db := config.Conf.GetMongoDb()
	_, err := db.Collection(collectionName).InsertOne(context.Background(),
		bson.NewDocument(
			bson.EC.String("type", "User info"),
			bson.EC.Time("created_at", userInfo.CreatedAt),
			bson.EC.String("created_in", userInfo.CreatedIn),
			bson.EC.Int32("disable_auto_group_change", userInfo.DisableAutoGroupChange),
			bson.EC.String("email", userInfo.Email),
			bson.EC.String("firstname", userInfo.FirstName),
			bson.EC.String("lastname", userInfo.LastName),
			bson.EC.Int32("group_id", userInfo.GroupID),
			bson.EC.String("id", userInfo.ID),
			bson.EC.Int32("store_id", userInfo.StoreID),
			bson.EC.Time("updated_at", userInfo.UpdatedAt),
			bson.EC.Int32("website_id", userInfo.WebsiteID),
			bson.EC.Interface("address", userInfo.Addresses)))
	helpers.PanicErr(err)
}

func getUserFromMongo(id string) CustomerData {
	db := config.Conf.GetMongoDb()
	var userInfo CustomerData
	err := db.Collection(collectionName).FindOne(context.Background(), bson.NewDocument(
		bson.EC.Interface("id", id),
		bson.EC.String("type", "User info"))).Decode(&userInfo)
	helpers.PanicErr(err)
	if len(userInfo.Addresses) == 0 {
		userInfo.Addresses = []Address{}
	}
	fmt.Println(userInfo.Addresses)
	return userInfo
}

func getUserOrderHistoryFromMongo(id string) OrderHistory {
	db := config.Conf.GetMongoDb()

	cur, err := db.Collection(collectionName).Find(context.Background(), bson.NewDocument(
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
