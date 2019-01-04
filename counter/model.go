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

type QuoteCounter struct {
	ID int64 `json:"id"`
	Name string `json:"name"`
	Value int64 `json:"value"`
}

// collectionName Collection name
const collectionName = "counters"
const quoteCounter = "quote"

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

func GetAndIncreaseQuoteCounterInMySQL() (value int64) {
	db, err := config.Conf.GetDb()
	helpers.PanicErr(err)

	//defer db.Close()

	tx, err := db.Begin()
	helpers.PanicErr(err)

	err = tx.QueryRow("SELECT value FROM counters WHERE name = ?", quoteCounter).Scan(&value)
	if err != nil {
		tx.Rollback()
		helpers.PanicErr(err)
	}

	_, err = tx.Exec("UPDATE counters SET value = value + 1 WHERE name = ?", quoteCounter)
	if err != nil {
		tx.Rollback()
		helpers.PanicErr(err)
	}

	err = tx.Commit()
	helpers.PanicErr(err)

	return
}
