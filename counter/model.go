package counter

import (
	"github.com/mongodb/mongo-go-driver/bson"
	"go-api-ws/config"
	"go-api-ws/helpers"
)

type ItemCounter struct {
	ID  string `json:"_id,omitempty" bson:"_id"`
	Value int    `json:"value,omitempty" bson:"value"`
}

// CONNECTIONSTRING DB connection string
//const CONNECTIONSTRING = "mongodb://localhost:27017"

// DBNAME Database name
//const DBNAME = "go-api-ws"

// COLLNAME Collection name
const COLLNAME = "counters"

//var db *mongo.Database

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

func GetAndIncreaseItemIdCounterInMongo() int {
	var itemCounter ItemCounter

	db := config.Conf.GetMongoDb()

	err := db.Collection(COLLNAME).FindOneAndUpdate(nil,
		bson.NewDocument(
			bson.EC.String("_id", "itemid")),
			bson.NewDocument(
				bson.EC.SubDocumentFromElements("$inc",
					bson.EC.Interface("value", 1)))).Decode(&itemCounter)
	helpers.PanicErr(err)
	return itemCounter.Value
}