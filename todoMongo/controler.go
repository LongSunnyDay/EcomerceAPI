package todoMongo

import (
	"context"
	"fmt"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/objectid"
	"go-api-ws/config"
	"go-api-ws/helpers"
	m "go-api-ws/todoMongo/models"
)


// CONNECTIONSTRING DB connection string
//const CONNECTIONSTRING = "mongodb://localhost:27017"

// DBNAME Database name
//const DBNAME = "go-api-ws"

// COLLNAME Collection name
const COLLNAME = "todos"

//var db *mongo.Database

// Connect establish a connection to database
//func init() {
//	client, err := mongo.NewClient(CONNECTIONSTRING)
//	helpers.PanicErr(err)
//
//	err = client.Connect(context.Background())
//	helpers.PanicErr(err)
//
//	// Collection types can be used to access the database
//	db = client.Database(DBNAME)
//}

func InsertTodo(todo m.Todo) {
	fmt.Println(todo)
	db := config.Conf.GetMongoDb()

	_, err := db.Collection(COLLNAME).InsertOne(context.Background(), todo)

	helpers.PanicErr(err)

}
func GetOneTodo(todoId string) interface{}{
	var todo m.Todo

	objId, err := objectid.FromHex(todoId)
	filter := bson.NewDocument(bson.EC.ObjectID("_id", objId))
	fmt.Println(filter)

	db := config.Conf.GetMongoDb()

	err = db.Collection(COLLNAME).FindOne(context.Background(), filter).Decode(&todo)
	helpers.PanicErr(err)
	fmt.Println(todo)

	fmt.Println(todo.ObjectId)

	return  todo
}

func GetAllTodos() []m.Todo {
	db := config.Conf.GetMongoDb()

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

func ReplaceTodoByID(todo m.Todo, todoId string) interface{}{
	var todoUpdated m.Todo

	objId, err := objectid.FromHex(todoId)
	helpers.PanicErr(err)
	filter := bson.NewDocument(bson.EC.ObjectID("_id", objId))
	fmt.Println(filter)

	db := config.Conf.GetMongoDb()

	err = db.Collection(COLLNAME).FindOneAndReplace(context.Background(), filter, todo).Decode(&todoUpdated)
	fmt.Println(todoUpdated)
	helpers.PanicErr(err)

	return  todoUpdated
}

func UpdateTodoById(todo m.Todo, todoId string) interface{}{

	var todoUpdated m.Todo
	bsonTodo, err := helpers.StructToBson(todo)
	helpers.PanicErr(err)

	objId, err := objectid.FromHex(todoId)
	helpers.PanicErr(err)
	filter := bson.NewDocument(bson.EC.ObjectID("_id", objId))

	db := config.Conf.GetMongoDb()

	err = db.Collection(COLLNAME).FindOneAndUpdate(context.Background(), filter, bson.NewDocument(bson.EC.SubDocument("$set", bsonTodo))).Decode(&todoUpdated)
	fmt.Println(todoUpdated)
	helpers.PanicErr(err)

	return  todoUpdated
}

// deletes an existing todo
func DeleteTodo(todoId string) interface{}{
	var todo m.Todo
	objId, err := objectid.FromHex(todoId)
	filter := bson.NewDocument(bson.EC.ObjectID("_id", objId))
	fmt.Println(filter)

	db := config.Conf.GetMongoDb()

	err = db.Collection(COLLNAME).FindOneAndDelete(context.Background(), filter).Decode(&todo)
	if err != nil {
		panic(err)
	}
	 {
		fmt.Println("todo ID: "+ todoId  +" has been deleted")
	}

	return nil
}