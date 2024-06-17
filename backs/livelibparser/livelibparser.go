package livelibparser

import (
	"context"
	"fmt"
	"livelib/backs/book_and_ex"
	"livelib/components/livelibwebpageiterator"
	"livelib/components/requestsession"
	"livelib/constants"
	"livelib/models/livelib_ex"
	"sync"
	"time"
)

var recoverMu = sync.Mutex{}
var recoverMuWaiter = sync.Mutex{}

func StartBatchParsing() {
	for {
		timer := time.Now()
		booksToParse, err := livelib_ex.FindUnscannedValues(constants.LivelibBatchCount)
		if err != nil {
			panic(err)
		}

		if len(booksToParse) == 0 {
			break
		}

		for _, data := range booksToParse {
			for {
				if incrParallelCounter() {
					break
				}
			}

			go runWorkerWithTimeout(data.Id)

		}

		fmt.Println(constants.LivelibBatchCount, "values parsed at:", time.Since(timer))
	}
}

func runWorkerWithTimeout(v int) {
	// timer := time.Now()
	// create channel for end of goroutine signal
	ch := make(chan int)

	//create timeout
	ctx, cancel := context.WithTimeout(context.Background(), constants.ParseTimeout)

	defer decrParallelCounter()
	defer cancel()

	go func() {
		//recover after recaptcha
		defer recoverLivelib(v, 0)

		Book, err := livelibwebpageiterator.ParseLivelibId(ctx, v)
		book_and_ex.WriteToModel(err, Book, v)

		// statsReport := fmt.Sprintf("%d  %s", v, time.Since(timer))
		// if err != nil {
		// 	statsReport = fmt.Sprintf("%s %s", statsReport, err.Error())
		// } else {
		// 	statsReport = fmt.Sprintf("%s success", statsReport)
		// }

		// fmt.Println(statsReport)
		close(ch)
	}()

	//end func after timeout or success
	select {
	case <-ctx.Done():
	case <-ch:
	}

}

func recoverLivelib(id int, retryCount int) {
	if r := recover(); r != nil {
		if r != "got recaptcha" {
			panic(r)
		}
		if retryCount > constants.MaxRetryRecoverCount {
			panic(r.(string) + " recover failed")
		}

		// try to recover 1 more time
		defer recoverLivelib(id, retryCount+1)
		defer recoverMuWaiter.Unlock()

		// lock one proc to solve recaptcha
		if recoverMu.TryLock() {
			defer recoverMu.Unlock()

			// lock all other go's
			recoverMuWaiter.Lock()
			requestsession.RestoreBadCookies()
		} else {
			// other go's wait here for solve recaptcha
			recoverMuWaiter.Lock()
			return
		}
	}
}
