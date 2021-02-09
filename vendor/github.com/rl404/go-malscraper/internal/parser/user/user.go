package user

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/model"
)

// Parser is user parser interface.
type Parser interface {
	GetDetails(*goquery.Selection) *model.User
	GetStats(*goquery.Selection) *model.UserStats
	GetFavorites(*goquery.Selection) *model.UserFavorite
	GetFriends(*goquery.Selection) []model.UserFriend
	GetHistory(*goquery.Selection) []model.UserHistory
	GetReviews(*goquery.Selection) []model.Review
	GetRecommendations(*goquery.Selection) []model.Recommendation
	GetClubs(*goquery.Selection) []model.Item
	GetAnime([]model.UserRawAnime) []model.UserAnime
	GetManga([]model.UserRawManga) []model.UserManga
}

type parser struct {
	cleanImg bool
}

// New to create new user parser.
func New(cleanImg bool) Parser {
	return &parser{
		cleanImg: cleanImg,
	}
}
