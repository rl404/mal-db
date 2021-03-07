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

// GetStatsHistory to get entry stats history.
func (v *Validator) GetStatsHistory(t string, id int) ([]model.StatsHistory, int, error) {
	if !utils.InArrayStr(constant.MainTypes, t) {
		return nil, http.StatusBadRequest, errors.ErrInvalidType
	}
	if id <= 0 {
		return nil, http.StatusBadRequest, errors.ErrInvalidID
	}
	return v.api.GetStatsHistory(t, id)
}

// CompareScore to get entry score comparison.
func (v *Validator) CompareScore(query model.CompareQuery) ([]model.ScoreComparison, map[string]interface{}, int, error) {
	if len(query.Title) > 0 && len(query.Title) < 3 {
		return nil, nil, http.StatusBadRequest, errors.Err3LettersSearch
	}
	if query.Page <= 0 {
		return nil, nil, http.StatusBadRequest, errors.ErrInvalidPage
	}
	if query.Limit <= 0 {
		return nil, nil, http.StatusBadRequest, errors.ErrInvalidLimit
	}
	if query.Order != "" && !utils.InArrayStr(constant.Orders, query.Order) {
		return nil, nil, http.StatusBadRequest, errors.ErrInvalidOrder
	}

	// Get data.
	data, meta, code, err := v.api.CompareScore(query)
	if err != nil {
		return nil, nil, code, err
	}

	// Handle pagination.
	start, current := query.Limit*(query.Page-1), len(data)-(query.Page-1)*query.Limit
	if current <= 0 {
		data = []model.ScoreComparison{}
	} else {
		if current < query.Limit {
			query.Limit = current
		}
		data = data[start : start+query.Limit]
	}

	return data, meta, http.StatusOK, nil
}
