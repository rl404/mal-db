package validator

import (
	"net/http"

	"github.com/rl404/go-malscraper/errors"
	"github.com/rl404/go-malscraper/model"
)

// GetGenres to get anime/manga genre list.
func (v *Validator) GetGenres(t string) ([]model.ItemCount, int, error) {
	if t != AnimeType && t != MangaType {
		return nil, http.StatusBadRequest, errors.ErrInvalidType
	}
	return v.api.GetGenres(t)
}

// GetAnimeWithGenre to get anime list with specific genre.
func (v *Validator) GetAnimeWithGenre(id int, page int) ([]model.AnimeItem, int, error) {
	if page <= 0 {
		return nil, http.StatusBadRequest, errors.ErrInvalidPage
	}
	if !v.isAnimeGenreValid(id) {
		return nil, http.StatusBadRequest, errors.ErrInvalidID
	}
	return v.api.GetAnimeWithGenre(id, page)
}

// GetMangaWithGenre to get manga list with specific genre.
func (v *Validator) GetMangaWithGenre(id int, page int) ([]model.MangaItem, int, error) {
	if page <= 0 {
		return nil, http.StatusBadRequest, errors.ErrInvalidPage
	}
	if !v.isMangaGenreValid(id) {
		return nil, http.StatusBadRequest, errors.ErrInvalidID
	}
	return v.api.GetMangaWithGenre(id, page)
}
