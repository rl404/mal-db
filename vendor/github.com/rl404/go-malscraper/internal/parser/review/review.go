package review

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/model"
)

// Parser is review parser interface.
type Parser interface {
	GetDetails(*goquery.Selection) *model.Review
	GetReviews(*goquery.Selection) []model.Review
}

type parser struct {
	cleanImg bool
}

// New to create new review parser.
func New(cleanImg bool) Parser {
	return &parser{
		cleanImg: cleanImg,
	}
}
