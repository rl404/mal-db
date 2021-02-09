package malscraper

import (
	"github.com/rl404/go-malscraper/model"
)

// GetManga to get manga detail information.
//
// Example: https://myanimelist.net/manga/1.
func (m *Malscraper) GetManga(id int) (*model.Manga, int, error) {
	return m.api.GetManga(id)
}

// GetMangaReview to get manga review list.
//
// Example: https://myanimelist.net/manga/1/Monster/reviews.
func (m *Malscraper) GetMangaReview(id int, page ...int) ([]model.Review, int, error) {
	p := 1
	if len(page) > 0 {
		p = page[0]
	}
	return m.api.GetMangaReview(id, p)
}

// GetMangaRecommendation to get manga recommendation list.
//
// Example: https://myanimelist.net/manga/1/Monster/userrecs.
func (m *Malscraper) GetMangaRecommendation(id int) ([]model.Recommendation, int, error) {
	return m.api.GetMangaRecommendation(id)
}

// GetMangaStats to get manga stats list.
//
// Example: https://myanimelist.net/manga/1/Monster/stats.
func (m *Malscraper) GetMangaStats(id int) (*model.Stats, int, error) {
	return m.api.GetMangaStats(id)
}

// GetMangaCharacter to get manga character list.
//
// Example: https://myanimelist.net/manga/1/Monster/characters.
func (m *Malscraper) GetMangaCharacter(id int) ([]model.Role, int, error) {
	return m.api.GetMangaCharacter(id)
}

// GetMangaNews to get manga news list.
//
// Example: https://myanimelist.net/manga/1/Monster/news.
func (m *Malscraper) GetMangaNews(id int) ([]model.NewsItem, int, error) {
	return m.api.GetMangaNews(id)
}

// GetMangaArticle to get manga featured article list.
//
// Example: https://myanimelist.net/manga/1/Monster/featured.
func (m *Malscraper) GetMangaArticle(id int) ([]model.ArticleItem, int, error) {
	return m.api.GetMangaArticle(id)
}

// GetMangaClub to get manga club list.
//
// Example: https://myanimelist.net/manga/1/Monster/clubs.
func (m *Malscraper) GetMangaClub(id int) ([]model.ClubItem, int, error) {
	return m.api.GetMangaClub(id)
}

// GetMangaPicture to get manga picture list.
//
// Example: https://myanimelist.net/manga/1/Monster/pics.
func (m *Malscraper) GetMangaPicture(id int) ([]string, int, error) {
	return m.api.GetMangaPicture(id)
}

// GetMangaMoreInfo to get manga more info.
//
// Example: https://myanimelist.net/manga/2/Berserk/moreinfo.
func (m *Malscraper) GetMangaMoreInfo(id int) (string, int, error) {
	return m.api.GetMangaMoreInfo(id)
}
