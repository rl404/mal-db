package parser

import (
	"net/http"

	"github.com/rl404/go-malscraper/model"
	"github.com/rl404/go-malscraper/pkg/utils"
)

// GetNews to get news detail information.
func (p *Parser) GetNews(id int) (*model.News, int, error) {
	doc, code, err := p.getDoc(utils.BuildURL(malURL, "news", id), ".content-left")
	if err != nil {
		return nil, code, err
	}
	return p.news.GetDetails(doc), http.StatusOK, nil
}

// GetNewsList to get news list.
func (p *Parser) GetNewsList(page int, tag string) ([]model.NewsItem, int, error) {
	q := map[string]interface{}{"p": page}
	dir := []interface{}{"news"}
	if tag != "" {
		dir = append(dir, "tag", tag)
	}
	doc, code, err := p.getDoc(utils.BuildURLWithQuery(q, malURL, dir...), "#content")
	if err != nil {
		return nil, code, err
	}
	return p.news.GetList(doc), http.StatusOK, nil
}

// GetNewsTag to get news tag list.
func (p *Parser) GetNewsTag() (*model.NewsTag, int, error) {
	doc, code, err := p.getDoc(utils.BuildURL(malURL, "news", "tag"), ".content-left")
	if err != nil {
		return nil, code, err
	}
	return p.news.GetTags(doc), http.StatusOK, nil
}
