package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoInstance struct {
	Client *mongo.Client
	DB     *mongo.Database
}

func ConnectDB() *mongo.client {
	database := os.Getenv("DB_URL")

	client, err := mongo.NewClient(options.Client().ApplyURI(database))
	if err != nil {
		log.Fatal(err)
	}

	//we increment the connection time to 25 seconds
	ctx, _ := context.WithTimeout(context.Background(), 25*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB")

	return client
}
