package producermagazine

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/model"
)

// Parser is producer/magazine parser interface.
type Parser interface {
	GetProducers(*goquery.Selection) []model.ItemCount
	GetAnime(*goquery.Selection) []model.AnimeItem
	GetMagazines(*goquery.Selection) []model.ItemCount
	GetManga(*goquery.Selection) []model.MangaItem
}

type parser struct {
	cleanImg bool
}

// New to create new producer/magazine parser.
func New(cleanImg bool) Parser {
	return &parser{
		cleanImg: cleanImg,
	}
}
