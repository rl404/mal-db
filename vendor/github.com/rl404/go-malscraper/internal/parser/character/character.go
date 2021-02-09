package character

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/model"
)

// Parser is character parser interface.
type Parser interface {
	GetDetails(*goquery.Selection) *model.Character
	GetArticle(*goquery.Selection) []model.ArticleItem
	GetOgraphy(*goquery.Selection, string) []model.Role
	GetVA(*goquery.Selection) []model.Role
	GetPictures(*goquery.Selection) []string
	GetClubs(*goquery.Selection) []model.ClubItem
}

type parser struct {
	cleanImg bool
}

// New to create new character parser.
func New(cleanImg bool) Parser {
	return &parser{
		cleanImg: cleanImg,
	}
}
