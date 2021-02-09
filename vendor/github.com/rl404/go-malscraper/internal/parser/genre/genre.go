package genre

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/model"
)

// Parser is genre parser interface.
type Parser interface {
	GetGenres(*goquery.Selection) []model.ItemCount
}

type parser struct {
	cleanImg bool
}

// New to create new genre parser.
func New(cleanImg bool) Parser {
	return &parser{
		cleanImg: cleanImg,
	}
}
