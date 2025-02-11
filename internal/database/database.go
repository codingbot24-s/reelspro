package database

import (
	"context"

	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client = ConnectToDatabase()

func ConnectToDatabase() *mongo.Client {
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("MONGO_URI is not set")
		panic("MONGO_URI is not set")
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))

	if err != nil {
		log.Fatal("error connecting to database: ", err)
		return nil
	}

	// Ping the database to verify connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("error pinging database: ", err)
		return nil
	}

	log.Println("Successfully connected to MongoDB!")
	return client
}

func DisconnectDatabase() {
	if Client != nil {
		if err := Client.Disconnect(context.TODO()); err != nil {
			log.Fatal("Error disconnecting from database:", err)
		}
	}
}
