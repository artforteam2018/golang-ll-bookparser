package mongo

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoStorage[T interface{}] struct {
	Collname string
	Coll     *mongo.Collection
}

func (storage *MongoStorage[T]) Init() error {

	if _, ok := collections[storage.Collname]; ok {
		return fmt.Errorf("collection %s is already associated", storage.Collname)
	}

	collections[storage.Collname] = 1

	storage.Coll = GetDBClient().Collection(storage.Collname)

	return nil
}

func (storage *MongoStorage[T]) InsertOne(data T) (string, error) {
	res, err := storage.Coll.InsertOne(context.Background(), data)
	if err != nil {
		return "", fmt.Errorf("error mongo insertOne: %s", err)
	}
	id := res.InsertedID.(primitive.ObjectID)

	return string(id[:]), nil
}

func (storage *MongoStorage[T]) FindOne(dataInput *T, search interface{}, opts *options.FindOneOptions) error {
	if search == nil {
		search = bson.M{}
	}
	res := storage.Coll.FindOne(context.TODO(), search, opts)

	if err := res.Err(); err != nil {
		return err
	}

	if err := res.Decode(dataInput); err != nil {
		return err
	}
	return nil
}

func (storage *MongoStorage[T]) Find(dataInput *[]*T, search interface{}, opts *options.FindOptions) error {
	if search == nil {
		search = bson.M{}
	}
	cursor, err := storage.Coll.Find(context.TODO(), search, opts)

	if err != nil {
		return err
	}

	if err := cursor.All(context.TODO(), dataInput); err != nil {
		return err
	}

	// // dataInputVal := *dataInput
	// for cursor.Next(context.TODO()) {
	// 	d := new(T)

	// 	if err := cursor.Decode(d); err != nil {
	// 		return err
	// 	}
	// 	*dataInput = append(*dataInput, d)
	// }
	return nil
}
