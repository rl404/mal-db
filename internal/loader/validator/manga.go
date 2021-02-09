package validator

import (
	"net/http"

	"github.com/rl404/mal-db/internal/errors"
	"github.com/rl404/mal-db/internal/model"
)

// GetManga to get manga.
func (v *Validator) GetManga(id int) (*model.Manga, map[string]interface{}, int, error) {
	if id <= 0 {
		return nil, nil, http.StatusBadRequest, errors.ErrInvalidID
	}
	return v.api.GetManga(id)
}

// GetMangaCharacter to get manga character list.
func (v *Validator) GetMangaCharacter(id int, page int, limit int) ([]model.Role, map[string]interface{}, int, error) {
	if id <= 0 {
		return nil, nil, http.StatusBadRequest, errors.ErrInvalidID
	}
	if page <= 0 {
		return nil, nil, http.StatusBadRequest, errors.ErrInvalidPage
	}
	if limit <= 0 {
		return nil, nil, http.StatusBadRequest, errors.ErrInvalidLimit
	}

	// Get data.
	data, meta, code, err := v.api.GetMangaCharacter(id, page, limit)
	if err != nil {
		return nil, nil, code, err
	}

	// Handle pagination.
	start, current := limit*(page-1), len(data)-(page-1)*limit
	if current <= 0 {
		data = []model.Role{}
	} else {
		if current < limit {
			limit = current
		}
		data = data[start : start+limit]
	}

	return data, meta, http.StatusOK, nil
}
