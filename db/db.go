package db

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	middlewares "github.com/umangraval/Go-Mongodb-REST-boilerplate/handlers"
	"github.com/umangraval/Go-Mongodb-REST-boilerplate/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io/ioutil"
	"log"
	"path/filepath"
	"reflect"
)

var client *mongo.Client

// Dbconnect -> connects mongo
func Dbconnect() *mongo.Client {
	clientOptions := options.Client().ApplyURI(middlewares.DotEnvVariable("MONGO_URL"))
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal("⛒ Connection Failed to Database")
		log.Fatal(err)
	}
	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("⛒ Connection Failed to Database")
		log.Fatal(err)
	}
	color.Green("⛁ Connected to Database")

	loadDataFromJsonFile(err, client)

	return client
}

func loadDataFromJsonFile(err error, client *mongo.Client) {
	// Load values from JSON file to model
	docsPath, _ := filepath.Abs("db/db.json")

	byteValues, err := ioutil.ReadFile(docsPath)
	if err != nil {
		// Print any IO errors with the .json file
		fmt.Println("ioutil.ReadFile ERROR:", err)
	} else {
		// Access a MongoDB collection through a database
		col := client.Database("wajve").Collection("trivia")
		fmt.Println("Collection type:", reflect.TypeOf(col), "n")

		// Print the values of the JSON docs, and insert them if no error
		fmt.Println("ioutil.ReadFile byteValues TYPE:", reflect.TypeOf(byteValues))
		fmt.Println("byteValues:", byteValues, "n")
		fmt.Println("byteValues:", string(byteValues))
		var docs []models.Trivia
		// Unmarshal the encoded JSON byte string into the slice
		err = json.Unmarshal(byteValues, &docs)
		fmt.Println("nMongoFields Docs:", reflect.TypeOf(docs))
		for i := range docs {
			// Put the document element in a new variable
			doc := docs[i]
			fmt.Println("ndoc _id:", doc.Text)
			fmt.Println("doc Field Str:", doc.Text)

			// Call the InsertOne() method and pass the context and doc objects
			result, insertErr := col.InsertOne(context.TODO(), doc)
			// Check for any insertion errors
			if insertErr != nil {
				fmt.Println("InsertOne ERROR:", insertErr)
			} else {
				fmt.Println("InsertOne() API result:", result)
			}
		}
	}
}
