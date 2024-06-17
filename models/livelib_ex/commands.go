package livelib_ex

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func FindUnscannedValues(limit int) ([]*LivelibEx, error) {
	booksToParse := []*LivelibEx{}

	lim64 := int64(limit)

	if err := LivelibExStorage.Find(&booksToParse, bson.M{"scanned": false}, &options.FindOptions{Limit: &lim64}); err != nil {
		return nil, err
	}

	return booksToParse, nil
}

func RemoveFoundedValue(id int) error {
	updateValue := bson.M{"exists": false, "scanned": false}
	_, err := LivelibExStorage.Coll.UpdateOne(context.TODO(), bson.M{"id": id}, bson.M{"$set": updateValue})

	return err
}

func UpdateFoundedValue(id int, exists bool) error {
	updateValue := bson.M{"exists": exists, "scanned": true}
	_, err := LivelibExStorage.Coll.UpdateOne(context.TODO(), bson.M{"id": id}, bson.M{"$set": updateValue})

	return err
}
