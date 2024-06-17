package book

import (
	"livelib/components/mongo"
)

type Info struct {
	Type  string `bson:"type,omitempty"`
	Value string `bson:"value,omitempty"`
}

type Author struct {
	Name string `bson:"name,omitempty"`
	Link string `bson:"link,omitempty"`
}

type Genre struct {
	Name string `bson:"name,omitempty"`
	Link string `bson:"link,omitempty"`
}

type Book struct {
	Id          int      `bson:"id,omitempty"`
	TextId      string   `bson:"textid,omitempty"`
	Title       string   `bson:"title,omitempty"`
	Authors     []Author `bson:"authors,omitempty"`
	Editions    []Info   `bson:"editions,omitempty"`
	Image       string   `bson:"image,omitempty"`
	Infos       []Info   `bson:"infos,omitempty"`
	Genres      []Genre  `bson:"genres,omitempty"`
	Description string   `bson:"description,omitempty"`
}

var BookStorage mongo.MongoStorage[Book] = mongo.MongoStorage[Book]{Collname: "books", Coll: nil}
