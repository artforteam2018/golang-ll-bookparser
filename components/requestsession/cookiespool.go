package requestsession

import (
	"livelib/components/chromewebpageemu"
	"livelib/constants"
	"net/http"
	"sync"
)

var storedCookiesPool = make([]StoredCookies, constants.CookiePoolSize)

func InitStoredCookiesPool() {
	for i := range storedCookiesPool {
		storedCookiesPool[i] = StoredCookies{i, []*http.Cookie{}, false, sync.Mutex{}, false}
		storedCookiesPool[i].getStoredCookies()
	}
	RestoreBadCookies()
}

func RestoreBadCookies() {
	for i := range storedCookiesPool {

		if !storedCookiesPool[i].is_recaptched {
			continue
		}
		cookiesTo, err := chromewebpageemu.EmulateRunWebpage()
		if err != nil {
			panic(err)
		}
		storedCookiesPool[i].addHttpCookies(cookiesTo)
		// livelibwebpageiterator.CheckLivelibResponse(cookies, constants.FirstBookId)
	}
}
