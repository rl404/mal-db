package validator

import (
	"net/http"

	"github.com/rl404/go-malscraper/pkg/utils"
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

// GetEntryCount to get all entry count.
func (v *Validator) GetEntryCount() (*model.Total, map[string]interface{}, int, error) {
	return v.api.GetEntryCount()
}

// GetYearSummary to get yearly anime & manga summary.
func (v *Validator) GetYearSummary() ([]model.YearSummary, map[string]interface{}, int, error) {
	return v.api.GetYearSummary()
}

// Enqueue to enqueue to be re-parsed.
func (v *Validator) Enqueue(t string, id int) (int, error) {
	if !utils.InArrayStr(constant.MainTypes, t) {
		return http.StatusBadRequest, errors.ErrInvalidType
	}
	if id <= 0 {
		return http.StatusBadRequest, errors.ErrInvalidID
	}
	return v.api.Enqueue(t, id)
}
