package book_and_ex

import (
	"fmt"
	"livelib/models/book"
	"livelib/models/livelib_ex"
	"regexp"
)

var regularSearch, _ = regexp.Compile(`404|503|stopped after 10 redirects`)

func WriteToModel(err error, Book *book.Book, bookId int) {
	writeToEx := err == nil || regularSearch.MatchString(err.Error())
	writeToBs := err == nil

	// todo transaction || outbox
	if writeToEx {
		err := livelib_ex.UpdateFoundedValue(bookId, err == nil)

		if err != nil {
			fmt.Println("write to ex err", err.Error())
			return
		}
	} else {
		fmt.Println("unusual error write", err)
	}

	if writeToBs {
		_, err := book.BookStorage.InsertOne(*Book)

		if err != nil {
			fmt.Println("write to bs err", err.Error())
			err := livelib_ex.RemoveFoundedValue(bookId)

			if err != nil {
				panic("Cannot revert operation: " + err.Error())
			}

		}
	}
}
