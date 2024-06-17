package databasescripts

import (
	"context"
	"fmt"
	"livelib/constants"
	"livelib/models/livelib_ex"

	"go.mongodb.org/mongo-driver/bson"
)

func FillDatabase() {
	allCount, err := livelib_ex.LivelibExStorage.Coll.CountDocuments(context.TODO(), bson.M{})

	if err != nil {
		panic(err)
	}

	if allCount != constants.AllBookCount {
		fmt.Println("FILL DB WITH -- Start")

		docs := bson.A{}
		y := 0
		for i := 0; i < constants.AllBookCount; i++ {
			docs = append(docs, livelib_ex.LivelibEx{Id: constants.FirstBookId + i, Exists: false, Scanned: false})

			if y > 100000 {
				fmt.Println("inserted", i, "docs")
				_, err := livelib_ex.LivelibExStorage.Coll.InsertMany(context.TODO(), docs)
				if err != nil {
					panic(err)
				}
				y = 0
				docs = bson.A{}
			}
			y++
		}
		_, err := livelib_ex.LivelibExStorage.Coll.InsertMany(context.TODO(), docs)
		if err != nil {
			panic(err)
		}
		fmt.Println("FILL DB WITH -- Done")
	}
}
