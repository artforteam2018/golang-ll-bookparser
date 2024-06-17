package livelibwebpageiterator

import (
	"context"
	"livelib/components/requestsession"
	"livelib/models/book"
)

func CheckLivelibResponse(id int, cookies *requestsession.StoredCookies) error {
	req := &requestsession.RequestSession{Req: nil, Cli: nil, Res: nil, Cookies: cookies}

	if err := req.PrepareRequestSession(); err != nil {
		return err
	}

	req.PrepareBookRequest(id)
	if err := req.SendRequest(); err != nil {
		return err
	}

	return nil
}

func ParseLivelibId(ctx context.Context, id int) (*book.Book, error) {

	req := &requestsession.RequestSession{}

	if err := req.PrepareRequestSession(); err != nil {
		return nil, err
	}

	req.PrepareBookRequest(id)
	if err := req.SendRequest(); err != nil {
		return nil, err
	}

	if ctx.Err() != nil {
		return nil, ctx.Err()
	}

	if err := req.DecodeRequest(); err != nil {
		return nil, err
	}

	Book := &book.Book{}
	Book.Id = id

	if err := ParseHtmlToBook(req, Book); err != nil {
		return nil, err
	}

	referrer := req.Res.RawResponse.Request.URL.String()
	if err := req.PrepareRequestSession(); err != nil {
		return nil, err
	}

	req.PrepareFullTextInfoRequest(id, referrer)

	if err := req.SendRequest(); err != nil {
		return nil, err
	}

	if ctx.Err() != nil {
		return nil, ctx.Err()
	}

	if err := req.DecodeRequest(); err != nil {
		return nil, err
	}

	if err := ParseFullTextFromRequest(req, Book); err != nil {
		return nil, err
	}

	if ctx.Err() != nil {
		return nil, ctx.Err()
	}

	return Book, nil
}
