package validator

import (
	"net/http"

	"github.com/rl404/mal-db/internal/errors"
	"github.com/rl404/mal-db/internal/model"
)

// GetPeople to get people.
func (v *Validator) GetPeople(id int) (*model.People, map[string]interface{}, int, error) {
	if id <= 0 {
		return nil, nil, http.StatusBadRequest, errors.ErrInvalidID
	}
	return v.api.GetPeople(id)
}

// GetPeopleVA to get people voice actor role list.
func (v *Validator) GetPeopleVA(id int, page int, limit int) ([]model.VoiceActor, map[string]interface{}, int, error) {
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
	data, meta, code, err := v.api.GetPeopleVA(id, page, limit)
	if err != nil {
		return nil, nil, code, err
	}

	// Handle pagination.
	start, current := limit*(page-1), len(data)-(page-1)*limit
	if current <= 0 {
		data = []model.VoiceActor{}
	} else {
		if current < limit {
			limit = current
		}
		data = data[start : start+limit]
	}

	return data, meta, http.StatusOK, nil
}

// GetPeopleStaff to get people anime staff role list.
func (v *Validator) GetPeopleStaff(id int, page int, limit int) ([]model.Role, map[string]interface{}, int, error) {
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
	data, meta, code, err := v.api.GetPeopleStaff(id, page, limit)
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

// GetPeopleManga to get people published manga list.
func (v *Validator) GetPeopleManga(id int, page int, limit int) ([]model.Role, map[string]interface{}, int, error) {
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
	data, meta, code, err := v.api.GetPeopleManga(id, page, limit)
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
