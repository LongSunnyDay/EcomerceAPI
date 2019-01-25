package payment_methods

import (
	"context"
	"github.com/mongodb/mongo-go-driver/bson"
	"go-api-ws/config"
	"go-api-ws/helpers"
)

func insertPaymentMethodsToMongo(methods []interface{}) {
	db := config.Conf.GetMongoDb()

	_, err := db.Collection(collectionName).InsertMany(nil, methods)
	helpers.PanicErr(err)
}

func getPaymentMethodsFromMongo() []Method {
	var paymentMethod Method
	var paymentMethods []Method

	db := config.Conf.GetMongoDb()

	cur, err := db.Collection(collectionName).Find(nil, bson.NewDocument(
		bson.EC.String("type", "Payment Method")))
	helpers.PanicErr(err)
	for cur.Next(context.Background()) {
		err := cur.Decode(&paymentMethod)
		helpers.PanicErr(err)
		paymentMethods = append(paymentMethods, paymentMethod)
	}
	err = cur.Close(context.Background())
	helpers.CheckErr(err)
	return paymentMethods
}

func (m Method) insertToDb() {
	db, err := config.Conf.GetDb()
	helpers.PanicErr(err)
	_, err = db.Exec("INSERT INTO payment_methods("+
		"Code, "+
		"Title, "+
		"IsServerMethod)"+
		" VALUES(?, ?, ?)",
		m.Code, m.Title, m.IsServerMethod)
	helpers.PanicErr(err)
}

func getPaymentMethodsFromDb() Methods {
	db, err := config.Conf.GetDb()
	helpers.PanicErr(err)
	rows, err := db.Query("SELECT Code, Title, IsServerMethod FROM payment_methods")
	helpers.PanicErr(err)
	var methods []Method
	defer rows.Close()
	for rows.Next() {
		var method Method
		if err := rows.Scan(&method.Code, &method.Title, &method.IsServerMethod); err != nil {
			helpers.PanicErr(err)
		}
		methods = append(methods, method)
	}
	return methods
}

func (m Method) updatePaymentMethodInDb() {
	db, err := config.Conf.GetDb()
	helpers.PanicErr(err)
	_, err = db.Exec("UPDATE payment_methods p SET "+
		"p.Code = ?, "+
		"p.Title = ?, "+
		"p.IsServerMethod = ? "+
		"WHERE p.Id = ?",
		m.Code, m.Title, m.IsServerMethod, m.Id)
	helpers.PanicErr(err)
}

func removePaymentMethodFromDb(id string) {
	db, err := config.Conf.GetDb()
	helpers.PanicErr(err)
	_, err = db.Exec("DELETE FROM payment_methods  WHERE Id = ?", id)
	helpers.PanicErr(err)
}

func GetActualPaymentMethodsFromDb() Methods {
	db, err := config.Conf.GetDb()
	helpers.PanicErr(err)
	rows, err := db.Query("SELECT Code, Title, IsServerMethod FROM payment_methods WHERE id != 17")
	helpers.PanicErr(err)
	var methods []Method
	for rows.Next() {
		var method Method
		if err := rows.Scan(&method.Code, &method.Title, &method.IsServerMethod); err != nil {
			helpers.PanicErr(err)
		}
		methods = append(methods, method)
	}
	return methods
}

func GetPaymentMethodFromDbByMethodCode(methodCode string) Method {
	db, err := config.Conf.GetDb()
	helpers.PanicErr(err)
	var method Method
	err = db.QueryRow("SELECT Code, Title, IsServerMethod FROM payment_methods WHERE code = ?", methodCode).Scan(&method.Code, &method.Title, &method.IsServerMethod)
	helpers.PanicErr(err)
	return method
}
