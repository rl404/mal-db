package people

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/model"
)

// Parser is people parser interface.
type Parser interface {
	GetDetails(*goquery.Selection) *model.People
	GetCharacters(*goquery.Selection) []model.PeopleCharacter
	GetStaffManga(*goquery.Selection, string) []model.Role
	GetNews(*goquery.Selection) []model.NewsItem
	GetArticle(*goquery.Selection) []model.ArticleItem
	GetPictures(*goquery.Selection) []string
}

type parser struct {
	cleanImg bool
}

// New to create new people parser.
func New(cleanImg bool) Parser {
	return &parser{
		cleanImg: cleanImg,
	}
}
