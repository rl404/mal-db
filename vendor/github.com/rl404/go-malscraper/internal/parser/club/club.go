package club

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/model"
)

const cdnMyAnimeListURL = "https://cdn.myanimelist.net"

// Parser is club parser interface.
type Parser interface {
	GetDetails(*goquery.Selection) *model.Club
	GetMembers(*goquery.Selection) []model.ClubMember
	GetPictures(*goquery.Selection) []string
	GetRelated(*goquery.Selection) *model.ClubRelated
}

type parser struct {
	cleanImg bool
}

// New to create new club parser.
func New(cleanImg bool) Parser {
	return &parser{
		cleanImg: cleanImg,
	}
}
