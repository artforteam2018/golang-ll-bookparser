package constants

import "time"

const FirstBookId = 1000000001
const AllBookCount = 2302050
const LastBookId = FirstBookId + AllBookCount
const LivelibBatchCount = 100

const MaxRetryRecoverCount = 2

// parsing timeouts and sleeps
const ParseTimeout = time.Second * 12

// parallel livelib parsing
const MaxParallelworkers = 10
const CookiePoolSize = 5

const PauseForPageRender = time.Second * 2
