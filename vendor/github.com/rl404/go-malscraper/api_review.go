package malscraper

import (
	"github.com/rl404/go-malscraper/model"
)

// GetReview to get review detail information.
//
// Example: https://myanimelist.net/reviews.php?id=1.
func (m *Malscraper) GetReview(id int) (*model.Review, int, error) {
	return m.api.GetReview(id)
}

// GetReviews to get anime/manga/best review list.
//
// Param `_type` should be one of these constants.
//
//  AnimeReview
//  MangaReview
//  BestReview
//
// Or just use method `GetAnimeReviews()`, `GetMangaReviews()` or `GetBestReviews()`.
func (m *Malscraper) GetReviews(_type int, page ...int) ([]model.Review, int, error) {
	p := 1
	if len(page) > 0 {
		p = page[0]
	}
	return m.api.GetReviews(reviewStr[_type], p)
}

// GetAnimeReviews to get anime review list.
//
// Example: https://myanimelist.net/reviews.php?t=anime.
func (m *Malscraper) GetAnimeReviews(page ...int) ([]model.Review, int, error) {
	return m.GetReviews(AnimeReview, page...)
}

// GetMangaReviews to get manga review list.
//
// Example: https://myanimelist.net/reviews.php?t=manga.
func (m *Malscraper) GetMangaReviews(page ...int) ([]model.Review, int, error) {
	return m.GetReviews(MangaReview, page...)
}

// GetBestReviews to get best anime & manga review list.
//
// Example: https://myanimelist.net/reviews.php?st=bestvoted.
func (m *Malscraper) GetBestReviews(page ...int) ([]model.Review, int, error) {
	return m.GetReviews(BestReview, page...)
}
