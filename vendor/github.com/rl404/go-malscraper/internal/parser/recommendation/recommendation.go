package recommendation

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/model"
)

// Parser is recommendation parser interface.
type Parser interface {
	GetDetails(*goquery.Selection) *model.Recommendation
	GetRecommendations(*goquery.Selection) []model.Recommendation
}

type parser struct {
	cleanImg bool
}

// New to create new recommendation parser.
func New(cleanImg bool) Parser {
	return &parser{
		cleanImg: cleanImg,
	}
}
