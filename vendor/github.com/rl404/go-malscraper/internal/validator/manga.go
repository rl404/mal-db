package validator

import (
	"net/http"

	"github.com/rl404/go-malscraper/errors"
	"github.com/rl404/go-malscraper/internal"
	"github.com/rl404/go-malscraper/model"
)

// GetManga to get manga detail information.
func (v *Validator) GetManga(id int) (*model.Manga, int, error) {
	if id <= 0 {
		return nil, http.StatusBadRequest, errors.ErrInvalidID
	}

	// Check empty id.
	key := internal.GetKey(internal.KeyEmptyManga, id)
	if v.isEmptyID(key) {
		return nil, http.StatusNotFound, errors.ErrNot200
	}

	// Parse.
	data, code, err := v.api.GetManga(id)

	// Save empty id.
	v.saveEmptyID(code, key)

	return data, code, err
}

// GetMangaReview to get manga review list.
func (v *Validator) GetMangaReview(id int, page int) ([]model.Review, int, error) {
	if id <= 0 {
		return nil, http.StatusBadRequest, errors.ErrInvalidID
	}
	if page <= 0 {
		return nil, http.StatusBadRequest, errors.ErrInvalidPage
	}

	// Check empty id.
	key := internal.GetKey(internal.KeyEmptyManga, id)
	if v.isEmptyID(key) {
		return nil, http.StatusNotFound, errors.ErrNot200
	}

	// Parse.
	data, code, err := v.api.GetMangaReview(id, page)

	// Save empty id.
	v.saveEmptyID(code, key)

	return data, code, err
}

// GetMangaRecommendation to get manga recommendation list.
func (v *Validator) GetMangaRecommendation(id int) ([]model.Recommendation, int, error) {
	if id <= 0 {
		return nil, http.StatusBadRequest, errors.ErrInvalidID
	}

	// Check empty id.
	key := internal.GetKey(internal.KeyEmptyManga, id)
	if v.isEmptyID(key) {
		return nil, http.StatusNotFound, errors.ErrNot200
	}

	// Parse.
	data, code, err := v.api.GetMangaRecommendation(id)

	// Save empty id.
	v.saveEmptyID(code, key)

	return data, code, err
}

// GetMangaStats to get manga stats list.
func (v *Validator) GetMangaStats(id int) (*model.Stats, int, error) {
	if id <= 0 {
		return nil, http.StatusBadRequest, errors.ErrInvalidID
	}

	// Check empty id.
	key := internal.GetKey(internal.KeyEmptyManga, id)
	if v.isEmptyID(key) {
		return nil, http.StatusNotFound, errors.ErrNot200
	}

	// Parse.
	data, code, err := v.api.GetMangaStats(id)

	// Save empty id.
	v.saveEmptyID(code, key)

	return data, code, err
}

// GetMangaCharacter to get manga character list.
func (v *Validator) GetMangaCharacter(id int) ([]model.Role, int, error) {
	if id <= 0 {
		return nil, http.StatusBadRequest, errors.ErrInvalidID
	}

	// Check empty id.
	key := internal.GetKey(internal.KeyEmptyManga, id)
	if v.isEmptyID(key) {
		return nil, http.StatusNotFound, errors.ErrNot200
	}

	// Parse.
	data, code, err := v.api.GetMangaCharacter(id)

	// Save empty id.
	v.saveEmptyID(code, key)

	return data, code, err
}

// GetMangaNews to get manga news list.
func (v *Validator) GetMangaNews(id int) ([]model.NewsItem, int, error) {
	if id <= 0 {
		return nil, http.StatusBadRequest, errors.ErrInvalidID
	}

	// Check empty id.
	key := internal.GetKey(internal.KeyEmptyManga, id)
	if v.isEmptyID(key) {
		return nil, http.StatusNotFound, errors.ErrNot200
	}

	// Parse.
	data, code, err := v.api.GetMangaNews(id)

	// Save empty id.
	v.saveEmptyID(code, key)

	return data, code, err
}

// GetMangaArticle to get manga featured article list.
func (v *Validator) GetMangaArticle(id int) ([]model.ArticleItem, int, error) {
	if id <= 0 {
		return nil, http.StatusBadRequest, errors.ErrInvalidID
	}

	// Check empty id.
	key := internal.GetKey(internal.KeyEmptyManga, id)
	if v.isEmptyID(key) {
		return nil, http.StatusNotFound, errors.ErrNot200
	}

	// Parse.
	data, code, err := v.api.GetMangaArticle(id)

	// Save empty id.
	v.saveEmptyID(code, key)

	return data, code, err
}

// GetMangaClub to get manga club list.
func (v *Validator) GetMangaClub(id int) ([]model.ClubItem, int, error) {
	if id <= 0 {
		return nil, http.StatusBadRequest, errors.ErrInvalidID
	}

	// Check empty id.
	key := internal.GetKey(internal.KeyEmptyManga, id)
	if v.isEmptyID(key) {
		return nil, http.StatusNotFound, errors.ErrNot200
	}

	// Parse.
	data, code, err := v.api.GetMangaClub(id)

	// Save empty id.
	v.saveEmptyID(code, key)

	return data, code, err
}

// GetMangaPicture to get manga picture list.
func (v *Validator) GetMangaPicture(id int) ([]string, int, error) {
	if id <= 0 {
		return nil, http.StatusBadRequest, errors.ErrInvalidID
	}

	// Check empty id.
	key := internal.GetKey(internal.KeyEmptyManga, id)
	if v.isEmptyID(key) {
		return nil, http.StatusNotFound, errors.ErrNot200
	}

	// Parse.
	data, code, err := v.api.GetMangaPicture(id)

	// Save empty id.
	v.saveEmptyID(code, key)

	return data, code, err
}

// GetMangaMoreInfo to get manga more info.
func (v *Validator) GetMangaMoreInfo(id int) (string, int, error) {
	if id <= 0 {
		return "", http.StatusBadRequest, errors.ErrInvalidID
	}

	// Check empty id.
	key := internal.GetKey(internal.KeyEmptyManga, id)
	if v.isEmptyID(key) {
		return "", http.StatusNotFound, errors.ErrNot200
	}

	// Parse.
	data, code, err := v.api.GetMangaMoreInfo(id)

	// Save empty id.
	v.saveEmptyID(code, key)

	return data, code, err
}
