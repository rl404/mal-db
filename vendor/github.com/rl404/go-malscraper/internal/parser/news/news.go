package news

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/model"
)

// Parser is news parser interface.
type Parser interface {
	GetDetails(*goquery.Selection) *model.News
	GetList(*goquery.Selection) []model.NewsItem
	GetTags(*goquery.Selection) *model.NewsTag
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
