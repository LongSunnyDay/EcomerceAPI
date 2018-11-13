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

type CartItem struct {
	Item Item `json:"cartItem,omitempty" bson:"cartItem"`
}

type Item struct {
	SKU           string  `json:"sku,omitempty" bson:"sku"`
	QTY           int32   `json:"qty,omitempty" bson:"qty"`
	Price         float64 `json:"price,omitempty" bson:"price"`
	ProductType   string  `json:"product_type,omitempty" bson:"product_type"`
	Name          string  `json:"name,omitempty" bson:"name"`
	ItemID        int64   `json:"item_id,omitempty" bson:"item_id"`
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

func getCartIDFromMongo(userId string, userType string) string {
	bsonData := bson.NewDocument()
	err := db.Collection(COLLNAME).FindOne(nil, bson.NewDocument(
		bson.EC.Interface("_id", userId))).Decode(&bsonData)
	helpers.PanicErr(err)
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

func createGuestCartInMongo(id string) string {
	_, err := db.Collection(COLLNAME).InsertOne(context.Background(), bson.NewDocument(
		bson.EC.String("id", id)))
	helpers.PanicErr(err)

	bsonData := bson.NewDocument()
	err = db.Collection(COLLNAME).FindOne(nil, bson.NewDocument(
		bson.EC.String("id", id))).Decode(&bsonData)
	helpers.PanicErr(err)
	idFromMongo := bsonData.LookupElement("_id").Value().ObjectID().Hex()
	return idFromMongo
}

func getCartFromMongoByID(userId string) Cart {
	cart := Cart{Items: []Item{}}
	err := db.Collection(COLLNAME).FindOne(context.Background(), bson.NewDocument(
		bson.EC.Interface("_id", userId))).Decode(&cart)
	if err != nil {
		createGuestCartInMongo(userId)
	}
	return cart
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

func updateUserCartInMongo(cartId string, item CartItem) {
	data := bson.NewDocument()
	result := db.Collection(COLLNAME).FindOneAndUpdate(nil,
		bson.NewDocument(
			bson.EC.String("_id", cartId)),
		bson.NewDocument(
			bson.EC.SubDocumentFromElements("$addToSet",
				bson.EC.SubDocument("items",
					bson.NewDocument(
						bson.EC.Int64("item_id", item.Item.ItemID),
						bson.EC.String("sku", item.Item.SKU),
						bson.EC.Int32("qty", item.Item.QTY),
						bson.EC.String("name", item.Item.Name),
						bson.EC.Double("price", item.Item.Price),
						bson.EC.String("product_type", item.Item.ProductType),
						bson.EC.String("quoteId", item.Item.QuoteId),
						bson.EC.SubDocument("product_option",
							bson.NewDocument(bson.EC.SubDocument("extension_attributes",
								bson.NewDocument(bson.EC.Array("configurable_item_options", bson.NewArray(
									bson.VC.DocumentFromElements(
										bson.EC.String("option_id", item.Item.ProductOption.ExtensionAttributes.ConfigurableItemOptions[0].OptionsID),
										bson.EC.String("option_value", item.Item.ProductOption.ExtensionAttributes.ConfigurableItemOptions[0].OptionValue)),
									bson.VC.DocumentFromElements(
										bson.EC.String("option_id", item.Item.ProductOption.ExtensionAttributes.ConfigurableItemOptions[1].OptionsID),
										bson.EC.String("option_value", item.Item.ProductOption.ExtensionAttributes.ConfigurableItemOptions[1].OptionValue))))))))))))).Decode(&data)
	fmt.Println(result)
	fmt.Println("updateUserCartInMongo - ", item)
	fmt.Println("bsonData - ", data)
}

func deleteItemFromUserCartInMongo(carId string, item CartItem) {
	data := bson.NewDocument()
	result, err := db.Collection(COLLNAME).UpdateOne(nil,
		bson.NewDocument(
			bson.EC.String("_id", carId)),
		bson.NewDocument(
			bson.EC.SubDocumentFromElements("$pull",
				bson.EC.SubDocument("items",
					bson.NewDocument(
						bson.EC.String("sku", item.Item.SKU),
						bson.EC.String("quoteId", item.Item.QuoteId))))))
	helpers.PanicErr(err)
	fmt.Println("Result", result)
	fmt.Println("Data", data)
}
