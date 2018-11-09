package cart

import (
	"context"
	"fmt"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"go-api-ws/helpers"
	"net/http"
)

type Response struct {
	Acknowledged bool          `json:"acknowledged,omitempty"`
	Code         int           `json:"code,omitempty"`
	CreatedAt    string        `json:"created_at,omitempty"`
	Payload      *http.Request `json:"payload,omitempty"`
	Result       interface{}   `json:"result,omitempty"`
	ResultCode   int           `json:"result_code,omitempty"`
	TaskID       string        `json:"task_id,omitempty"`
	Transmited   bool          `json:"transmited,omitempty"`
	TransmitedAt string        `json:"transmited_at,omitempty"`
	UpdatedAt    string        `json:"updated_at,omitempty"`
	Url          string        `json:"url,omitempty"`
	Meta         interface{}   `json:"meta,omitempty"`
}

type cartData struct {
	ObjectId string `json:"_id,omitempty" bson:"_id"`
	ID       string `json:"id,omitempty" bson:"id"`
}

type Cart struct {
	Items []Item `json:"items" bson:"items"`
}

type Item struct {
	SKU           string `json:"sku,omitempty" bson:"sku"`
	QTY           int    `json:"qty,omitempty" bson:"qty"`
	Price         int    `json:"price,omitempty" bson:"price"`
	ProductType   string `json:"product_type,omitempty" bson:"product_type"`
	Name          string `json:"name,omitempty" bson:"name"`
	ItemID        int    `json:"item_id,omitempty" bson:"item_id"`
	QuoteId       string `json:"quote_id,omitempty" bson:"quote_id"`
	ProductOption struct {
		ExtensionAttributes struct {
			ConfigurableItemOptions []Options `json:"configurable_item_options,omitempty" bson:"configurable_item_options"`
		} `json:"extension_attributes,omitempty" bson:"extension_attributes"`
	} `json:"product_option,omitempty" bson:"product_options"`
}

type Options struct {
	OptionsID   string `json:"option_id,omitempty" bson:"option_id"`
	OptionValue int    `json:"option_value,omitempty" bson:"option_value"`
}

// CONNECTIONSTRING DB connection string
const CONNECTIONSTRING = "mongodb://localhost:32768"

// DBNAME Database name
const DBNAME = "go-api-ws"

// COLLNAME Collection name
const COLLNAME = "cart"

var db *mongo.Database

// Connect establish a connection to database
func init() {
	client, err := mongo.NewClient(CONNECTIONSTRING)
	helpers.PanicErr(err)

	err = client.Connect(context.Background())
	helpers.PanicErr(err)

	// Collection types can be used to access the database
	db = client.Database(DBNAME)
}

func getCartIDFromMongo(userId string, userType string) (string) {
	cur, err := db.Collection(COLLNAME).Find(context.Background(), bson.NewDocument(
		bson.EC.Interface("_id", userId)))
	helpers.PanicErr(err)
	bsonData := bson.NewDocument()
	for cur.Next(context.Background()) {
		err := cur.Decode(&bsonData)
		helpers.PanicErr(err)
	}
	cur.Close(context.Background())
	if userType == "" {
		cartID := bsonData.LookupElement("_id").Value().ObjectID().Hex()
		return cartID
	}
	cartID := bsonData.LookupElement("_id").Value().StringValue()
	return cartID
}

func CreateUserCartInMongo(id string) {
	_, err := db.Collection(COLLNAME).InsertOne(context.Background(),
		bson.NewDocument(
			bson.EC.Interface("_id", id)))
	helpers.PanicErr(err)
}

func createGuestCartInMongo(id string) (string) {
	_, err := db.Collection(COLLNAME).InsertOne(context.Background(), bson.NewDocument(
		bson.EC.String("id", id)))
	helpers.PanicErr(err)

	cur, err := db.Collection(COLLNAME).Find(context.Background(), bson.NewDocument(
		bson.EC.String("id", id)))
	helpers.PanicErr(err)
	bsonData := bson.NewDocument()
	for cur.Next(context.Background()) {
		err := cur.Decode(&bsonData)
		helpers.PanicErr(err)
	}
	cur.Close(context.Background())
	idFromMongo := bsonData.LookupElement("_id").Value().ObjectID().Hex()
	return idFromMongo
}

func getCartFromMongoByID(userId string) {
	var cart Cart
	err := db.Collection(COLLNAME).FindOne(context.Background(), bson.NewDocument(
		bson.EC.String("_id", userId))).Decode(&cart)
	helpers.PanicErr(err)
	fmt.Println(cart)
}
