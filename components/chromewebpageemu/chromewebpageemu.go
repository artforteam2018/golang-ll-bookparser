package chromewebpageemu

import (
	"livelib/constants"
	"net/http"
	"time"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

func EmulateRunWebpage() ([]*http.Cookie, error) {
	service, err := selenium.NewChromeDriverService("chromedriver", 4444)
	if err != nil {
		panic(err)
	}
	defer service.Stop()

	caps := selenium.Capabilities{}
	caps.AddChrome(chrome.Capabilities{Args: []string{
		"window-size=1920x1080",
		"--no-sandbox",
		"--disable-dev-shm-usage",
		"disable-gpu",
		// "--headless",  // comment out this line to see the browser
	}})

	driver, err := selenium.NewRemote(caps, "")
	if err != nil {
		return nil, err
	}

	driver.Get("https://www.livelib.ru/book/1008345310")
	driver.Wait(func(wd selenium.WebDriver) (bool, error) {
		time.Sleep(constants.PauseForPageRender)
		return true, nil
	})
	cookies, err := driver.GetCookies()

	if err != nil {
		return nil, err
	}

	var cookiesReady []*http.Cookie

	for _, cookieChrome := range cookies {
		cookie := http.Cookie{
			Name:     cookieChrome.Name,
			Value:    cookieChrome.Value,
			Path:     cookieChrome.Path,
			Domain:   cookieChrome.Domain,
			Expires:  time.Now().Add(time.Hour * 24 * 180),
			MaxAge:   0,
			Secure:   cookieChrome.Secure,
			HttpOnly: false,
			SameSite: 0,
			Raw:      "",
			Unparsed: []string{},
		}
		cookiesReady = append(cookiesReady, &cookie)
	}

	return cookiesReady, nil
}
