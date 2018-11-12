package todoMongo

import (
	"context"
	"fmt"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/objectid"
	"github.com/mongodb/mongo-go-driver/mongo"
	"go-api-ws/helpers"
	m "go-api-ws/todoMongo/models"
	)


// CONNECTIONSTRING DB connection string
const CONNECTIONSTRING = "mongodb://localhost:27017"

// DBNAME Database name
const DBNAME = "go-api-ws"

// COLLNAME Collection name
const COLLNAME = "todos"

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

func InsertTodo(todo m.Todo) {
	fmt.Println(todo)
	_, err := db.Collection(COLLNAME).InsertOne(context.Background(), todo)

	helpers.PanicErr(err)

}
func GetOneTodo(todoId string) interface{}{
	var todo m.Todo

	objId, err := objectid.FromHex(todoId)
	filter := bson.NewDocument(bson.EC.ObjectID("_id", objId))
	fmt.Println(filter)

	err = db.Collection(COLLNAME).FindOne(context.Background(), filter).Decode(&todo)
	helpers.PanicErr(err)
	fmt.Println(todo)

	fmt.Println(todo.ObjectId)

	return  todo
}

func GetAllTodos() []m.Todo {
	cur, err := db.Collection(COLLNAME).Find(context.Background(), nil, nil)
	helpers.PanicErr(err)

	var elements []m.Todo
	var elem m.Todo
	// Get the next result from the cursor
	for cur.Next(context.Background()) {
		err := cur.Decode(&elem)
		helpers.PanicErr(err)

		elements = append(elements, elem)
	}
	helpers.PanicErr(err)

	cur.Close(context.Background())
	return elements
}

func UpdateTodoByID(todo m.Todo, todoId string)  interface{}{
	var todoUpdated m.Todo


	objId, err := objectid.FromHex(todoId)
	helpers.PanicErr(err)
	filter := bson.NewDocument(bson.EC.ObjectID("_id", objId))
	fmt.Println(filter)

	err = db.Collection(COLLNAME).FindOneAndReplace(context.Background(), filter, todo).Decode(&todoUpdated)

	fmt.Println(todoUpdated)

	helpers.PanicErr(err)


	return  todoUpdated

}

// deletes an existing todo
func DeleteTodo(todoId string) {
	objectIDS, err := objectid.FromHex(todoId)
	helpers.PanicErr(err)
	idDoc := bson.NewDocument(bson.EC.ObjectID("_id", objectIDS))
	_, err = db.Collection(COLLNAME).DeleteOne(context.Background(), idDoc, nil)
	helpers.PanicErr(err)
}





//func UpdateTodoByID(todo m.Todo, todoID string) {
//	doc := db.Collection(COLLNAME).FindOneAndUpdate(
//		context.Background(),
//		bson.NewDocument(
//			bson.EC.String("id", todoID),
//		),
//		bson.NewDocument(
//			bson.EC.SubDocumentFromElements("$set",
//				bson.EC.String("id", todo.ID),
//				bson.EC.String("title", todo.Title),
//				bson.EC.String("category", todo.Category),
//				bson.EC.String("content", todo.Content),
//				bson.EC.String("state", todo.State)),
//			),
//			nil)
//	fmt.Println(doc)
//}


