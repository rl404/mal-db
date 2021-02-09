package article

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/model"
)

// Parser is article parser interface.
type Parser interface {
	GetDetails(*goquery.Selection) *model.Article
	GetList(*goquery.Selection) []model.ArticleItem
	GetTags(*goquery.Selection) []model.ArticleTagItem
}

type parser struct {
	cleanImg bool
}

// New to create new news parser.
func New(cleanImg bool) Parser {
	return &parser{
		cleanImg: cleanImg,
	}
}
