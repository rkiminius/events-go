package db

import (
	"context"
	"events/env"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

const defaultDbName = "events"

const MongoTimeout = 5 * time.Second

var client *mongo.Client

// MongoConnect used to establish a connection to MongoDB server using the provided host and port information from the environment variables.
func MongoConnect() *mongo.Client {

	if client != nil {
		return client
	}

	uri := fmt.Sprintf("mongodb://%s:%s", env.Env["DB_HOST"], env.Env["DB_PORT"])
	clientOpts := options.Client().ApplyURI(uri)

	var err error
	client, err = mongo.Connect(context.TODO(), clientOpts)
	if err != nil {
		log.Fatal(err)
	}

	if err = client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	return client
}

func GetDatabase() *mongo.Database {
	dbName := env.Env["DB_NAME"]
	if dbName == "" {
		dbName = defaultDbName
	}

	return MongoConnect().Database(dbName)
}

func GetMongoCollection(name string) *mongo.Collection {
	return GetDatabase().Collection(name)
}

func GetTimeoutContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), MongoTimeout)
}

// MongoDisconnect used to close the MongoDB client connection.
func MongoDisconnect() {
	if client == nil {
		return
	}

	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MongoDB connection closed!")
}
