package todoMongo

import (
	"context"
	"fmt"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/objectid"
	"github.com/mongodb/mongo-go-driver/mongo"
	"go-api-ws/helpers"
	"go-api-ws/todoMongo/models"
)

// CONNECTIONSTRING DB connection string
const CONNECTIONSTRING = "mongodb://localhost:32777"

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

func InsertTodo(todo models.Todo) {
	fmt.Println(todo)
	_, err := db.Collection(COLLNAME).InsertOne(context.Background(), todo)

	helpers.PanicErr(err)

}

func GetAllTodos() []models.Todo {
	cur, err := db.Collection(COLLNAME).Find(context.Background(), nil, nil)
	helpers.PanicErr(err)

	var elements []models.Todo
	var elem models.Todo
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

// deletes an existing todo
func DeleteTodo(todoId string) {
	objectIDS, err := objectid.FromHex(todoId)
	helpers.PanicErr(err)
	idDoc := bson.NewDocument(bson.EC.ObjectID("_id", objectIDS))
	_, err = db.Collection(COLLNAME).DeleteOne(context.Background(), idDoc, nil)
	helpers.PanicErr(err)
}

func UpdateTodoByID(todo models.Todo, todoID string) {
	doc := db.Collection(COLLNAME).FindOneAndUpdate(
		context.Background(),
		bson.NewDocument(
			bson.EC.String("id", todoID),
		),
		bson.NewDocument(
			bson.EC.SubDocumentFromElements("$set",
				bson.EC.String("title", todo.Title),
				bson.EC.String("category", todo.Category),
				bson.EC.String("content", todo.Content),
				bson.EC.String("state", todo.State)),
			),
			nil)
	fmt.Println(doc)
}



















//var todo m.Todo
//
//// connect to MongoDB
////func connection() {
////
////	client, err := mongo.Connect(context.Background(), "mongodb://localhost:32768", nil)
////	if err != nil {
////		log.Fatal(err)
////	}
////	db := client.Database("go-api-ws")
////	inventory := db.Collection("todos")
////}
//
//func InsertTodo(w http.ResponseWriter, r *http.Request) {
//
//
//	client, err := mongo.Connect(context.Background(), "mongodb://localhost:32768", nil)
//	if err != nil {
//		log.Fatal(err)
//	}
//	db := client.Database("go-api-ws")
//	inventory := db.Collection("todos")
//
//
//	var schemaLoader = gojsonschema.NewReferenceLoader("file://todo/jsonSchemaModels/todoSchema.json")
//	_ = json.NewDecoder(r.Body).Decode(&todo)
//	documentLoader := gojsonschema.NewGoLoader(todo)
//	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
//	helpers.PanicErr(err)
//
//	if  result.Valid(){
//		inventory.InsertOne(context.Background(), result)
//	}
//	w.Header().Set("Content-Type", "application/json")
//	json.NewEncoder(w).Encode("Todo: " + todo.Title + " has been registered. " )
//
//
//
//
//
//
////	if err := jsonSchemaModels.Insert(movie); err != nil {
////		respondWithError(w, http.StatusInternalServerError, err.Error())
////		return
////	}
////	respondWithJson(w, http.StatusCreated, movie)
//}
//
//
//
//func AllMoviesEndPoint(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintln(w, "not implemented yet !")
//}
//
//func FindMovieEndpoint(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintln(w, "not implemented yet !")
//}
//
//
//func UpdateMovieEndPoint(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintln(w, "not implemented yet !")
//}
//
//func DeleteMovieEndPoint(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintln(w, "not implemented yet !")
//}


