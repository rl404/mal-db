package validator

import (
	"net/http"

	"github.com/rl404/go-malscraper/errors"
	"github.com/rl404/go-malscraper/model"
	"github.com/rl404/go-malscraper/pkg/utils"
)

// SearchAnime to search anime with advanced query.
func (v *Validator) SearchAnime(query model.Query) ([]model.AnimeSearch, int, error) {
	if len(query.Title) < 3 {
		return nil, http.StatusBadRequest, errors.Err3LettersSearch
	}
	if query.Page < 0 {
		return nil, http.StatusBadRequest, errors.ErrInvalidPage
	}
	if query.Page == 0 {
		query.Page = 1
	}
	if !utils.InArrayInt(animeTypes, query.Type) {
		return nil, http.StatusBadRequest, errors.ErrInvalidType
	}
	if query.Score < 0 || query.Score > 10 {
		return nil, http.StatusBadRequest, errors.ErrInvalidScore
	}
	if !utils.InArrayInt(animeStatuses, query.Status) {
		return nil, http.StatusBadRequest, errors.ErrInvalidStatus
	}
	if query.ProducerID != 0 && !v.isProducerValid(query.ProducerID) {
		return nil, http.StatusBadRequest, errors.ErrInvalidProducer
	}
	for _, g := range query.GenreIDs {
		if !v.isAnimeGenreValid(g) {
			return nil, http.StatusBadRequest, errors.ErrInvalidGenre
		}
	}
	if !utils.InArrayInt(ratings, query.Rating) {
		return nil, http.StatusBadRequest, errors.ErrInvalidRating
	}
	if len(query.FirstLetter) > 1 {
		return nil, http.StatusBadRequest, errors.ErrInvalidFirstLetter
	}
	return v.api.SearchAnime(query)
}

// SearchManga to search manga with advanced query.
func (v *Validator) SearchManga(query model.Query) ([]model.MangaSearch, int, error) {
	if len(query.Title) < 3 {
		return nil, http.StatusBadRequest, errors.Err3LettersSearch
	}
	if query.Page < 0 {
		return nil, http.StatusBadRequest, errors.ErrInvalidPage
	}
	if query.Page == 0 {
		query.Page = 1
	}
	if !utils.InArrayInt(mangaTypes, query.Type) {
		return nil, http.StatusBadRequest, errors.ErrInvalidType
	}
	if query.Score < 0 || query.Score > 10 {
		return nil, http.StatusBadRequest, errors.ErrInvalidScore
	}
	if !utils.InArrayInt(mangaStatuses, query.Status) {
		return nil, http.StatusBadRequest, errors.ErrInvalidStatus
	}
	if query.MagazineID != 0 && !v.isMagazineValid(query.MagazineID) {
		return nil, http.StatusBadRequest, errors.ErrInvalidMagazine
	}
	for _, g := range query.GenreIDs {
		if !v.isMangaGenreValid(g) {
			return nil, http.StatusBadRequest, errors.ErrInvalidGenre
		}
	}
	if len(query.FirstLetter) > 1 {
		return nil, http.StatusBadRequest, errors.ErrInvalidFirstLetter
	}
	return v.api.SearchManga(query)
}

// SearchCharacter to search character.
func (v *Validator) SearchCharacter(name string, page int) ([]model.CharacterSearch, int, error) {
	if len(name) < 3 {
		return nil, http.StatusBadRequest, errors.Err3LettersSearch
	}
	if page <= 0 {
		return nil, http.StatusBadRequest, errors.ErrInvalidPage
	}
	return v.api.SearchCharacter(name, page)
}

// SearchPeople to search people.
func (v *Validator) SearchPeople(name string, page int) ([]model.PeopleSearch, int, error) {
	if len(name) < 3 {
		return nil, http.StatusBadRequest, errors.Err3LettersSearch
	}
	if page <= 0 {
		return nil, http.StatusBadRequest, errors.ErrInvalidPage
	}
	return v.api.SearchPeople(name, page)
}

// SearchClub to search club with advanced query.
func (v *Validator) SearchClub(query model.ClubQuery) ([]model.ClubSearch, int, error) {
	if len(query.Name) < 3 {
		return nil, http.StatusBadRequest, errors.Err3LettersSearch
	}
	if query.Page < 0 {
		return nil, http.StatusBadRequest, errors.ErrInvalidPage
	}
	if query.Page == 0 {
		query.Page = 1
	}
	if !utils.InArrayInt(categories, query.Category) {
		return nil, http.StatusBadRequest, errors.ErrInvalidClubCategory
	}
	if !utils.InArrayInt(sorts, query.Sort) {
		return nil, http.StatusBadRequest, errors.ErrInvalidSortType
	}
	return v.api.SearchClub(query)
}

// SearchUser to search club with advanced query.
func (v *Validator) SearchUser(query model.UserQuery) ([]model.UserSearch, int, error) {
	if len(query.Username) < 3 {
		return nil, http.StatusBadRequest, errors.Err3LettersSearch
	}
	if query.Page < 0 {
		return nil, http.StatusBadRequest, errors.ErrInvalidPage
	}
	if query.Page == 0 {
		query.Page = 1
	}
	if query.MinAge < 0 || query.MaxAge < 0 {
		return nil, http.StatusBadRequest, errors.ErrInvalidAge
	}
	if !utils.InArrayInt(genders, query.Gender) {
		return nil, http.StatusBadRequest, errors.ErrInvalidGender
	}
	return v.api.SearchUser(query)
}
