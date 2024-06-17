package mongo

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var singleMongoClient *mongo.Client
var collections map[string]int

func ConnectDB() error {

	timer := time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(""))
	if err != nil {
		return fmt.Errorf("connect to mongodb error: %s", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return fmt.Errorf("connect to mongodb error: %s", err)
	}

	fmt.Println("connect to mongodb successful in: ", time.Since(timer))
	singleMongoClient = client
	collections = make(map[string]int)

	return nil
}

func GetDBClient() *mongo.Database {
	db := singleMongoClient.Database("main")

	return db
}
