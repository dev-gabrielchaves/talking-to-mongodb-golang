package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Class repr in GoLang
type User struct {
	Name   string  `bson:"name"`
	Age    int     `bson:"age"`
	Salary float64 `bson:"salary"`
}

func main() {

	// Connecting to mongoDB
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))

	// Error handler
	if err != nil {
		panic(err)
	}

	// Creating/Defining collection
	collection := client.Database("user").Collection("userData")

	// Defining object from class User
	gabrielUser := User{
		Name:   "Gabriel",
		Age:    25,
		Salary: 1000.00,
	}

	// Inserting user to DB
	result, err := collection.InsertOne(context.Background(), gabrielUser)

	// Error handler
	if err != nil {
		panic(err)
	}

	// Showing ID of the added user in DB
	fmt.Println(result.InsertedID)
}
