package db

import (
	"context"
	"log"
	"time"

	"NokiaAssesmentGo/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

// CreatePersonEndpoint will insert data into DB
func CreatePersonEndpoint(person utils.Person) *mongo.InsertOneResult {
	collection := client.Database("NokiaAssesmentGo").Collection("people")
	ctx, CancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer CancelFunc()
	result, err := collection.InsertOne(ctx, person)
	if err != nil {
		log.Println("Insert Person failed: ", err)
	}
	return result
}

// GetPersonFromDB from DB
func GetPersonFromDB() {
	collection := client.Database("NokiaAssesmentGo").Collection("people")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var person utils.Person
		cursor.Decode(&person)
		utils.StoreInCache(person.ID.Hex(), &person)
	}
}

//ConnectToDB connect to DB
func ConnectToDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, _ = mongo.Connect(ctx, clientOptions)

}
