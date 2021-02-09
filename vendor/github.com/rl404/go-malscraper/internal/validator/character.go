package validator

import (
	"net/http"

	"github.com/rl404/go-malscraper/errors"
	"github.com/rl404/go-malscraper/internal"
	"github.com/rl404/go-malscraper/model"
)

// GetCharacter to get character detail information.
func (v *Validator) GetCharacter(id int) (*model.Character, int, error) {
	if id <= 0 {
		return nil, http.StatusBadRequest, errors.ErrInvalidID
	}

	// Check empty id.
	key := internal.GetKey(internal.KeyEmptyChar, id)
	if v.isEmptyID(key) {
		return nil, http.StatusNotFound, errors.ErrNot200
	}

	// Parse.
	data, code, err := v.api.GetCharacter(id)

	// Save empty id.
	v.saveEmptyID(code, key)

	return data, code, err
}

// GetCharacterArticle to get character featured article list.
func (v *Validator) GetCharacterArticle(id int) ([]model.ArticleItem, int, error) {
	if id <= 0 {
		return nil, http.StatusBadRequest, errors.ErrInvalidID
	}

	// Check empty id.
	key := internal.GetKey(internal.KeyEmptyChar, id)
	if v.isEmptyID(key) {
		return nil, http.StatusNotFound, errors.ErrNot200
	}

	// Parse.
	data, code, err := v.api.GetCharacterArticle(id)

	// Save empty id.
	v.saveEmptyID(code, key)

	return data, code, err
}

// GetCharacterOgraphy to get character animeography/mangaography list.
func (v *Validator) GetCharacterOgraphy(t string, id int) ([]model.Role, int, error) {
	if t != AnimeType && t != MangaType {
		return nil, http.StatusBadRequest, errors.ErrInvalidType
	}
	if id <= 0 {
		return nil, http.StatusBadRequest, errors.ErrInvalidID
	}

	// Check empty id.
	key := internal.GetKey(internal.KeyEmptyChar, id)
	if v.isEmptyID(key) {
		return nil, http.StatusNotFound, errors.ErrNot200
	}

	// Parse.
	data, code, err := v.api.GetCharacterOgraphy(t, id)

	// Save empty id.
	v.saveEmptyID(code, key)

	return data, code, err
}

// GetCharacterPicture to get character picture list.
func (v *Validator) GetCharacterPicture(id int) ([]string, int, error) {
	if id <= 0 {
		return nil, http.StatusBadRequest, errors.ErrInvalidID
	}

	// Check empty id.
	key := internal.GetKey(internal.KeyEmptyChar, id)
	if v.isEmptyID(key) {
		return nil, http.StatusNotFound, errors.ErrNot200
	}

	// Parse.
	data, code, err := v.api.GetCharacterPicture(id)

	// Save empty id.
	v.saveEmptyID(code, key)

	return data, code, err
}

// GetCharacterClub to get character club list.
func (v *Validator) GetCharacterClub(id int) ([]model.ClubItem, int, error) {
	if id <= 0 {
		return nil, http.StatusBadRequest, errors.ErrInvalidID
	}

	// Check empty id.
	key := internal.GetKey(internal.KeyEmptyChar, id)
	if v.isEmptyID(key) {
		return nil, http.StatusNotFound, errors.ErrNot200
	}

	// Parse.
	data, code, err := v.api.GetCharacterClub(id)

	// Save empty id.
	v.saveEmptyID(code, key)

	return data, code, err
}

// GetCharacterVA to get character voice actor list.
func (v *Validator) GetCharacterVA(id int) ([]model.Role, int, error) {
	if id <= 0 {
		return nil, http.StatusBadRequest, errors.ErrInvalidID
	}

	// Check empty id.
	key := internal.GetKey(internal.KeyEmptyChar, id)
	if v.isEmptyID(key) {
		return nil, http.StatusNotFound, errors.ErrNot200
	}

	// Parse.
	data, code, err := v.api.GetCharacterVA(id)

	// Save empty id.
	v.saveEmptyID(code, key)

	return data, code, err
}
