package validator

import (
	"net/http"
	"strconv"

	"github.com/rl404/go-malscraper/errors"
	"github.com/rl404/go-malscraper/model"
	"github.com/rl404/go-malscraper/pkg/utils"
)

// GetSeason to get seasonal anime list.
func (v *Validator) GetSeason(season string, year int) ([]model.AnimeItem, int, error) {
	if !utils.InArrayStr(seasons, season) {
		return nil, http.StatusBadRequest, errors.ErrInvalidSeason
	}
	if len(strconv.Itoa(year)) != 4 {
		return nil, http.StatusBadRequest, errors.ErrInvalidYear
	}
	return v.api.GetSeason(season, year)
}
