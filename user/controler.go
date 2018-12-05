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

func UpdateUserByIdMongo(user CustomerData) {
	fmt.Println(user)
	bsonUser, err := helpers.StructToBson(user)
	helpers.PanicErr(err)

	filter := bson.NewDocument(bson.EC.Interface("id", user.ID))
	doc := bson.NewDocument(bson.EC.SubDocument("$set", bsonUser))

	db := config.Conf.GetMongoDb()

	db.Collection(collectionName).FindOneAndUpdate(context.Background(), filter, doc)
	//helpers.PanicErr(err)

}


