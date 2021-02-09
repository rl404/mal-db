package top

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/model"
)

// Parser is top parser interface.
type Parser interface {
	GetAnime(*goquery.Selection) []model.TopAnime
	GetManga(*goquery.Selection) []model.TopManga
	GetCharacter(*goquery.Selection) []model.TopCharacter
	GetPeople(*goquery.Selection) []model.TopPeople
}

type parser struct {
	cleanImg bool
}

// New to create new top parser.
func New(cleanImg bool) Parser {
	return &parser{
		cleanImg: cleanImg,
	}
}
