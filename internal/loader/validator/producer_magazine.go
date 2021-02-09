package validator

import (
	"net/http"

	"github.com/rl404/mal-db/internal/constant"
	"github.com/rl404/mal-db/internal/errors"
	"github.com/rl404/mal-db/internal/model"
)

// GetProducerMagazine to get all producer/magazine list.
func (v *Validator) GetProducerMagazine(t string) ([]model.Item, map[string]interface{}, int, error) {
	if t != constant.AnimeType && t != constant.MangaType {
		return nil, nil, http.StatusBadRequest, errors.ErrInvalidType
	}
	return v.api.GetProducerMagazine(t)
}
