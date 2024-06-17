package requestsession

import (
	"compress/gzip"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"gopkg.in/h2non/gentleman.v2"
	"gopkg.in/h2non/gentleman.v2/plugins/multipart"
	"gopkg.in/h2non/gentleman.v2/plugins/redirect"
)

type RequestSession struct {
	Req     *gentleman.Request
	Cli     *gentleman.Client
	Res     *gentleman.Response
	Cookies *StoredCookies
}

const livelibURL = "https://www.livelib.ru"

func (req *RequestSession) PrepareRequestSession() error {
	cli := gentleman.New()
	cli.URL(livelibURL)

	request := cli.Request()

	req.Cli = cli
	req.Req = request
	if err := AddCookiesToRequest(req); err != nil {
		return fmt.Errorf("cookies restore error: %s", err)
	}

	return nil
}

func (req *RequestSession) SendRequest() error {
	res, err := req.Req.Send()
	req.Res = res

	if err != nil {
		return err
	}

	if res.RawResponse.Request.URL.Path == "/service/ratelimitcaptcha" {
		panic("got recaptcha")
	}

	if !res.Ok {
		return fmt.Errorf("invalid server response: %d", res.StatusCode)
	}

	return nil
}

func (req *RequestSession) PrepareBookRequest(id int) {
	req.Req.Path("/book/" + strconv.Itoa(id))
	setHeaders(req.Req)
	req.Req.Method("GET")
}

func (req *RequestSession) PrepareFullTextInfoRequest(id int, referrer string) {
	data := url.Values{
		"object_alias":  multipart.Values{"edition"},
		"object_id":     multipart.Values{strconv.Itoa(id)},
		"is_new_design": multipart.Values{"ll2019"},
	}
	req.Req.Path("/feed/getfullobjecttext")
	setHeadersForFullText(req.Req, referrer)
	req.Req.Body(strings.NewReader(data.Encode()))
	req.Req.Method("POST")
	req.Cli.Use(redirect.Limit(0))
}

func (req *RequestSession) DecodeRequest() error {
	if req.Res.Header.Get("Content-Encoding") == "gzip" {
		req.Res.Header.Del("Content-Length")
		zr, err := gzip.NewReader(req.Res.RawResponse.Body)
		if err != nil {
			return fmt.Errorf("invalid decoding: %d", req.Res.StatusCode)
		}
		req.Res.RawResponse.Body = zr
	}
	return nil
}
