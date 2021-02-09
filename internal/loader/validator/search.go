package validator

import (
	"net/http"
	"strconv"

	"github.com/rl404/go-malscraper/pkg/utils"
	"github.com/rl404/mal-db/internal/constant"
	"github.com/rl404/mal-db/internal/errors"
	"github.com/rl404/mal-db/internal/model"
)

// SearchAnime to search anime.
func (v *Validator) SearchAnime(query model.AnimeQuery) ([]model.Media, map[string]interface{}, int, error) {
	if len(query.Title) > 0 && len(query.Title) < 3 {
		return nil, nil, http.StatusBadRequest, errors.Err3LettersSearch
	}
	if query.Page <= 0 {
		return nil, nil, http.StatusBadRequest, errors.ErrInvalidPage
	}
	if query.Limit <= 0 {
		return nil, nil, http.StatusBadRequest, errors.ErrInvalidLimit
	}
	if query.Score < 0 || query.Score > 10 {
		return nil, nil, http.StatusBadRequest, errors.ErrInvalidScore
	}
	if query.Type != 0 && constant.Types[constant.AnimeType][query.Type] == "" {
		return nil, nil, http.StatusBadRequest, errors.ErrInvalidType
	}
	if query.Status != 0 && constant.Statuses[constant.AnimeType][query.Status] == "" {
		return nil, nil, http.StatusBadRequest, errors.ErrInvalidStatus
	}
	if query.Rating != 0 && constant.Ratings[query.Rating] == "" {
		return nil, nil, http.StatusBadRequest, errors.ErrInvalidRating
	}
	if query.Source != 0 && constant.Sources[query.Source] == "" {
		return nil, nil, http.StatusBadRequest, errors.ErrInvalidSource
	}
	if query.Year < 0 || (query.Year > 0 && len(strconv.Itoa(query.Year)) != 4) {
		return nil, nil, http.StatusBadRequest, errors.ErrInvalidYear
	}
	if query.Season != "" && !utils.InArrayStr(constant.Seasons, query.Season) {
		return nil, nil, http.StatusBadRequest, errors.ErrInvalidSeason
	}
	if query.Order != "" && !utils.InArrayStr(constant.Orders, query.Order) {
		return nil, nil, http.StatusBadRequest, errors.ErrInvalidOrder
	}
	if query.StartYear < 0 || query.EndYear < 0 {
		return nil, nil, http.StatusBadRequest, errors.ErrInvalidYear
	}
	if query.StartEpisode < 0 || query.EndEpisode < 0 {
		return nil, nil, http.StatusBadRequest, errors.ErrInvalidEpisode
	}
	if query.StartDuration < 0 || query.EndDuration < 0 {
		return nil, nil, http.StatusBadRequest, errors.ErrInvalidDuration
	}
	if query.Producer < 0 {
		return nil, nil, http.StatusBadRequest, errors.ErrInvalidProducer
	}
	return v.api.SearchAnime(query)
}

// SearchManga to search manga.
func (v *Validator) SearchManga(query model.MangaQuery) ([]model.Media, map[string]interface{}, int, error) {
	if len(query.Title) > 0 && len(query.Title) < 3 {
		return nil, nil, http.StatusBadRequest, errors.Err3LettersSearch
	}
	if query.Page <= 0 {
		return nil, nil, http.StatusBadRequest, errors.ErrInvalidPage
	}
	if query.Limit <= 0 {
		return nil, nil, http.StatusBadRequest, errors.ErrInvalidLimit
	}
	if query.Score < 0 || query.Score > 10 {
		return nil, nil, http.StatusBadRequest, errors.ErrInvalidScore
	}
	if query.Type != 0 && constant.Types[constant.MangaType][query.Type] == "" {
		return nil, nil, http.StatusBadRequest, errors.ErrInvalidType
	}
	if query.Status != 0 && constant.Statuses[constant.MangaType][query.Status] == "" {
		return nil, nil, http.StatusBadRequest, errors.ErrInvalidStatus
	}
	if query.Year < 0 || (query.Year > 0 && len(strconv.Itoa(query.Year)) != 4) {
		return nil, nil, http.StatusBadRequest, errors.ErrInvalidYear
	}
	if query.Order != "" && !utils.InArrayStr(constant.Orders, query.Order) {
		return nil, nil, http.StatusBadRequest, errors.ErrInvalidOrder
	}
	if query.StartYear < 0 || query.EndYear < 0 {
		return nil, nil, http.StatusBadRequest, errors.ErrInvalidYear
	}
	if query.StartChapter < 0 || query.EndChapter < 0 {
		return nil, nil, http.StatusBadRequest, errors.ErrInvalidChapter
	}
	if query.StartVolume < 0 || query.EndVolume < 0 {
		return nil, nil, http.StatusBadRequest, errors.ErrInvalidVolume
	}
	if query.Magazine < 0 {
		return nil, nil, http.StatusBadRequest, errors.ErrInvalidMagazine
	}
	return v.api.SearchManga(query)
}

// SearchCharacter to search character.
func (v *Validator) SearchCharacter(query model.EntryQuery) ([]model.Entry, map[string]interface{}, int, error) {
	if len(query.Name) > 0 && len(query.Name) < 3 {
		return nil, nil, http.StatusBadRequest, errors.Err3LettersSearch
	}
	if query.Page <= 0 {
		return nil, nil, http.StatusBadRequest, errors.ErrInvalidPage
	}
	if query.Limit <= 0 {
		return nil, nil, http.StatusBadRequest, errors.ErrInvalidLimit
	}
	if query.Order != "" && !utils.InArrayStr(constant.Orders2, query.Order) {
		return nil, nil, http.StatusBadRequest, errors.ErrInvalidOrder
	}
	return v.api.SearchCharacter(query)
}

// SearchPeople to search people.
func (v *Validator) SearchPeople(query model.EntryQuery) ([]model.Entry, map[string]interface{}, int, error) {
	if len(query.Name) > 0 && len(query.Name) < 3 {
		return nil, nil, http.StatusBadRequest, errors.Err3LettersSearch
	}
	if query.Page <= 0 {
		return nil, nil, http.StatusBadRequest, errors.ErrInvalidPage
	}
	if query.Limit <= 0 {
		return nil, nil, http.StatusBadRequest, errors.ErrInvalidLimit
	}
	if query.Order != "" && !utils.InArrayStr(constant.Orders2, query.Order) {
		return nil, nil, http.StatusBadRequest, errors.ErrInvalidOrder
	}
	return v.api.SearchPeople(query)
}
