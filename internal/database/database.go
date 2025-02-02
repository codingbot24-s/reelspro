package database

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/mongo"
)

var Client *mongo.Client = ConnectToDatabase()

func ConnectToDatabase() *mongo.Client {
	
	uri := os.Getenv("MONGO_URI")

	if uri == "" {
		panic("MONGO_URI is not set")
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	if err != nil {
		panic(err)
	}

	return client

}
