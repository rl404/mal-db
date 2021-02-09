package parser

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/errors"
	"github.com/rl404/go-malscraper/model"
)

const malURL string = "https://myanimelist.net"

// Mockable functions.
var httpRequest = http.NewRequest
var parseHTML = goquery.NewDocumentFromReader
var timeSince = time.Since

func (p *Parser) getBody(url string) (io.ReadCloser, int, error) {
	// Prepare request.
	request, err := httpRequest("GET", url, nil)
	if err != nil {
		p.logger.Error("failed preparing request: %s", err.Error())
		return nil, http.StatusInternalServerError, errors.ErrPrepareRequest
	}

	// Do request.
	t := time.Now()
	resp, err := p.http.Do(request)
	if err != nil {
		return nil, http.StatusInternalServerError, errors.ErrHTTPRequest
	}

	// Header check.
	p.logger.Debug("%s %v (%s)", url, resp.StatusCode, timeSince(t).Truncate(time.Microsecond))
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, resp.StatusCode, errors.ErrNot200
	}

	return resp.Body, resp.StatusCode, nil
}

func (p *Parser) getDoc(url string, area string) (*goquery.Selection, int, error) {
	// Get HTML body.
	p.logger.Trace("parsing %s", url)
	body, code, err := p.getBody(url)
	if err != nil {
		return nil, code, err
	}
	defer body.Close()

	// Parse HTML.
	doc, err := parseHTML(body)
	if err != nil {
		p.logger.Error("failed parsing body: %s", err.Error())
		return nil, http.StatusInternalServerError, errors.ErrParseBody
	}

	return doc.Find(area).First(), http.StatusOK, nil
}

func (p *Parser) queryToMap(queryObj model.Query) map[string]interface{} {
	query := make(map[string]interface{})

	for i, c := range []string{"a", "b", "c", "d", "e", "f", "g"} {
		query[fmt.Sprintf("c[%v]", i)] = c
	}

	query["q"] = strings.Replace(queryObj.Title, " ", "+", -1)
	query["show"] = 50 * (queryObj.Page - 1)
	query["type"] = queryObj.Type
	query["score"] = queryObj.Score
	query["status"] = queryObj.Status
	query["r"] = queryObj.Rating
	query["p"] = queryObj.ProducerID
	query["mid"] = queryObj.MagazineID

	if queryObj.FirstLetter != "" {
		query["letter"] = strings.ToUpper(queryObj.FirstLetter)
	}

	excludeGenre := 0
	if queryObj.ExcludeGenre {
		excludeGenre = 1
	}
	query["gx"] = excludeGenre

	if !queryObj.StartDate.IsZero() {
		sYear, sMonth, sDay := queryObj.StartDate.Date()
		query["sd"] = sDay
		query["sm"] = int(sMonth)
		query["sy"] = sYear
	}

	if !queryObj.EndDate.IsZero() {
		eYear, eMonth, eDay := queryObj.EndDate.Date()
		query["ed"] = eDay
		query["em"] = int(eMonth)
		query["ey"] = eYear
	}

	for i, g := range queryObj.GenreIDs {
		query[fmt.Sprintf("genre[%v]", i)] = g
	}

	return query
}
