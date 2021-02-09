package malscraper

import (
	"github.com/rl404/go-malscraper/model"
)

// GetUser to get user detail information.
//
// Example: https://myanimelist.net/profile/rl404.
func (m *Malscraper) GetUser(username string) (*model.User, int, error) {
	return m.api.GetUser(username)
}

// GetUserStats to get user stats detail information.
//
// Example: https://myanimelist.net/profile/rl404.
func (m *Malscraper) GetUserStats(username string) (*model.UserStats, int, error) {
	return m.api.GetUserStats(username)
}

// GetUserFavorite to get user favorite list.
//
// Example: https://myanimelist.net/profile/rl404.
func (m *Malscraper) GetUserFavorite(username string) (*model.UserFavorite, int, error) {
	return m.api.GetUserFavorite(username)
}

// GetUserFriend to get user friend list.
//
// Example: https://myanimelist.net/profile/rl404/friends.
func (m *Malscraper) GetUserFriend(username string, page ...int) ([]model.UserFriend, int, error) {
	p := 1
	if len(page) > 0 {
		p = page[0]
	}
	return m.api.GetUserFriend(username, p)
}

// GetUserHistory to get user history list.
//
// Type should be one of these constants.
//
//  AllType
//  AnimeType
//  MangaType
//
// Example: https://myanimelist.net/history/rl404.
func (m *Malscraper) GetUserHistory(username string, _type ...int) ([]model.UserHistory, int, error) {
	t := 0
	if len(_type) > 0 {
		t = _type[0]
	}
	return m.api.GetUserHistory(username, mainTypes[t])
}

// GetUserReview to get user review list.
//
// Example: https://myanimelist.net/profile/Archaeon/reviews.
func (m *Malscraper) GetUserReview(username string, page ...int) ([]model.Review, int, error) {
	p := 1
	if len(page) > 0 {
		p = page[0]
	}
	return m.api.GetUserReview(username, p)
}

// GetUserRecommendation to get user recommendation list.
//
// Example: https://myanimelist.net/profile/Archaeon/recommendations.
func (m *Malscraper) GetUserRecommendation(username string, page ...int) ([]model.Recommendation, int, error) {
	p := 1
	if len(page) > 0 {
		p = page[0]
	}
	return m.api.GetUserRecommendation(username, p)
}

// GetUserClub to get user club list.
//
// Example: https://myanimelist.net/profile/Archaeon/clubs.
func (m *Malscraper) GetUserClub(username string) ([]model.Item, int, error) {
	return m.api.GetUserClub(username)
}

// GetUserAnime to quick get user anime list.
//
// Example: https://myanimelist.net/animelist/rl404.
func (m *Malscraper) GetUserAnime(username string, page ...int) ([]model.UserAnime, int, error) {
	p := 1
	if len(page) > 0 {
		p = page[0]
	}
	return m.GetUserAnimeAdv(model.UserListQuery{Username: username, Status: StatusAll, Page: p})
}

// GetUserAnimeAdv to get user anime list.
//
// Available constant options.
//
//  Status            Order
//  ---------------   -------------------
//  StatusDefault     OrderDefault
//  StatusCurrent     OrderAnimeTitle
//  StatusCompleted   OrderAnimeFinishDate
//  StatusOnHold      OrderAnimeStartDate
//  StatusDropped     OrderAnimeScore
//  StatusPlanned     OrderAnimeType
//  StatusAll         OrderAnimeRated
//                    OrderAnimePriority
//                    OrderAnimeProgress
//                    OrderAnimeStorage
//                    OrderAnimeAirStart
//                    OrderAnimeAirEnd
//
// Example: https://myanimelist.net/animelist/rl404.
func (m *Malscraper) GetUserAnimeAdv(query model.UserListQuery) ([]model.UserAnime, int, error) {
	if query.Page == 0 {
		query.Page = 1
	}
	return m.api.GetUserAnime(query)
}

// GetUserManga to quick get user manga list.
//
// Example: https://myanimelist.net/mangalist/rl404.
func (m *Malscraper) GetUserManga(username string, page ...int) ([]model.UserManga, int, error) {
	p := 1
	if len(page) > 0 {
		p = page[0]
	}
	return m.GetUserMangaAdv(model.UserListQuery{Username: username, Status: StatusAll, Page: p})
}

// GetUserMangaAdv to get user manga list.
//
// Available constant options.
//
//  Status            Order
//  ---------------   -------------------
//  StatusDefault     OrderDefault
//  StatusCurrent     OrderMangaTitle
//  StatusCompleted   OrderMangaFinishDate
//  StatusOnHold      OrderMangaStartDate
//  StatusDropped     OrderMangaScore
//  StatusPlanned     OrderMangaPriority
//  StatusAll         OrderMangaChapter
//                    OrderMangaVolume
//                    OrderMangaType
//                    OrderMangaPublishStart
//                    OrderMangaPublishEnd
//
// Example: https://myanimelist.net/mangalist/rl404.
func (m *Malscraper) GetUserMangaAdv(query model.UserListQuery) ([]model.UserManga, int, error) {
	if query.Page == 0 {
		query.Page = 1
	}
	return m.api.GetUserManga(query)
}
