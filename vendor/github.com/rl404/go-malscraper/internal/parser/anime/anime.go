package anime

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/model"
)

// Parser is anime parser interface.
type Parser interface {
	GetDetails(*goquery.Selection) *model.Anime
	GetCharacters(*goquery.Selection) []model.CharacterItem
	GetStaff(*goquery.Selection) []model.Role
	GetVideos(*goquery.Selection) *model.Video
	GetEpisodes(*goquery.Selection) []model.Episode
	GetStats(*goquery.Selection) *model.Stats
	GetReviews(*goquery.Selection) []model.Review
	GetRecommendations(*goquery.Selection) []model.Recommendation
	GetNews(*goquery.Selection) []model.NewsItem
	GetArticle(*goquery.Selection) []model.ArticleItem
	GetClubs(*goquery.Selection) []model.ClubItem
	GetPictures(*goquery.Selection) []string
	GetMoreInfo(*goquery.Selection) string
}

type parser struct {
	cleanImg bool
	cleanVid bool
}

// New to create new anime parser.
func New(cleanImg, cleanVid bool) Parser {
	return &parser{
		cleanImg: cleanImg,
		cleanVid: cleanVid,
	}
}
