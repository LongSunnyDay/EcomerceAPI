package cart

import (
	"context"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/objectid"
	"github.com/mongodb/mongo-go-driver/mongo"
	"go-api-ws/helpers"
	"net/http"
	"time"
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
	ID        string    `json:"id,omitempty" bson:"id,omitempty"`
	Items     []Item    `json:"items" bson:"items"`
	CreatedAt time.Time `json:"created_at,omitempty" bson:"createdAt,omitempty"`
}

type CartItem struct {
	Item Item `json:"cartItem,omitempty" bson:"cartItem"`
}

type Item struct {
	SKU           string  `json:"sku,omitempty" bson:"sku"`
	QTY           int32   `json:"qty,omitempty" bson:"qty"`
	Price         float64 `json:"price,omitempty" bson:"price"`
	ProductType   string  `json:"product_type,omitempty" bson:"product_type"`
	Name          string  `json:"name,omitempty" bson:"name"`
	ItemID        int     `json:"item_id,omitempty" bson:"item_id,omitempty"`
	QuoteId       string  `json:"quoteId,omitempty" bson:"quoteId"`
	ProductOption struct {
		ExtensionAttributes struct {
			ConfigurableItemOptions []Options `json:"configurable_item_options,omitempty" bson:"configurable_item_options"`
		} `json:"extension_attributes,omitempty" bson:"extension_attributes"`
	} `json:"product_option,omitempty" bson:"product_options"`
}

type Options struct {
	OptionsID   string `json:"option_id,omitempty" bson:"option_id"`
	OptionValue string `json:"option_value,omitempty" bson:"option_value"`
}

type PaymentMethod struct {
	Type           string `bson:"type"`
	Code           string `json:"code,omitempty" bson:"code"`
	Title          string `json:"title,omitempty" bson:"title"`
	IsServerMethod bool   `json:"is_server_method,omitempty" bson:"is_server_method"`
}

// CONNECTIONSTRING DB connection string
const CONNECTIONSTRING = "mongodb://localhost:27017"

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

func getGuestCartIDFromMongo(guestCartID string) (cartID string, err error) {
	bsonData := bson.NewDocument()
	err = db.Collection(COLLNAME).FindOne(nil, bson.NewDocument(
		bson.EC.Interface("_id", guestCartID))).Decode(&bsonData)
	if err != nil {
		return "", err
	}
	cartID = bsonData.LookupElement("_id").Value().ObjectID().Hex()
	return cartID, nil
}

func getUserCartIDFromMongo(userID string) (cartID string, err error) {
	bsonData := bson.NewDocument()
	err = db.Collection(COLLNAME).FindOne(nil, bson.NewDocument(
		bson.EC.String("id", userID))).Decode(&bsonData)
	if err != nil {
		return "", err
	}
	cartID = bsonData.LookupElement("id").Value().StringValue()
	return cartID, nil
}

func CreateCartInMongoDB(userID string) (cartID string) {
	if userID == "" {
		cart := Cart{
			Items:     []Item{},
			CreatedAt: time.Now()}
		bsonCart, err := helpers.StructToBson(cart)
		helpers.PanicErr(err)
		result, err := db.Collection(COLLNAME).InsertOne(context.Background(), bsonCart)
		cartID := result.InsertedID.(objectid.ObjectID).Hex()
		return cartID
	} else {
		cart := Cart{
			Items: []Item{},
			ID:    userID}
		bsonCart, err := helpers.StructToBson(cart)
		helpers.PanicErr(err)
		_, err = db.Collection(COLLNAME).InsertOne(context.Background(), bsonCart)
		helpers.PanicErr(err)
		cartID, err = getUserCartIDFromMongo(userID)
		helpers.PanicErr(err)
		return cartID
	}
}

func getGuestCartFromMongoByID(guestCartID string) []Item {
	cart := Cart{Items: []Item{}}
	objectIDFromUserID, err := objectid.FromHex(guestCartID)
	helpers.PanicErr(err)
	err = db.Collection(COLLNAME).FindOne(context.Background(), bson.NewDocument(
		bson.EC.ObjectID("_id", objectIDFromUserID))).Decode(&cart)
	if err != nil {
		carID := CreateCartInMongoDB(guestCartID)
		getGuestCartFromMongoByID(carID)
	}
	return cart.Items
}

func getUserCartFromMongoByID(userID string) []Item {
	cart := Cart{Items: []Item{}}
	err := db.Collection(COLLNAME).FindOne(context.Background(), bson.NewDocument(
		bson.EC.String("id", userID))).Decode(&cart)
	if err != nil {
		carID := CreateCartInMongoDB(userID)
		getGuestCartFromMongoByID(carID)
	}
	return cart.Items
}

func insertPaymentMethodsToMongo(methods []interface{}) {
	_, err := db.Collection(COLLNAME).InsertMany(nil, methods)
	helpers.PanicErr(err)
}

func getPaymentMethodsFromMongo() []PaymentMethod {
	var paymentMethod PaymentMethod
	var paymentMethods []PaymentMethod

	cur, err := db.Collection(COLLNAME).Find(nil, bson.NewDocument(
		bson.EC.String("type", "Payment method")))
	helpers.PanicErr(err)
	for cur.Next(context.Background()) {
		err := cur.Decode(&paymentMethod)
		helpers.PanicErr(err)
		paymentMethods = append(paymentMethods, paymentMethod)
	}
	cur.Close(context.Background())
	return paymentMethods
}

func updateUserCartInMongo(cartID string, item Item) {
	bsonItem, err := helpers.StructToBson(item)
	helpers.PanicErr(err)

	db.Collection(COLLNAME).UpdateOne(nil,
		bson.NewDocument(
			bson.EC.String("id", cartID)),
		bson.NewDocument(
			bson.EC.SubDocumentFromElements("$pull",
				bson.EC.SubDocument("items",
					bson.NewDocument(
						bson.EC.String("sku", item.SKU))))))

	_, err = db.Collection(COLLNAME).UpdateOne(nil,
		bson.NewDocument(
			bson.EC.String("id", cartID)),
		bson.NewDocument(
			bson.EC.SubDocumentFromElements("$push",
				bson.EC.SubDocument("items", bsonItem))))
	helpers.PanicErr(err)
}

func updateGuestCartInMongo(cartID string, item Item) {
	bsonItem, err := helpers.StructToBson(item)
	helpers.PanicErr(err)

	bsonCartID, err := objectid.FromHex(cartID)
	helpers.PanicErr(err)

	db.Collection(COLLNAME).UpdateOne(nil,
		bson.NewDocument(
			bson.EC.ObjectID("_id", bsonCartID)),
		bson.NewDocument(
			bson.EC.SubDocumentFromElements("$pull",
				bson.EC.SubDocument("items",
					bson.NewDocument(
						bson.EC.String("sku", item.SKU))))))

	_, err = db.Collection(COLLNAME).UpdateOne(nil,
		bson.NewDocument(
			bson.EC.ObjectID("_id", bsonCartID)),
		bson.NewDocument(
			bson.EC.SubDocumentFromElements("$push",
				bson.EC.SubDocument("items", bsonItem))))
	helpers.PanicErr(err)
}

func deleteItemFromCartInMongo(carId string, item CartItem) {
	_, err := db.Collection(COLLNAME).UpdateOne(nil,
		bson.NewDocument(
			bson.EC.String("id", carId)),
		bson.NewDocument(
			bson.EC.SubDocumentFromElements("$pull",
				bson.EC.SubDocument("items",
					bson.NewDocument(
						bson.EC.String("sku", item.Item.SKU))))))
	helpers.PanicErr(err)
}

func deleteItemFromGuestCartInMongo(cartID string, item CartItem) {
	bsonCartID, err := objectid.FromHex(cartID)
	helpers.PanicErr(err)

	_, err = db.Collection(COLLNAME).UpdateOne(nil,
		bson.NewDocument(
			bson.EC.ObjectID("_id", bsonCartID)),
		bson.NewDocument(
			bson.EC.SubDocumentFromElements("$pull",
				bson.EC.SubDocument("items",
					bson.NewDocument(
						bson.EC.String("sku", item.Item.SKU))))))
	helpers.PanicErr(err)
}
