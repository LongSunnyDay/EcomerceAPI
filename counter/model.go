package counter

import (
	"github.com/mongodb/mongo-go-driver/bson"
	"go-api-ws/config"
	"go-api-ws/helpers"
)

type ItemCounter struct {
	ID    string `json:"_id,omitempty" bson:"_id"`
	Value int    `json:"value,omitempty" bson:"value"`
}

// collectionName Collection name
const collectionName = "counters"

func GetAndIncreaseItemIdCounterInMongo() int {
	var itemCounter ItemCounter

	db := config.Conf.GetMongoDb()

	err := db.Collection(collectionName).FindOneAndUpdate(nil,
		bson.NewDocument(
			bson.EC.String("_id", "itemid")),
		bson.NewDocument(
			bson.EC.SubDocumentFromElements("$inc",
				bson.EC.Interface("value", 1)))).Decode(&itemCounter)
	helpers.PanicErr(err)
	return itemCounter.Value
}
