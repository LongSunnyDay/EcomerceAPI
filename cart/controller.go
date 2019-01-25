package cart

import (
	"context"
	"fmt"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/objectid"
	"go-api-ws/config"
	"go-api-ws/counter"
	"go-api-ws/helpers"
	"strconv"
	"time"
)

// MySQL operations

func CreateCartInMongoDB(userID string) string {
	db := config.Conf.GetMongoDb()
	quoteId := counter.GetAndIncreaseQuoteCounterInMySQL()

	quoteIdString := strconv.Itoa(int(quoteId))

	if userID == "" {
		cart := Cart{
			Items:     []Item{},
			CreatedAt: time.Now(),
			QuoteId:   quoteId,
			CartId:    quoteIdString,
			Status:    "Active"}

		bsonCart, err := helpers.StructToBson(cart)
		helpers.PanicErr(err)

		_, err = db.Collection(collectionName).InsertOne(context.Background(), bsonCart)
		helpers.PanicErr(err)

		return quoteIdString
	} else {
		cart := Cart{
			Items:   []Item{},
			UserId:  userID,
			CartId:  quoteIdString,
			QuoteId: quoteId,
			Status:  "Active"}

		bsonCart, err := helpers.StructToBson(cart)
		helpers.PanicErr(err)

		_, err = db.Collection(collectionName).InsertOne(context.Background(), bsonCart)
		helpers.PanicErr(err)
		return quoteIdString
	}
}

// MongoDB operations

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

func getGuestCartFromMongoByID(guestCartID int64) []Item {
	db := config.Conf.GetMongoDb()

	cart := Cart{Items: []Item{}}
	err := db.Collection(collectionName).FindOne(context.Background(), bson.NewDocument(
		bson.EC.Int64("quote_id", guestCartID))).Decode(&cart)
	if err != nil {
		helpers.PanicErr(err)
	}
	return cart.Items
}

func GetUserCartFromMongoByID(cartId string) Cart {
	db := config.Conf.GetMongoDb()
	var cartIdInt64 int64
	if len(cartId) > 2 {
		cartIdInt, err := strconv.Atoi(cartId)
		helpers.PanicErr(err)
		cartIdInt64 = int64(cartIdInt)
	}

	cart := Cart{Items: []Item{}}
	err := db.Collection(collectionName).FindOne(context.Background(), bson.NewDocument(
		bson.EC.Int64("quote_id", cartIdInt64),
		bson.EC.String("status", "Active"))).Decode(&cart)
	if err != nil { // ToDo gets in to fucking loop
		fmt.Println("ERROR IN GetUserCartFromMongoByID: ", err)
		cartId = CreateCartInMongoDB(cartId)
		GetUserCartFromMongoByID(cartId)
	}
	return cart
}

func CheckDoesUserHasACart(userId string) (cartId string, err error) {
	db := config.Conf.GetMongoDb()

	var cart Cart
	err = db.Collection(collectionName).FindOne(nil, bson.NewDocument(
		bson.EC.String("user_id", userId),
		bson.EC.String("status", "Active"))).
		Decode(&cart)
	return cart.CartId, err
}

func updateUserCartInMongo(cartId string, item Item) {
	bsonItem, err := helpers.StructToBson(item)
	helpers.PanicErr(err)

	db := config.Conf.GetMongoDb()

	cartIdInt, err := strconv.Atoi(cartId)
	helpers.PanicErr(err)
	cartIdInt64 := int64(cartIdInt)

	_, err = db.Collection(collectionName).UpdateOne(nil,
		bson.NewDocument(
			bson.EC.Int64("quote_id", cartIdInt64),
			bson.EC.String("status", "Active")),
		bson.NewDocument(
			bson.EC.SubDocumentFromElements("$pull",
				bson.EC.SubDocument("items",
					bson.NewDocument(
						bson.EC.String("sku", item.SKU))))))
	helpers.PanicErr(err)

	_, err = db.Collection(collectionName).UpdateOne(nil,
		bson.NewDocument(
			bson.EC.Int64("quote_id", cartIdInt64),
			bson.EC.String("status", "Active")),
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

	_, err = db.Collection(collectionName).UpdateOne(nil,
		bson.NewDocument(
			bson.EC.ObjectID("_id", bsonCartID)),
		bson.NewDocument(
			bson.EC.SubDocumentFromElements("$pull",
				bson.EC.SubDocument("items",
					bson.NewDocument(
						bson.EC.String("sku", item.SKU))))))
	helpers.PanicErr(err)

	_, err = db.Collection(collectionName).UpdateOne(nil,
		bson.NewDocument(
			bson.EC.ObjectID("_id", bsonCartID)),
		bson.NewDocument(
			bson.EC.SubDocumentFromElements("$push",
				bson.EC.SubDocument("items", bsonItem))))
	helpers.PanicErr(err)
}

func deleteItemFromCartInMongo(cartId string, item CustomerCart) {
	cartIdInt, err := strconv.Atoi(cartId)
	helpers.PanicErr(err)
	cartIdInt64 := int64(cartIdInt)

	db := config.Conf.GetMongoDb()
	_, err = db.Collection(collectionName).UpdateOne(nil,
		bson.NewDocument(
			bson.EC.Int64("quote_id", cartIdInt64),
			bson.EC.String("status", "Active")),
		bson.NewDocument(
			bson.EC.SubDocumentFromElements("$pull",
				bson.EC.SubDocument("items",
					bson.NewDocument(
						bson.EC.String("sku", item.Item.SKU))))))
	helpers.PanicErr(err)
}

func deleteItemFromGuestCartInMongo(cartID string, item CustomerCart) {
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

func UpdateCartStatus(cartId int64) {
	db := config.Conf.GetMongoDb()
	_, err := db.Collection(collectionName).UpdateOne(nil,
		bson.NewDocument(
			bson.EC.Int64("quote_id", cartId),
			bson.EC.String("status", "Active")),
		bson.NewDocument(
			bson.EC.SubDocumentFromElements("$set",
				bson.EC.String("status", "Inactive"))))
	helpers.PanicErr(err)
}
