package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Currency struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Code         string             `bson:"code,omitempty"`
	ExchangeRate float64            `bson:"exchange_rate,omitempty"`
	Date         time.Time          `bson:"date,omitempty"`
	Base         bool               `bson:"base,omitempty"`
}

func main() {

	// create a new context
	ctx := context.Background()

	// create a mongo client
	// create a mongo client
	client, err := mongo.Connect(
		ctx,
		options.Client().ApplyURI("mongodb://localhost:27017/"),
	)
	if err != nil {
		log.Fatal(err)
	}

	// disconnects from mongo
	defer client.Disconnect(ctx)

	db := client.Database("budget")
	col := db.Collection("currencies")

	// find all documents
	cursor, err := col.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("cursor is done %+v\n", cursor)
	var curs []Currency
	// iterate through all documents
	for cursor.Next(ctx) {
		var cur Currency
		// decode the document
		if err := cursor.Decode(&cur); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("currency: %+v\n", cur)
		curs = append(curs, cur)
	}

	log.Printf("currencies %+v", curs)

	// check if the cursor encountered any errors while iterating
	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}

}
