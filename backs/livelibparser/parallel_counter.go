package livelibparser

import (
	"livelib/constants"
	"sync/atomic"
)

var atomCounter int32 = 0

/*
returns true if counter incremented
returns false if counter is already full or cannot be incremented
*/
func incrParallelCounter() bool {
	if atomCounter < constants.MaxParallelworkers {
		return atomic.CompareAndSwapInt32(&atomCounter, atomCounter, atomCounter+1)
	}
	return false
}

func decrParallelCounter() {
	atomic.AddInt32(&atomCounter, -1)
}
