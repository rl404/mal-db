package search

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/model"
)

// Parser is search parser interface.
type Parser interface {
	GetAnime(*goquery.Selection) []model.AnimeSearch
	GetManga(*goquery.Selection) []model.MangaSearch
	GetCharacter(*goquery.Selection) []model.CharacterSearch
	GetPeople(*goquery.Selection) []model.PeopleSearch
	GetClub(*goquery.Selection) []model.ClubSearch
	GetUser(*goquery.Selection) []model.UserSearch
}

type parser struct {
	cleanImg bool
}

// New to create new search parser.
func New(cleanImg bool) Parser {
	return &parser{
		cleanImg: cleanImg,
	}
}
