package user

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/mongodb/mongo-go-driver/bson"
	"go-api-ws/config"
	"go-api-ws/helpers"
)


// Connect establish a connection to database
//func init() {
//	client, err := mongo.NewClient(CONNECTIONSTRING)
//	helpers.PanicErr(err)
//
//	err = client.Connect(context.Background())
//	helpers.PanicErr(err)
//
//	// Collection types can be used to access the database
//	db = client.Database(DBNAME)
//}

func roleByGroupId(groupId int) (string) {
	if groupId < 1 {
		return adminRole
	}
	return userRole
}

func UpdateUserByIdMongo(user UpdateUser)interface{}{
	var updatedUser UpdateUser

	bsonUser, err := helpers.StructToBson(user)
	helpers.PanicErr(err)

	filter := bson.NewDocument(bson.EC.Interface("id", user.ID))
	doc := bson.NewDocument(bson.EC.SubDocument("$set", bsonUser))

	db := config.Conf.GetMongoDb()

	lopas := db.Collection(COLLNAME).FindOneAndUpdate(context.Background(), filter, doc)
	//helpers.PanicErr(err)
	fmt.Println(lopas)

	return  updatedUser
}

func UpdateUserByIdMySQL(user UpdateUser){
	dataBase, err := config.Conf.GetDb()
	helpers.PanicErr(err)

	query, err := dataBase.Prepare("Update users set email=? where id=?")
	helpers.PanicErr(err)

	_, er := query.Exec(user.Email, user.ID)
	helpers.PanicErr(er)
	fmt.Print(user.Email + " updated int mysql")
}
