package malscraper

import (
	"github.com/rl404/go-malscraper/model"
)

// GetRecommendation to get anime/manga recommendation detail information.
//
// Param `_type` should be one of these constants.
//
//  AnimeType
//  MangaType
//
// Or just use method `GetRecommendationAnime()` or `GetRecommendationManga()`.
func (m *Malscraper) GetRecommendation(_type int, id1, id2 int) (*model.Recommendation, int, error) {
	return m.api.GetRecommendation(mainTypes[_type], id1, id2)
}

// GetRecommendationAnime to get anime recommendation.
//
// Example: https://myanimelist.net/recommendations/anime/1-205.
func (m *Malscraper) GetRecommendationAnime(id1, id2 int) (*model.Recommendation, int, error) {
	return m.GetRecommendation(AnimeType, id1, id2)
}

// GetRecommendationManga to get manga recommendation.
//
// Example: https://myanimelist.net/recommendations/manga/1-21.
func (m *Malscraper) GetRecommendationManga(id1, id2 int) (*model.Recommendation, int, error) {
	return m.GetRecommendation(MangaType, id1, id2)
}

// GetRecommendations to get anime/manga recommendation list.
//
// Param `_type` should be one of these constants.
//
//  AnimeType
//  MangaType
//
// Or just use method `GetAnimeRecommendations()` or `GetMangaRecommendations()`.
func (m *Malscraper) GetRecommendations(_type int, page ...int) ([]model.Recommendation, int, error) {
	p := 1
	if len(page) > 0 {
		p = page[0]
	}
	return m.api.GetRecommendations(mainTypes[_type], p)
}

// GetAnimeRecommendations to get anime recommendation list.
//
// Example: https://myanimelist.net/recommendations.php?s=recentrecs&t=anime.
func (m *Malscraper) GetAnimeRecommendations(page ...int) ([]model.Recommendation, int, error) {
	return m.GetRecommendations(AnimeType, page...)
}

// GetMangaRecommendations to get manga recommendation list.
//
// Example: https://myanimelist.net/recommendations.php?s=recentrecs&t=manga.
func (m *Malscraper) GetMangaRecommendations(page ...int) ([]model.Recommendation, int, error) {
	return m.GetRecommendations(MangaType, page...)
}
