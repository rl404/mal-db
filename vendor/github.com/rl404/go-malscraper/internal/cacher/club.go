package cacher

import (
	"net/http"

	"github.com/rl404/go-malscraper/internal"
	"github.com/rl404/go-malscraper/model"
)

// GetClubs to get club list.
func (c *Cacher) GetClubs(page int) (data []model.ClubSearch, code int, err error) {
	// Get from cache.
	key := internal.GetKey(internal.KeyClubs, page)
	if c.cacher.Get(key, &data) == nil {
		return data, http.StatusOK, nil
	}

	// Parse.
	data, code, err = c.api.GetClubs(page)
	if err != nil {
		return nil, code, err
	}

	// Save to cache. Won't return error.
	_ = c.cacher.Set(key, data)
	return data, http.StatusOK, nil
}

// GetClub to get club detail information.
func (c *Cacher) GetClub(id int) (data *model.Club, code int, err error) {
	// Get from cache.
	key := internal.GetKey(internal.KeyClub, id)
	if c.cacher.Get(key, &data) == nil {
		return data, http.StatusOK, nil
	}

	// Parse.
	data, code, err = c.api.GetClub(id)
	if err != nil {
		return nil, code, err
	}

	// Save to cache. Won't return error.
	_ = c.cacher.Set(key, data)
	return data, http.StatusOK, nil
}

// GetClubMember to get club member list.
func (c *Cacher) GetClubMember(id int, page int) (data []model.ClubMember, code int, err error) {
	// Get from cache.
	key := internal.GetKey(internal.KeyClubMember, id, page)
	if c.cacher.Get(key, &data) == nil {
		return data, http.StatusOK, nil
	}

	// Parse.
	data, code, err = c.api.GetClubMember(id, page)
	if err != nil {
		return nil, code, err
	}

	// Save to cache. Won't return error.
	_ = c.cacher.Set(key, data)
	return data, http.StatusOK, nil
}

// GetClubPicture to get club picture list.
func (c *Cacher) GetClubPicture(id int) (data []string, code int, err error) {
	// Get from cache.
	key := internal.GetKey(internal.KeyClubPicture, id)
	if c.cacher.Get(key, &data) == nil {
		return data, http.StatusOK, nil
	}

	// Parse.
	data, code, err = c.api.GetClubPicture(id)
	if err != nil {
		return nil, code, err
	}

	// Save to cache. Won't return error.
	_ = c.cacher.Set(key, data)
	return data, http.StatusOK, nil
}

// GetClubRelated to get club related list.
func (c *Cacher) GetClubRelated(id int) (data *model.ClubRelated, code int, err error) {
	// Get from cache.
	key := internal.GetKey(internal.KeyClubRelated, id)
	if c.cacher.Get(key, &data) == nil {
		return data, http.StatusOK, nil
	}

	// Parse.
	data, code, err = c.api.GetClubRelated(id)
	if err != nil {
		return nil, code, err
	}

	// Save to cache. Won't return error.
	_ = c.cacher.Set(key, data)
	return data, http.StatusOK, nil
}
