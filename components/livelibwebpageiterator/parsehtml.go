package livelibwebpageiterator

import (
	"fmt"
	"livelib/components/requestsession"
	"livelib/models/book"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var regularSearch, _ = regexp.Compile(`(\d*) стр`)

type FullTextAnswer struct {
	Error_code int    `json:"error_code"`
	Content    string `json:"content"`
}

func ParseFullTextFromRequest(req *requestsession.RequestSession, Book *book.Book) error {
	Answer := FullTextAnswer{}
	if err := req.Res.JSON(&Answer); err != nil {
		return fmt.Errorf("error json unmarshalling full text: %s", err)
	}

	Book.Description = Answer.Content
	return nil
}

func ParseHtmlToBook(req *requestsession.RequestSession, Book *book.Book) error {

	doc, err := goquery.NewDocumentFromReader(req.Res.RawResponse.Body)

	zoza := req.Res.RawResponse.Request.URL.String()
	_, textId, _ := strings.Cut(zoza, "-")

	Book.TextId = textId

	if err != nil {
		return fmt.Errorf("parsing to goquery error: %s", err)
	}

	if len(doc.Has(".page-404").Nodes) > 0 {
		return fmt.Errorf("captcha error")
	}

	//parsing request

	doc.Find(".bc__book-title").Each(func(i int, q *goquery.Selection) {
		Book.Title = q.Text()
	})

	doc.Find(".bc-author").Each(func(i int, q *goquery.Selection) {
		q.Find(".bc-author__link").Each(func(i int, q *goquery.Selection) {
			href, _ := q.Attr("href")
			Book.Authors = append(Book.Authors, book.Author{Name: q.Text(), Link: href})
		})
	})

	doc.Find(".bc-edition").Each(func(i int, q *goquery.Selection) {
		q.Find("tr").Each(func(i int, q *goquery.Selection) {
			q.Find("td").Each(func(i int, q *goquery.Selection) {
				q.Find("a").Each(func(i int, q *goquery.Selection) {
					href, _ := q.Attr("href")
					editionInfo := strings.Split(href, "/")
					edType, edVal := editionInfo[1], editionInfo[2]
					Book.Editions = append(Book.Editions, book.Info{Type: edType, Value: edVal})
				})
			})
		})
	})

	doc.Find(".bc-menu__image").Each(func(i int, q *goquery.Selection) {
		src, _ := q.Attr("src")
		Book.Image = src
	})

	doc.Find(".bc-info__wrapper>label").Each(func(i int, q *goquery.Selection) {
		forField, _ := q.Attr("for")

		// ISBN and more
		if forField == "5" {
			q.Next().Children().Each(func(i int, q *goquery.Selection) {

				text := strings.Trim(q.Text(), "\n ")
				fieldsInfo := strings.Split(text, ":")
				for i, s := range fieldsInfo {
					fieldsInfo[i] = strings.Trim(s, "\n ")
				}
				if len(fieldsInfo) == 2 && strings.Count(text, "\n") < 1 {
					fieldsKey, fieldsVal := fieldsInfo[0], fieldsInfo[1]
					if fieldsInfo[0] == "ISBN" {
						fieldsInfo[1], _, _ = strings.Cut(fieldsInfo[1], ",")
					}
					Book.Infos = append(Book.Infos, book.Info{Type: fieldsKey, Value: fieldsVal})
				} else {
					dirtyInfo := strings.Join(fieldsInfo, ";")
					dirtyField := strings.ToLower(strings.Trim(dirtyInfo, "\n "))

					var bookInfo book.Info

					if strings.Contains(dirtyField, "супер") {
						bookInfo = book.Info{Type: "Обложка", Value: "Суперобложка"}
					}
					if strings.Contains(dirtyField, "мягк") {
						bookInfo = book.Info{Type: "Обложка", Value: "Мягкий переплет"}
					}
					if strings.Contains(dirtyField, "тверд") || strings.Contains(dirtyField, "твёрд") {
						bookInfo = book.Info{Type: "Обложка", Value: "Твердый переплет"}
					}

					if bookInfo.Type != "" {
						Book.Infos = append(Book.Infos, bookInfo)
					}

					if strings.Contains(dirtyField, "стр") {

						matched := regularSearch.FindStringSubmatch(dirtyField)
						if len(matched) > 1 {
							bookInfo = book.Info{Type: "Кол-во страниц", Value: matched[1]}
						}
					}

					if bookInfo.Type != "" {
						Book.Infos = append(Book.Infos, bookInfo)
					}
				}
			})
		}
		if forField == "6" {
			q.Next().Children().Each(func(i int, q *goquery.Selection) {
				q.Find("a").Each(func(i int, q *goquery.Selection) {
					href, _ := q.Attr("href")
					genresInfo := strings.Split(href, "/")
					if len(genresInfo) > 1 && genresInfo[1] == "genre" {
						Book.Genres = append(Book.Genres, book.Genre{Name: strings.Replace(genresInfo[2], "-", " ", -1), Link: href})
					}
				})

			})
		}
	})
	return nil
}
