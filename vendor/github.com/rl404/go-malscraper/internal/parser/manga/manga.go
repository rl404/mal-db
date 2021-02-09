package manga

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/model"
)

// Parser is manga parser interface.
type Parser interface {
	GetDetails(*goquery.Selection) *model.Manga
	GetReviews(*goquery.Selection) []model.Review
	GetRecommendations(*goquery.Selection) []model.Recommendation
	GetStats(*goquery.Selection) *model.Stats
	GetCharacters(*goquery.Selection) []model.Role
	GetNews(*goquery.Selection) []model.NewsItem
	GetArticle(*goquery.Selection) []model.ArticleItem
	GetClubs(*goquery.Selection) []model.ClubItem
	GetPictures(*goquery.Selection) []string
	GetMoreInfo(*goquery.Selection) string
}

type parser struct {
	cleanImg bool
}

// New to create new manga parser.
func New(cleanImg bool) Parser {
	return &parser{
		cleanImg: cleanImg,
	}
}
