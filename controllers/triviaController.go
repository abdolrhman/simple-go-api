package controllers

import (
	"context"
	"fmt"
	"github.com/abdolrhman/simple-go-api/db"
	middlewares "github.com/abdolrhman/simple-go-api/handlers"
	"github.com/abdolrhman/simple-go-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"net/http"
)

var client = db.Dbconnect()

// GetTriviaEndpoint -> get Trivia
var GetTriviaEndpoint = http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
	var people []*models.Trivia
	collection := client.Database("wajve").Collection("trivia")

	text := request.URL.Query().Get("text")
	number := request.URL.Query().Get("number")

	filter := bson.M{}
	if text != "" {
		filter["text"] = text
	}
	if number != "" {
		filter["number"] = number
	}

	fmt.Println(text, number)

	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		middlewares.ServerErrResponse(err.Error(), response)
		return
	}
	for cursor.Next(context.TODO()) {
		var person models.Trivia
		err := cursor.Decode(&person)
		if err != nil {
			log.Fatal(err)
		}

		people = append(people, &person)
	}
	if err := cursor.Err(); err != nil {
		middlewares.ServerErrResponse(err.Error(), response)
		return
	}
	middlewares.SuccessArrRespond(people, response)
})
