package validator

import (
	"net/http"

	"github.com/rl404/go-malscraper/errors"
	"github.com/rl404/go-malscraper/internal"
	"github.com/rl404/go-malscraper/model"
)

// GetClubs to get club list.
func (v *Validator) GetClubs(page int) ([]model.ClubSearch, int, error) {
	if page <= 0 {
		return nil, http.StatusBadRequest, errors.ErrInvalidPage
	}
	return v.api.GetClubs(page)
}

// GetClub to get club detail information.
func (v *Validator) GetClub(id int) (*model.Club, int, error) {
	if id <= 0 {
		return nil, http.StatusBadRequest, errors.ErrInvalidID
	}

	// Check empty id.
	key := internal.GetKey(internal.KeyEmptyClub, id)
	if v.isEmptyID(key) {
		return nil, http.StatusNotFound, errors.ErrNot200
	}

	// Parse.
	data, code, err := v.api.GetClub(id)

	// Save empty id.
	v.saveEmptyID(code, key)

	return data, code, err
}

// GetClubMember to get club member list.
func (v *Validator) GetClubMember(id int, page int) ([]model.ClubMember, int, error) {
	if id <= 0 {
		return nil, http.StatusBadRequest, errors.ErrInvalidID
	}
	if page <= 0 {
		return nil, http.StatusBadRequest, errors.ErrInvalidPage
	}

	// Check empty id.
	key := internal.GetKey(internal.KeyEmptyClub, id)
	if v.isEmptyID(key) {
		return nil, http.StatusNotFound, errors.ErrNot200
	}

	// Parse.
	data, code, err := v.api.GetClubMember(id, page)

	// Save empty id.
	v.saveEmptyID(code, key)

	return data, code, err
}

// GetClubPicture to get club picture list.
func (v *Validator) GetClubPicture(id int) ([]string, int, error) {
	if id <= 0 {
		return nil, http.StatusBadRequest, errors.ErrInvalidID
	}

	// Check empty id.
	key := internal.GetKey(internal.KeyEmptyClub, id)
	if v.isEmptyID(key) {
		return nil, http.StatusNotFound, errors.ErrNot200
	}

	// Parse.
	data, code, err := v.api.GetClubPicture(id)

	// Save empty id.
	v.saveEmptyID(code, key)

	return data, code, err
}

// GetClubRelated to get club related list.
func (v *Validator) GetClubRelated(id int) (*model.ClubRelated, int, error) {
	if id <= 0 {
		return nil, http.StatusBadRequest, errors.ErrInvalidID
	}

	// Check empty id.
	key := internal.GetKey(internal.KeyEmptyClub, id)
	if v.isEmptyID(key) {
		return nil, http.StatusNotFound, errors.ErrNot200
	}

	// Parse.
	data, code, err := v.api.GetClubRelated(id)

	// Save empty id.
	v.saveEmptyID(code, key)

	return data, code, err
}
