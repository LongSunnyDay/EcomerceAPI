package user

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/mongodb/mongo-go-driver/bson"
	"go-api-ws/config"
	"go-api-ws/helpers"
)

func roleByGroupId(groupId int) string {
	if groupId < 1 {
		return adminRole
	}
	return userRole
}

func UpdateUserByIdMongo(user UpdateUser) interface{} {
	var updatedUser UpdateUser

	bsonUser, err := helpers.StructToBson(user)
	helpers.PanicErr(err)

	filter := bson.NewDocument(bson.EC.Interface("id", user.ID))
	doc := bson.NewDocument(bson.EC.SubDocument("$set", bsonUser))

	db := config.Conf.GetMongoDb()

	update := db.Collection(collectionName).FindOneAndUpdate(context.Background(), filter, doc)
	//helpers.PanicErr(err)
	fmt.Println(update)

	return updatedUser
}

func UpdateUserByIdMySQL(user UpdateUser){
	dataBase, err := config.Conf.GetDb()
	helpers.PanicErr(err)

	query, err := dataBase.Prepare("Update users set email=? where id=?")
	helpers.PanicErr(err)

	_, er := query.Exec(user.Email, user.ID)
	helpers.PanicErr(er)
	fmt.Println(user.Email + " updated in mysql")
}
