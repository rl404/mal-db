package validator

import (
	"net/http"

	"github.com/rl404/go-malscraper/errors"
	"github.com/rl404/go-malscraper/internal"
	"github.com/rl404/go-malscraper/model"
)

// GetAnime to get anime details.
func (v *Validator) GetAnime(id int) (*model.Anime, int, error) {
	if id <= 0 {
		return nil, http.StatusBadRequest, errors.ErrInvalidID
	}

	// Check empty id.
	key := internal.GetKey(internal.KeyEmptyAnime, id)
	if v.isEmptyID(key) {
		return nil, http.StatusNotFound, errors.ErrNot200
	}

	// Parse.
	data, code, err := v.api.GetAnime(id)

	// Save empty id.
	v.saveEmptyID(code, key)

	return data, code, err
}

// GetAnimeCharacter to get anime character list.
func (v *Validator) GetAnimeCharacter(id int) ([]model.CharacterItem, int, error) {
	if id <= 0 {
		return nil, http.StatusBadRequest, errors.ErrInvalidID
	}

	// Check empty id.
	key := internal.GetKey(internal.KeyEmptyAnime, id)
	if v.isEmptyID(key) {
		return nil, http.StatusNotFound, errors.ErrNot200
	}

	// Parse.
	data, code, err := v.api.GetAnimeCharacter(id)

	// Save empty id.
	v.saveEmptyID(code, key)

	return data, code, err
}

// GetAnimeStaff to get anime staff list.
func (v *Validator) GetAnimeStaff(id int) ([]model.Role, int, error) {
	if id <= 0 {
		return nil, http.StatusBadRequest, errors.ErrInvalidID
	}

	// Check empty id.
	key := internal.GetKey(internal.KeyEmptyAnime, id)
	if v.isEmptyID(key) {
		return nil, http.StatusNotFound, errors.ErrNot200
	}

	// Parse.
	data, code, err := v.api.GetAnimeStaff(id)

	// Save empty id.
	v.saveEmptyID(code, key)

	return data, code, err
}

// GetAnimeVideo to get anime video list.
func (v *Validator) GetAnimeVideo(id int, page int) (*model.Video, int, error) {
	if id <= 0 {
		return nil, http.StatusBadRequest, errors.ErrInvalidID
	}
	if page <= 0 {
		return nil, http.StatusBadRequest, errors.ErrInvalidPage
	}

	// Check empty id.
	key := internal.GetKey(internal.KeyEmptyAnime, id)
	if v.isEmptyID(key) {
		return nil, http.StatusNotFound, errors.ErrNot200
	}

	// Parse.
	data, code, err := v.api.GetAnimeVideo(id, page)

	// Save empty id.
	v.saveEmptyID(code, key)

	return data, code, err
}

// GetAnimeEpisode to get anime episode list.
func (v *Validator) GetAnimeEpisode(id int, page int) ([]model.Episode, int, error) {
	if id <= 0 {
		return nil, http.StatusBadRequest, errors.ErrInvalidID
	}
	if page <= 0 {
		return nil, http.StatusBadRequest, errors.ErrInvalidPage
	}

	// Check empty id.
	key := internal.GetKey(internal.KeyEmptyAnime, id)
	if v.isEmptyID(key) {
		return nil, http.StatusNotFound, errors.ErrNot200
	}

	// Parse.
	data, code, err := v.api.GetAnimeEpisode(id, page)

	// Save empty id.
	v.saveEmptyID(code, key)

	return data, code, err
}

// GetAnimeStats to get anime stats.
func (v *Validator) GetAnimeStats(id int) (*model.Stats, int, error) {
	if id <= 0 {
		return nil, http.StatusBadRequest, errors.ErrInvalidID
	}

	// Check empty id.
	key := internal.GetKey(internal.KeyEmptyAnime, id)
	if v.isEmptyID(key) {
		return nil, http.StatusNotFound, errors.ErrNot200
	}

	// Parse.
	data, code, err := v.api.GetAnimeStats(id)

	// Save empty id.
	v.saveEmptyID(code, key)

	return data, code, err
}

// GetAnimeReview to get anime review list.
func (v *Validator) GetAnimeReview(id int, page int) ([]model.Review, int, error) {
	if id <= 0 {
		return nil, http.StatusBadRequest, errors.ErrInvalidID
	}
	if page <= 0 {
		return nil, http.StatusBadRequest, errors.ErrInvalidPage
	}

	// Check empty id.
	key := internal.GetKey(internal.KeyEmptyAnime, id)
	if v.isEmptyID(key) {
		return nil, http.StatusNotFound, errors.ErrNot200
	}

	// Parse.
	data, code, err := v.api.GetAnimeReview(id, page)

	// Save empty id.
	v.saveEmptyID(code, key)

	return data, code, err
}

// GetAnimeRecommendation to get anime recommendation list.
func (v *Validator) GetAnimeRecommendation(id int) ([]model.Recommendation, int, error) {
	if id <= 0 {
		return nil, http.StatusBadRequest, errors.ErrInvalidID
	}

	// Check empty id.
	key := internal.GetKey(internal.KeyEmptyAnime, id)
	if v.isEmptyID(key) {
		return nil, http.StatusNotFound, errors.ErrNot200
	}

	// Parse.
	data, code, err := v.api.GetAnimeRecommendation(id)

	// Save empty id.
	v.saveEmptyID(code, key)

	return data, code, err
}

// GetAnimeNews to get anime news list.
func (v *Validator) GetAnimeNews(id int) ([]model.NewsItem, int, error) {
	if id <= 0 {
		return nil, http.StatusBadRequest, errors.ErrInvalidID
	}

	// Check empty id.
	key := internal.GetKey(internal.KeyEmptyAnime, id)
	if v.isEmptyID(key) {
		return nil, http.StatusNotFound, errors.ErrNot200
	}

	// Parse.
	data, code, err := v.api.GetAnimeNews(id)

	// Save empty id.
	v.saveEmptyID(code, key)

	return data, code, err
}

// GetAnimeArticle to get anime featured article list.
func (v *Validator) GetAnimeArticle(id int) ([]model.ArticleItem, int, error) {
	if id <= 0 {
		return nil, http.StatusBadRequest, errors.ErrInvalidID
	}

	// Check empty id.
	key := internal.GetKey(internal.KeyEmptyAnime, id)
	if v.isEmptyID(key) {
		return nil, http.StatusNotFound, errors.ErrNot200
	}

	// Parse.
	data, code, err := v.api.GetAnimeArticle(id)

	// Save empty id.
	v.saveEmptyID(code, key)

	return data, code, err
}

// GetAnimeClub to get anime club list.
func (v *Validator) GetAnimeClub(id int) ([]model.ClubItem, int, error) {
	if id <= 0 {
		return nil, http.StatusBadRequest, errors.ErrInvalidID
	}

	// Check empty id.
	key := internal.GetKey(internal.KeyEmptyAnime, id)
	if v.isEmptyID(key) {
		return nil, http.StatusNotFound, errors.ErrNot200
	}

	// Parse.
	data, code, err := v.api.GetAnimeClub(id)

	// Save empty id.
	v.saveEmptyID(code, key)

	return data, code, err
}

// GetAnimePicture to get anime picture list.
func (v *Validator) GetAnimePicture(id int) ([]string, int, error) {
	if id <= 0 {
		return nil, http.StatusBadRequest, errors.ErrInvalidID
	}

	// Check empty id.
	key := internal.GetKey(internal.KeyEmptyAnime, id)
	if v.isEmptyID(key) {
		return nil, http.StatusNotFound, errors.ErrNot200
	}

	// Parse.
	data, code, err := v.api.GetAnimePicture(id)

	// Save empty id.
	v.saveEmptyID(code, key)

	return data, code, err
}

// GetAnimeMoreInfo to get anime more info.
func (v *Validator) GetAnimeMoreInfo(id int) (string, int, error) {
	if id <= 0 {
		return "", http.StatusBadRequest, errors.ErrInvalidID
	}

	// Check empty id.
	key := internal.GetKey(internal.KeyEmptyAnime, id)
	if v.isEmptyID(key) {
		return "", http.StatusNotFound, errors.ErrNot200
	}

	// Parse.
	data, code, err := v.api.GetAnimeMoreInfo(id)

	// Save empty id.
	v.saveEmptyID(code, key)

	return data, code, err
}
