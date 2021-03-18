package validator

import (
	"net/http"

	"github.com/rl404/mal-db/internal/errors"
	"github.com/rl404/mal-db/internal/model"
)

// GetAnime to get anime.
func (v *Validator) GetAnime(id int) (*model.Anime, map[string]interface{}, int, error) {
	if id <= 0 {
		return nil, nil, http.StatusBadRequest, errors.ErrInvalidID
	}
	return v.api.GetAnime(id)
}

// GetAnimeCharacter to get anime character list.
func (v *Validator) GetAnimeCharacter(id int, page int, limit int) ([]model.AnimeCharacter, map[string]interface{}, int, error) {
	if id <= 0 {
		return nil, nil, http.StatusBadRequest, errors.ErrInvalidID
	}
	if page <= 0 {
		return nil, nil, http.StatusBadRequest, errors.ErrInvalidPage
	}
	if limit <= 0 {
		return nil, nil, http.StatusBadRequest, errors.ErrInvalidLimit
	}
	return v.api.GetAnimeCharacter(id, page, limit)
}

// GetAnimeStaff to get anime staff list.
func (v *Validator) GetAnimeStaff(id int, page int, limit int) ([]model.Role, map[string]interface{}, int, error) {
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
	data, meta, code, err := v.api.GetAnimeStaff(id, page, limit)
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
