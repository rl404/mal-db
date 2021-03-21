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
	return v.api.GetMangaCharacter(id, page, limit)
}
