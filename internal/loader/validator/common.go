package validator

import (
	"net/http"

	"github.com/rl404/mal-db/internal/constant"
	"github.com/rl404/mal-db/internal/errors"
	"github.com/rl404/mal-db/internal/model"
)

// GetStats to get anime/manga stats.
func (v *Validator) GetStats(t string, id int) (*model.Stats, map[string]interface{}, int, error) {
	if t != constant.AnimeType && t != constant.MangaType {
		return nil, nil, http.StatusBadRequest, errors.ErrInvalidType
	}
	if id <= 0 {
		return nil, nil, http.StatusBadRequest, errors.ErrInvalidID
	}
	return v.api.GetStats(t, id)
}
