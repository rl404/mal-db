package validator

import (
	"net/http"

	"github.com/rl404/mal-db/internal/constant"
	"github.com/rl404/mal-db/internal/errors"
	"github.com/rl404/mal-db/internal/model"
)

// GetCharacter to get character.
func (v *Validator) GetCharacter(id int) (*model.Character, map[string]interface{}, int, error) {
	if id <= 0 {
		return nil, nil, http.StatusBadRequest, errors.ErrInvalidID
	}
	return v.api.GetCharacter(id)
}

// GetCharacterOgraphy to get character anime/manga ography list.
func (v *Validator) GetCharacterOgraphy(id int, t string, page int, limit int) ([]model.Role, map[string]interface{}, int, error) {
	if id <= 0 {
		return nil, nil, http.StatusBadRequest, errors.ErrInvalidID
	}
	if t != constant.AnimeType && t != constant.MangaType {
		return nil, nil, http.StatusBadRequest, errors.ErrInvalidType
	}
	if page <= 0 {
		return nil, nil, http.StatusBadRequest, errors.ErrInvalidPage
	}
	if limit <= 0 {
		return nil, nil, http.StatusBadRequest, errors.ErrInvalidLimit
	}

	// Get data.
	data, meta, code, err := v.api.GetCharacterOgraphy(id, t, page, limit)
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

// GetCharacterVA to get character voice actor list.
func (v *Validator) GetCharacterVA(id int, page int, limit int) ([]model.Role, map[string]interface{}, int, error) {
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
	data, meta, code, err := v.api.GetCharacterVA(id, page, limit)
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
