package requestsession

import (
	"encoding/json"
	"fmt"
	"livelib/constants"
	"math/rand"
	"net/http"
	"os"
	"sync"
)

type StoredCookies struct {
	uid           int
	cookies       []*http.Cookie
	ready         bool
	mu            sync.Mutex
	is_recaptched bool
}

func AddCookiesToRequest(req *RequestSession) error {
	randPoolInd := rand.Intn(constants.CookiePoolSize)
	cookies, err := storedCookiesPool[randPoolInd].getStoredCookies()

	if err != nil {
		return err
	}
	req.Cli.CookieJar().AddCookies(cookies)

	return nil
}

func (sc *StoredCookies) filename() string {
	return fmt.Sprintf("./files/cookies%d.json", sc.uid)
}

func (sc *StoredCookies) getStoredCookies() ([]*http.Cookie, error) {
	if sc.ready {
		return sc.cookies, nil
	}

	sc.mu.Lock()
	defer sc.mu.Unlock()

	file, err := os.OpenFile(sc.filename(), os.O_RDONLY|os.O_CREATE, 0644)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	json.NewDecoder(file).Decode(&sc.cookies)
	if len(sc.cookies) == 0 {
		sc.is_recaptched = true
	}
	sc.ready = true

	return sc.cookies, nil

}

func (sc *StoredCookies) addHttpCookies(cookiesTo []*http.Cookie) error {
	file, err := os.OpenFile(sc.filename(), os.O_WRONLY, 0644)

	if err != nil {
		return err
	}

	defer file.Close()

	err = json.NewEncoder(file).Encode(cookiesTo)

	if err != nil {
		return err
	}

	sc.ready = false
	sc.getStoredCookies()
	return nil
}
