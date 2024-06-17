package main

import (
	"livelib/backs/livelibparser"
)

func main() {

	bootstrap()
	livelibparser.StartBatchParsing()
}
