package todoMongo

import (
	"github.com/mongodb/mongo-go-driver/mongo"
		"context"
	"fmt"
	"go-api-ws/todoMongo/models"
	"go-api-ws/helpers"
	"github.com/mongodb/mongo-go-driver/bson"
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
	helpers.CheckErr(err)

	err = client.Connect(context.Background())
	helpers.CheckErr(err)

	// Collection types can be used to access the database
	db = client.Database(DBNAME)
}

func InsertTodo(todo models.Todo) {
	fmt.Println(todo)
	_, err := db.Collection(COLLNAME).InsertOne(context.Background(), todo)

	helpers.CheckErr(err)

}

func GetAllTodos() []models.Todo {
	cur, err := db.Collection(COLLNAME).Find(context.Background(), nil, nil)
	helpers.CheckErr(err)

	var elements []models.Todo
	var elem models.Todo
	// Get the next result from the cursor
	for cur.Next(context.Background()) {
		err := cur.Decode(&elem)
		helpers.CheckErr(err)

		elements = append(elements, elem)
	}
	helpers.CheckErr(err)

	cur.Close(context.Background())
	return elements
}

// deletes an existing todo
func DeleteTodo(todo models.Todo) {
	_, err := db.Collection(COLLNAME).DeleteOne(context.Background(), todo, nil)
	helpers.CheckErr(err)
}

func UpdateTodoByID(todo models.Todo, todoID string) {
	doc := db.Collection(COLLNAME).FindOneAndUpdate(
		context.Background(),
		bson.NewDocument(
			bson.EC.String("id", todoID),
		),
		bson.NewDocument(
			bson.EC.SubDocumentFromElements("$set",
				bson.EC.String("Title", todo.Title),
				bson.EC.String("Category", todo.Category),
				bson.EC.String("Content", todo.Content),
				bson.EC.String("State", todo.State)),
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
//	var schemaLoader = gojsonschema.NewReferenceLoader("file://todo/models/todoSchema.json")
//	_ = json.NewDecoder(r.Body).Decode(&todo)
//	documentLoader := gojsonschema.NewGoLoader(todo)
//	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
//	helpers.CheckErr(err)
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
////	if err := models.Insert(movie); err != nil {
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


