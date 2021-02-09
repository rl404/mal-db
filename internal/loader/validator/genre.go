package validator

import (
	"net/http"

	"github.com/rl404/mal-db/internal/constant"
	"github.com/rl404/mal-db/internal/errors"
	"github.com/rl404/mal-db/internal/model"
)

// GetGenres to get all anime/manga genre list.
func (v *Validator) GetGenres(t string) ([]model.Item, map[string]interface{}, int, error) {
	if t != constant.AnimeType && t != constant.MangaType {
		return nil, nil, http.StatusBadRequest, errors.ErrInvalidType
	}
	return v.api.GetGenres(t)
}
