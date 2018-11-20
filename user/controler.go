package user

import (
	"context"
		"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"go-api-ws/helpers"
	c "../config"
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

func UpdateUserByIdMongo(user UpdateUser)interface{}{
	var updatedUser UpdateUser

	bsonUser, err := helpers.StructToBson(user)
	helpers.PanicErr(err)

	filter := bson.NewDocument(bson.EC.Interface("id", user.ID))
	doc := bson.NewDocument(bson.EC.SubDocument("$set", bsonUser))

	err = db.Collection(COLLNAME).FindOneAndUpdate(context.Background(), filter, doc).Decode(&updatedUser)
	helpers.PanicErr(err)

	return  updatedUser
}

func UpdateUserByIdMySQL(user UpdateUser){
	dataBase, err := c.Conf.GetDb()
	helpers.PanicErr(err)

	err = dataBase.QueryRow("SELECT * FROM users u Where id=?", user.ID).Scan(&user.Email)
	helpers.PanicErr(err)

	//usr := user
	//for rows.Next() {
	//	rows.Scan(&usr.Email)
	//}

}
