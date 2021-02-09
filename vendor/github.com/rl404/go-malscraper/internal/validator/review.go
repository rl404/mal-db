package validator

import (
	"net/http"

	"github.com/rl404/go-malscraper/errors"
	"github.com/rl404/go-malscraper/internal"
	"github.com/rl404/go-malscraper/model"
)

// GetReview to get review detail information.
func (v *Validator) GetReview(id int) (*model.Review, int, error) {
	if id <= 0 {
		return nil, http.StatusBadRequest, errors.ErrInvalidID
	}

	// Check empty id.
	key := internal.GetKey(internal.KeyEmptyReview, id)
	if v.isEmptyID(key) {
		return nil, http.StatusNotFound, errors.ErrNot200
	}

	// Parse.
	data, code, err := v.api.GetReview(id)

	// Save empty id.
	v.saveEmptyID(code, key)

	return data, code, err
}

// GetReviews to get anime/manga/best review list.
func (v *Validator) GetReviews(t string, page int) ([]model.Review, int, error) {
	if t != AnimeType && t != MangaType && t != BestReview {
		return nil, http.StatusBadRequest, errors.ErrInvalidType
	}
	if page <= 0 {
		return nil, http.StatusBadRequest, errors.ErrInvalidPage
	}
	return v.api.GetReviews(t, page)
}
