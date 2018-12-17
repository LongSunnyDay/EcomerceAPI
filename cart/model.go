package cart

import (
	"context"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/objectid"
	"go-api-ws/config"
	"go-api-ws/helpers"
	"time"
)

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
	QTY           float64 `json:"qty,omitempty" bson:"qty"`
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

// collectionName Collection name
const collectionName = "cart"

func getGuestCartIDFromMongo(guestCartID string) (cartID string, err error) {
	db := config.Conf.GetMongoDb()

	bsonData := bson.NewDocument()
	err = db.Collection(collectionName).FindOne(nil, bson.NewDocument(
		bson.EC.Interface("_id", guestCartID))).Decode(&bsonData)
	if err != nil {
		return "", err
	}
	cartID = bsonData.LookupElement("_id").Value().ObjectID().Hex()
	return cartID, nil
}

func getUserCartIDFromMongo(userID string) (cartID string, err error) {
	db := config.Conf.GetMongoDb()

	bsonData := bson.NewDocument()
	err = db.Collection(collectionName).FindOne(nil, bson.NewDocument(
		bson.EC.String("id", userID))).Decode(&bsonData)
	if err != nil {
		return "", err
	}
	cartID = bsonData.LookupElement("id").Value().StringValue()
	return cartID, nil
}

func CreateCartInMongoDB(userID string) (cartID string) {
	db := config.Conf.GetMongoDb()

	if userID == "" {
		cart := Cart{
			Items:     []Item{},
			CreatedAt: time.Now()}
		bsonCart, err := helpers.StructToBson(cart)
		helpers.PanicErr(err)
		result, err := db.Collection(collectionName).InsertOne(context.Background(), bsonCart)
		cartID := result.InsertedID.(objectid.ObjectID).Hex()
		return cartID
	} else {
		cart := Cart{
			Items: []Item{},
			ID:    userID}
		bsonCart, err := helpers.StructToBson(cart)
		helpers.PanicErr(err)
		_, err = db.Collection(collectionName).InsertOne(context.Background(), bsonCart)
		helpers.PanicErr(err)
		cartID, err = getUserCartIDFromMongo(userID)
		helpers.PanicErr(err)
		return cartID
	}
}

func getGuestCartFromMongoByID(guestCartID string) []Item {
	db := config.Conf.GetMongoDb()

	cart := Cart{Items: []Item{}}
	objectIDFromUserID, err := objectid.FromHex(guestCartID)
	helpers.PanicErr(err)
	err = db.Collection(collectionName).FindOne(context.Background(), bson.NewDocument(
		bson.EC.ObjectID("_id", objectIDFromUserID))).Decode(&cart)
	if err != nil {
		carID := CreateCartInMongoDB(guestCartID)
		getGuestCartFromMongoByID(carID)
	}
	return cart.Items
}

func GetUserCartFromMongoByID(userID string) []Item {
	db := config.Conf.GetMongoDb()

	cart := Cart{Items: []Item{}}
	err := db.Collection(collectionName).FindOne(context.Background(), bson.NewDocument(
		bson.EC.String("id", userID))).Decode(&cart)
	if err != nil {
		carID := CreateCartInMongoDB(userID)
		getGuestCartFromMongoByID(carID)
	}
	return cart.Items
}

func updateUserCartInMongo(cartID string, item Item) {
	bsonItem, err := helpers.StructToBson(item)
	helpers.PanicErr(err)

	db := config.Conf.GetMongoDb()

	db.Collection(collectionName).UpdateOne(nil,
		bson.NewDocument(
			bson.EC.String("id", cartID)),
		bson.NewDocument(
			bson.EC.SubDocumentFromElements("$pull",
				bson.EC.SubDocument("items",
					bson.NewDocument(
						bson.EC.String("sku", item.SKU))))))

	_, err = db.Collection(collectionName).UpdateOne(nil,
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

	db := config.Conf.GetMongoDb()

	db.Collection(collectionName).UpdateOne(nil,
		bson.NewDocument(
			bson.EC.ObjectID("_id", bsonCartID)),
		bson.NewDocument(
			bson.EC.SubDocumentFromElements("$pull",
				bson.EC.SubDocument("items",
					bson.NewDocument(
						bson.EC.String("sku", item.SKU))))))

	_, err = db.Collection(collectionName).UpdateOne(nil,
		bson.NewDocument(
			bson.EC.ObjectID("_id", bsonCartID)),
		bson.NewDocument(
			bson.EC.SubDocumentFromElements("$push",
				bson.EC.SubDocument("items", bsonItem))))
	helpers.PanicErr(err)
}

func deleteItemFromCartInMongo(carId string, item CartItem) {
	db := config.Conf.GetMongoDb()

	_, err := db.Collection(collectionName).UpdateOne(nil,
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

	db := config.Conf.GetMongoDb()

	_, err = db.Collection(collectionName).UpdateOne(nil,
		bson.NewDocument(
			bson.EC.ObjectID("_id", bsonCartID)),
		bson.NewDocument(
			bson.EC.SubDocumentFromElements("$pull",
				bson.EC.SubDocument("items",
					bson.NewDocument(
						bson.EC.String("sku", item.Item.SKU))))))
	helpers.PanicErr(err)
}
