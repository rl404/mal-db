package validator

import (
	"net/http"

	"github.com/rl404/go-malscraper/errors"
	"github.com/rl404/go-malscraper/internal"
	"github.com/rl404/go-malscraper/model"
)

// GetPeople to get people detail information.
func (v *Validator) GetPeople(id int) (*model.People, int, error) {
	if id <= 0 {
		return nil, http.StatusBadRequest, errors.ErrInvalidID
	}

	// Check empty id.
	key := internal.GetKey(internal.KeyEmptyPeople, id)
	if v.isEmptyID(key) {
		return nil, http.StatusNotFound, errors.ErrNot200
	}

	// Parse.
	data, code, err := v.api.GetPeople(id)

	// Save empty id.
	v.saveEmptyID(code, key)

	return data, code, err
}

// GetPeopleCharacter to get people anime character list.
func (v *Validator) GetPeopleCharacter(id int) ([]model.PeopleCharacter, int, error) {
	if id <= 0 {
		return nil, http.StatusBadRequest, errors.ErrInvalidID
	}

	// Check empty id.
	key := internal.GetKey(internal.KeyEmptyPeople, id)
	if v.isEmptyID(key) {
		return nil, http.StatusNotFound, errors.ErrNot200
	}

	// Parse.
	data, code, err := v.api.GetPeopleCharacter(id)

	// Save empty id.
	v.saveEmptyID(code, key)

	return data, code, err
}

// GetPeopleStaff to get people anime staff list.
func (v *Validator) GetPeopleStaff(id int) ([]model.Role, int, error) {
	if id <= 0 {
		return nil, http.StatusBadRequest, errors.ErrInvalidID
	}

	// Check empty id.
	key := internal.GetKey(internal.KeyEmptyPeople, id)
	if v.isEmptyID(key) {
		return nil, http.StatusNotFound, errors.ErrNot200
	}

	// Parse.
	data, code, err := v.api.GetPeopleStaff(id)

	// Save empty id.
	v.saveEmptyID(code, key)

	return data, code, err
}

// GetPeopleManga to get people published manga list.
func (v *Validator) GetPeopleManga(id int) ([]model.Role, int, error) {
	if id <= 0 {
		return nil, http.StatusBadRequest, errors.ErrInvalidID
	}

	// Check empty id.
	key := internal.GetKey(internal.KeyEmptyPeople, id)
	if v.isEmptyID(key) {
		return nil, http.StatusNotFound, errors.ErrNot200
	}

	// Parse.
	data, code, err := v.api.GetPeopleManga(id)

	// Save empty id.
	v.saveEmptyID(code, key)

	return data, code, err
}

// GetPeopleNews to get people news list.
func (v *Validator) GetPeopleNews(id int) ([]model.NewsItem, int, error) {
	if id <= 0 {
		return nil, http.StatusBadRequest, errors.ErrInvalidID
	}

	// Check empty id.
	key := internal.GetKey(internal.KeyEmptyPeople, id)
	if v.isEmptyID(key) {
		return nil, http.StatusNotFound, errors.ErrNot200
	}

	// Parse.
	data, code, err := v.api.GetPeopleNews(id)

	// Save empty id.
	v.saveEmptyID(code, key)

	return data, code, err
}

// GetPeopleArticle to get people featured article list.
func (v *Validator) GetPeopleArticle(id int) ([]model.ArticleItem, int, error) {
	if id <= 0 {
		return nil, http.StatusBadRequest, errors.ErrInvalidID
	}

	// Check empty id.
	key := internal.GetKey(internal.KeyEmptyPeople, id)
	if v.isEmptyID(key) {
		return nil, http.StatusNotFound, errors.ErrNot200
	}

	// Parse.
	data, code, err := v.api.GetPeopleArticle(id)

	// Save empty id.
	v.saveEmptyID(code, key)

	return data, code, err
}

// GetPeoplePicture to get people picture list.
func (v *Validator) GetPeoplePicture(id int) ([]string, int, error) {
	if id <= 0 {
		return nil, http.StatusBadRequest, errors.ErrInvalidID
	}

	// Check empty id.
	key := internal.GetKey(internal.KeyEmptyPeople, id)
	if v.isEmptyID(key) {
		return nil, http.StatusNotFound, errors.ErrNot200
	}

	// Parse.
	data, code, err := v.api.GetPeoplePicture(id)

	// Save empty id.
	v.saveEmptyID(code, key)

	return data, code, err
}
