package user

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/mongodb/mongo-go-driver/bson"
	"go-api-ws/addresses"
	"go-api-ws/config"
	"go-api-ws/helpers"
	"time"
)

// MySQL operations

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

func GetUserFromMySQLById(id int64) CustomerData {
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

func GetGroupIdFromMySQLById(id int) int {
	db, err := config.Conf.GetDb()
	helpers.PanicErr(err)
	var groupId int
	err = db.QueryRow("SELECT group_id FROM users u WHERE id = ?", id).Scan(&groupId)
	helpers.PanicErr(err)

	return groupId
}

// MongoDB operations

func UpdateUserByIdMongo(user CustomerData) {
	fmt.Println(user)
	bsonUser, err := helpers.StructToBson(user)
	helpers.PanicErr(err)

	filter := bson.NewDocument(bson.EC.Interface("id", user.ID))
	doc := bson.NewDocument(bson.EC.SubDocument("$set", bsonUser))

	db := config.Conf.GetMongoDb()

	db.Collection(collectionName).FindOneAndUpdate(context.Background(), filter, doc)

}

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
			bson.EC.Int64("group_id", userInfo.GroupID),
			bson.EC.Int64("id", userInfo.ID),
			bson.EC.Int32("store_id", userInfo.StoreID),
			bson.EC.Time("updated_at", userInfo.UpdatedAt),
			bson.EC.Int32("website_id", userInfo.WebsiteID),
			bson.EC.Interface("addresses", userInfo.Addresses)))
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
		userInfo.Addresses = []addresses.Address{}
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
	err = cur.Close(context.Background())
	helpers.PanicErr(err)
	return orderHistory
}

// Other data operations

func roleByGroupId(groupId int) string {
	if groupId < 1 {
		return adminRole
	}
	return userRole
}
