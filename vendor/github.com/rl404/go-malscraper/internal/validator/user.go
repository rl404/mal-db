package validator

import (
	"net/http"

	"github.com/rl404/go-malscraper/errors"
	"github.com/rl404/go-malscraper/internal"
	"github.com/rl404/go-malscraper/model"
	"github.com/rl404/go-malscraper/pkg/utils"
)

// GetUser to get user detail information.
func (v *Validator) GetUser(username string) (*model.User, int, error) {
	if len(username) < 2 || len(username) > 16 {
		return nil, http.StatusBadRequest, errors.ErrInvalidUsername
	}

	// Check empty id.
	key := internal.GetKey(internal.KeyEmptyUser, username)
	if v.isEmptyID(key) {
		return nil, http.StatusNotFound, errors.ErrNot200
	}

	// Parse.
	data, code, err := v.api.GetUser(username)

	// Save empty id.
	v.saveEmptyID(code, key)

	return data, code, err
}

// GetUserStats to get user stats detail information.
func (v *Validator) GetUserStats(username string) (*model.UserStats, int, error) {
	if len(username) < 2 || len(username) > 16 {
		return nil, http.StatusBadRequest, errors.ErrInvalidUsername
	}

	// Check empty id.
	key := internal.GetKey(internal.KeyEmptyUser, username)
	if v.isEmptyID(key) {
		return nil, http.StatusNotFound, errors.ErrNot200
	}

	// Parse.
	data, code, err := v.api.GetUserStats(username)

	// Save empty id.
	v.saveEmptyID(code, key)

	return data, code, err
}

// GetUserFavorite to get user favorite list.
func (v *Validator) GetUserFavorite(username string) (*model.UserFavorite, int, error) {
	if len(username) < 2 || len(username) > 16 {
		return nil, http.StatusBadRequest, errors.ErrInvalidUsername
	}

	// Check empty id.
	key := internal.GetKey(internal.KeyEmptyUser, username)
	if v.isEmptyID(key) {
		return nil, http.StatusNotFound, errors.ErrNot200
	}

	// Parse.
	data, code, err := v.api.GetUserFavorite(username)

	// Save empty id.
	v.saveEmptyID(code, key)

	return data, code, err
}

// GetUserFriend to get user friend list.
func (v *Validator) GetUserFriend(username string, page int) ([]model.UserFriend, int, error) {
	if len(username) < 2 || len(username) > 16 {
		return nil, http.StatusBadRequest, errors.ErrInvalidUsername
	}
	if page <= 0 {
		return nil, http.StatusBadRequest, errors.ErrInvalidPage
	}

	// Check empty id.
	key := internal.GetKey(internal.KeyEmptyUser, username)
	if v.isEmptyID(key) {
		return nil, http.StatusNotFound, errors.ErrNot200
	}

	// Parse.
	data, code, err := v.api.GetUserFriend(username, page)

	// Save empty id.
	v.saveEmptyID(code, key)

	return data, code, err
}

// GetUserHistory to get user history list.
func (v *Validator) GetUserHistory(username string, t string) ([]model.UserHistory, int, error) {
	if len(username) < 2 || len(username) > 16 {
		return nil, http.StatusBadRequest, errors.ErrInvalidUsername
	}
	if t != "" && t != AnimeType && t != MangaType {
		return nil, http.StatusBadRequest, errors.ErrInvalidType
	}

	// Check empty id.
	key := internal.GetKey(internal.KeyEmptyUser, username)
	if v.isEmptyID(key) {
		return nil, http.StatusNotFound, errors.ErrNot200
	}

	// Parse.
	data, code, err := v.api.GetUserHistory(username, t)

	// Save empty id.
	v.saveEmptyID(code, key)

	return data, code, err
}

// GetUserReview to get user review list.
func (v *Validator) GetUserReview(username string, page int) ([]model.Review, int, error) {
	if len(username) < 2 || len(username) > 16 {
		return nil, http.StatusBadRequest, errors.ErrInvalidUsername
	}
	if page <= 0 {
		return nil, http.StatusBadRequest, errors.ErrInvalidPage
	}

	// Check empty id.
	key := internal.GetKey(internal.KeyEmptyUser, username)
	if v.isEmptyID(key) {
		return nil, http.StatusNotFound, errors.ErrNot200
	}

	// Parse.
	data, code, err := v.api.GetUserReview(username, page)

	// Save empty id.
	v.saveEmptyID(code, key)

	return data, code, err
}

// GetUserRecommendation to get user recommendation list.
func (v *Validator) GetUserRecommendation(username string, page int) ([]model.Recommendation, int, error) {
	if len(username) < 2 || len(username) > 16 {
		return nil, http.StatusBadRequest, errors.ErrInvalidUsername
	}
	if page <= 0 {
		return nil, http.StatusBadRequest, errors.ErrInvalidPage
	}

	// Check empty id.
	key := internal.GetKey(internal.KeyEmptyUser, username)
	if v.isEmptyID(key) {
		return nil, http.StatusNotFound, errors.ErrNot200
	}

	// Parse.
	data, code, err := v.api.GetUserRecommendation(username, page)

	// Save empty id.
	v.saveEmptyID(code, key)

	return data, code, err
}

// GetUserClub to get user club list.
func (v *Validator) GetUserClub(username string) ([]model.Item, int, error) {
	if len(username) < 2 || len(username) > 16 {
		return nil, http.StatusBadRequest, errors.ErrInvalidUsername
	}

	// Check empty id.
	key := internal.GetKey(internal.KeyEmptyUser, username)
	if v.isEmptyID(key) {
		return nil, http.StatusNotFound, errors.ErrNot200
	}

	// Parse.
	data, code, err := v.api.GetUserClub(username)

	// Save empty id.
	v.saveEmptyID(code, key)

	return data, code, err
}

// GetUserAnime to get user anime list.
func (v *Validator) GetUserAnime(query model.UserListQuery) ([]model.UserAnime, int, error) {
	if len(query.Username) < 2 || len(query.Username) > 16 {
		return nil, http.StatusBadRequest, errors.ErrInvalidUsername
	}
	if query.Page == 0 {
		query.Page = 1
	}
	if query.Page < -1 {
		return nil, http.StatusBadRequest, errors.ErrInvalidPage
	}
	if !utils.InArrayInt(statuses, query.Status) {
		return nil, http.StatusBadRequest, errors.ErrInvalidStatus
	}
	if !utils.InArrayInt(animeOrders, query.Order) {
		return nil, http.StatusBadRequest, errors.ErrInvalidOrder
	}

	// Check empty id.
	key := internal.GetKey(internal.KeyEmptyUser, query.Username)
	if v.isEmptyID(key) {
		return nil, http.StatusNotFound, errors.ErrNot200
	}

	// Parse.
	data, code, err := v.api.GetUserAnime(query)

	// Save empty id.
	v.saveEmptyID(code, key)

	return data, code, err
}

// GetUserManga to get user manga list.
func (v *Validator) GetUserManga(query model.UserListQuery) ([]model.UserManga, int, error) {
	if len(query.Username) < 2 || len(query.Username) > 16 {
		return nil, http.StatusBadRequest, errors.ErrInvalidUsername
	}
	if query.Page == 0 {
		query.Page = 1
	}
	if query.Page < -1 {
		return nil, http.StatusBadRequest, errors.ErrInvalidPage
	}
	if !utils.InArrayInt(statuses, query.Status) {
		return nil, http.StatusBadRequest, errors.ErrInvalidStatus
	}
	if !utils.InArrayInt(mangaOrders, query.Order) {
		return nil, http.StatusBadRequest, errors.ErrInvalidOrder
	}

	// Check empty id.
	key := internal.GetKey(internal.KeyEmptyUser, query.Username)
	if v.isEmptyID(key) {
		return nil, http.StatusNotFound, errors.ErrNot200
	}

	// Parse.
	data, code, err := v.api.GetUserManga(query)

	// Save empty id.
	v.saveEmptyID(code, key)

	return data, code, err
}
