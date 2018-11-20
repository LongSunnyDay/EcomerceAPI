package user

import (
	"github.com/mongodb/mongo-go-driver/mongo"
	"go-api-ws/helpers"
	"context"
	"github.com/mongodb/mongo-go-driver/bson"
	"fmt"
)


// Connect establish a connection to database
func init() {
	client, err := mongo.NewClient(CONNECTIONSTRING)
	helpers.PanicErr(err)

	err = client.Connect(context.Background())
	helpers.PanicErr(err)

	// Collection types can be used to access the database
	db = client.Database(DBNAME)
}

func roleByGroupId(groupId int) (string) {
	if groupId < 1 {
		return adminRole
	}
	return userRole
}

func UpdateUserById(user UpdateUser)interface{}{
	var updatedUser UpdateUser
	userId := user.ID
	fmt.Println(userId)
	bsonUser, err := helpers.StructToBson(user)
	helpers.PanicErr(err)

	filter := bson.NewDocument(bson.EC.Interface("id", userId))
	doc := bson.NewDocument(bson.EC.SubDocument("$set", bsonUser))


	res := db.Collection(COLLNAME).FindOneAndUpdate(context.Background(), filter, doc)

	res.Decode(&updatedUser)


	fmt.Println(updatedUser)
	helpers.PanicErr(err)

	return  updatedUser
}
