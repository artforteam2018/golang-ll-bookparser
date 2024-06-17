package main

import (
	databasescripts "livelib/backs/database_scripts"
	"livelib/components/mongo"
	"livelib/components/requestsession"
	"livelib/models/book"
	"livelib/models/livelib_ex"
)

func bootstrap() {
	if err := mongo.ConnectDB(); err != nil {
		panic(err)
	}

	loadModels()
	databasescripts.FillDatabase()

	requestsession.InitStoredCookiesPool()
}

func loadModels() {
	if err := book.BookStorage.Init(); err != nil {
		panic(err)
	}
	if err := livelib_ex.LivelibExStorage.Init(); err != nil {
		panic(err)
	}

}
