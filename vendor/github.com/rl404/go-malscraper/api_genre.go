package malscraper

import (
	"github.com/rl404/go-malscraper/model"
)

// GetGenres to get anime/manga genre list.
//
// Param `_type` should be one of these constants.
//
//  AnimeType
//  MangaType
//
// Or just use method `GetAnimeGenres()` or `GetMangaGenres()`.
func (m *Malscraper) GetGenres(_type int) ([]model.ItemCount, int, error) {
	return m.api.GetGenres(mainTypes[_type])
}

// GetAnimeGenres to get anime genre list.
//
// Example: https://myanimelist.net/anime.php.
func (m *Malscraper) GetAnimeGenres() ([]model.ItemCount, int, error) {
	return m.GetGenres(AnimeType)
}

// GetAnimeWithGenre to get anime list with specific genre.
//
// Example: https://myanimelist.net/anime/genre/1/Action.
func (m *Malscraper) GetAnimeWithGenre(id int, page ...int) ([]model.AnimeItem, int, error) {
	p := 1
	if len(page) > 0 {
		p = page[0]
	}
	return m.api.GetAnimeWithGenre(id, p)
}

// GetMangaGenres to get manga genre list.
//
// Example: https://myanimelist.net/manga.php.
func (m *Malscraper) GetMangaGenres() ([]model.ItemCount, int, error) {
	return m.GetGenres(MangaType)
}

// GetMangaWithGenre to get manga list with specific genre.
//
// Example: https://myanimelist.net/manga/genre/1/Action.
func (m *Malscraper) GetMangaWithGenre(id int, page ...int) ([]model.MangaItem, int, error) {
	p := 1
	if len(page) > 0 {
		p = page[0]
	}
	return m.api.GetMangaWithGenre(id, p)
}
