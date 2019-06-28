package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Currency - struct for currency object
type Currency struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Code         string             `bson:"code,omitempty"`
	ExchangeRate float64            `bson:"exchange_rate,omitempty"`
	Date         time.Time          `bson:"date,omitempty"`
	Base         bool               `bson:"base,omitempty"`
}

type mongoDB struct {
	client *mongo.Client
	db     *mongo.Database
	dbName string
	ctx    context.Context
}

var md *mongoDB

func getCurrencies() (currs []Currency) {
	ctx := md.ctx

	col := md.db.Collection("currencies")

	// find all documents
	cursor, err := col.Find(ctx, bson.D{})
	if err != nil {
		log.Println(err)
		return nil
	}

	fmt.Printf("cursor is done %+v\n", cursor)
	var curs []Currency
	// iterate through all documents
	for cursor.Next(ctx) {
		var cur Currency
		// decode the document
		if err := cursor.Decode(&cur); err != nil {
			log.Println(err)
			return nil
		}
		fmt.Printf("currency: %+v\n", cur)
		curs = append(curs, cur)
	}

	// check if the cursor encountered any errors while iterating
	if err := cursor.Err(); err != nil {
		log.Println(err)
		return nil
	}
	return curs
}

func main() {

	// create a new context
	ctx := context.Background()
	md = &mongoDB{}
	md.dbName = "budget"

	// create a mongo client
	db, err := mongo.Connect(
		ctx,
		options.Client().ApplyURI("mongodb://localhost:27017/"),
	)
	if err != nil {
		log.Fatal(err)
	}
	md.client = db
	md.db = db.Database(md.dbName)
	md.ctx = ctx
	// disconnects from mongo
	defer md.client.Disconnect(ctx)

	// log.Printf("currencies %+v", getCurrencies())

	http.HandleFunc("/", homePage)

	http.HandleFunc("/currencies", currencies)

	port := ":8082"
	log.Printf("Server is running on port %s", port)
	http.ListenAndServe(port, nil)

}

func homePage(w http.ResponseWriter, r *http.Request) {
	log.Println("### Home page")
	defer log.Println("### Home page bye")
	fmt.Fprintf(w, "Hi there use /currencies (GET) to get currencies or POST to submit currencies")
}

func currencies(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		w.WriteHeader(http.StatusOK)
		w.Header().Add("Content-Type", "application/json")
		crs := getCurrencies()
		if err := json.NewEncoder(w).Encode(crs); err != nil {
			log.Printf(" json Problem ... %v\n", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		// w.Write([]byte(fmt.Sprintf("list of currencies - comming soon")))
		// fmt.Fprintf(w, "list of currencies - comming soon")

	case http.MethodPost:
		postCur(w, r)

	default:
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		// fmt.Fprintf(w, fmt.Sprintf("Unsupported method was called %v", r.Method))
		messages := make(map[string]string)
		messages["message"] = "Unsupported method was called"
		json.NewEncoder(w).Encode(messages)

		log.Printf("Unsupported method was called %v", r.Method)

	}

}

func postCur(w http.ResponseWriter, r *http.Request) {
	var currency Currency
	if err := json.NewDecoder(r.Body).Decode(&currency); err != nil {
		log.Printf("Problem with decoding body %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	col := md.db.Collection("currencies")

	res, err := col.InsertOne(md.ctx, currency)
	if err != nil {
		log.Printf("Problem saving %T ... %+v", currency, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	currency.ID = res.InsertedID.(primitive.ObjectID)

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(currency); err != nil {
		log.Printf("Encode problem %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
