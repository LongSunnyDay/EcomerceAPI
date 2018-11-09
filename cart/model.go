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

type Cart struct {
	Items []item `json:"items" bson:"items"`
}

type item struct {
	sku string `json:"sku,omitempty" bson:"sku"`
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

func getUserCartFromMongo(userId float64) (Cart) {
	cur, err := db.Collection(COLLNAME).Find(context.Background(), bson.NewDocument(
		bson.EC.Interface("userId", userId),
		bson.EC.String("type", "User cart")))
	helpers.PanicErr(err)
	var cart Cart
	for cur.Next(context.Background()) {
		err := cur.Decode(&cart)
		helpers.PanicErr(err)
	}
	fmt.Println(cart)
	cur.Close(context.Background())
	return cart
}

func CreateUserCartInMongo(id interface{})  {
	_, err := db.Collection(COLLNAME).InsertOne(context.Background(),
		bson.NewDocument(
			bson.EC.Interface("id", id)))
	helpers.PanicErr(err)
}

func updateUserCartInMongo(userId float64)  {
	//cur, err := db.Collection(.)
}