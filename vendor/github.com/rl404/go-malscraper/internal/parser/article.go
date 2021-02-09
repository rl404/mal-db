package parser

import (
	"net/http"

	"github.com/rl404/go-malscraper/model"
	"github.com/rl404/go-malscraper/pkg/utils"
)

// GetArticle to get featured article detail information.
func (p *Parser) GetArticle(id int) (*model.Article, int, error) {
	doc, code, err := p.getDoc(utils.BuildURL(malURL, "featured", id), ".content-left")
	if err != nil {
		return nil, code, err
	}
	return p.article.GetDetails(doc), http.StatusOK, nil
}

// GetArticles to get featured article list.
func (p *Parser) GetArticles(page int, tag string) ([]model.ArticleItem, int, error) {
	q := map[string]interface{}{"p": page}
	dir := []interface{}{"featured"}
	if tag != "" {
		dir = append(dir, "tag", tag)
	}
	doc, code, err := p.getDoc(utils.BuildURLWithQuery(q, malURL, dir...), ".content-left")
	if err != nil {
		return nil, code, err
	}
	return p.article.GetList(doc), http.StatusOK, nil
}

// GetArticleTag to get featured article tag list.
func (p *Parser) GetArticleTag() ([]model.ArticleTagItem, int, error) {
	doc, code, err := p.getDoc(utils.BuildURL(malURL, "featured", "tag"), ".content-left")
	if err != nil {
		return nil, code, err
	}
	return p.article.GetTags(doc), http.StatusOK, nil
}
