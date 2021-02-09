package validator

import (
	"net/http"

	"github.com/rl404/go-malscraper/errors"
	"github.com/rl404/go-malscraper/model"
	"github.com/rl404/go-malscraper/pkg/utils"
)

// GetTopAnime to get top anime list.
func (v *Validator) GetTopAnime(t int, p int) ([]model.TopAnime, int, error) {
	if !utils.InArrayInt(topAnimeTypes, t) {
		return nil, http.StatusBadRequest, errors.ErrInvalidType
	}
	if p <= 0 {
		return nil, http.StatusBadRequest, errors.ErrInvalidPage
	}
	return v.api.GetTopAnime(t, p)
}

// GetTopManga to get top manga list.
func (v *Validator) GetTopManga(t int, p int) ([]model.TopManga, int, error) {
	if !utils.InArrayInt(topMangaTypes, t) {
		return nil, http.StatusBadRequest, errors.ErrInvalidType
	}
	if p <= 0 {
		return nil, http.StatusBadRequest, errors.ErrInvalidPage
	}
	return v.api.GetTopManga(t, p)
}

// GetTopCharacter to get top character list.
func (v *Validator) GetTopCharacter(page int) ([]model.TopCharacter, int, error) {
	if page <= 0 {
		return nil, http.StatusBadRequest, errors.ErrInvalidPage
	}
	return v.api.GetTopCharacter(page)
}

// GetTopPeople to get top people list.
func (v *Validator) GetTopPeople(page int) ([]model.TopPeople, int, error) {
	if page <= 0 {
		return nil, http.StatusBadRequest, errors.ErrInvalidPage
	}
	return v.api.GetTopPeople(page)
}
