package db

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var client *mongo.Client

func ConnectDb() (*mongo.Database, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://root:password@mongo:27010/?authSource=admin&maxPoolSize=20&w=majority"))

	if err != nil {
		panic(err)
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}

	fmt.Println("Successfully Connect MongoDB")

	db := client.Database("users")

	return db, nil
}

func CloseMongo() {
	if client == nil {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	if err := client.Disconnect(ctx); err != nil {
		panic(err)
	}

	fmt.Println("Connection Close")

}
