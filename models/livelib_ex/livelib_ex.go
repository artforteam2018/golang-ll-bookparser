package livelib_ex

import (
	"livelib/components/mongo"
)

type LivelibEx struct {
	Id      int  `bson:"id,omitempty"`
	Exists  bool `bson:"exists"`
	Scanned bool `bson:"scanned"`
}

var LivelibExStorage mongo.MongoStorage[LivelibEx] = mongo.MongoStorage[LivelibEx]{Collname: "livelibex", Coll: nil}
