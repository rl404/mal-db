package cacher

import (
	"net/http"

	"github.com/rl404/go-malscraper/internal"
	"github.com/rl404/go-malscraper/model"
)

// GetTopAnime to get top anime list.
func (c *Cacher) GetTopAnime(t int, page int) (data []model.TopAnime, code int, err error) {
	// Get from cache.
	key := internal.GetKey(internal.KeyTopAnime, t, page)
	if c.cacher.Get(key, &data) == nil {
		return data, http.StatusOK, nil
	}

	// Parse.
	data, code, err = c.api.GetTopAnime(t, page)
	if err != nil {
		return nil, code, err
	}

	// Save to cache. Won't return error.
	_ = c.cacher.Set(key, data)
	return data, http.StatusOK, nil
}

// GetTopManga to get top manga list.
func (c *Cacher) GetTopManga(t int, page int) (data []model.TopManga, code int, err error) {
	// Get from cache.
	key := internal.GetKey(internal.KeyTopManga, t, page)
	if c.cacher.Get(key, &data) == nil {
		return data, http.StatusOK, nil
	}

	// Parse.
	data, code, err = c.api.GetTopManga(t, page)
	if err != nil {
		return nil, code, err
	}

	// Save to cache. Won't return error.
	_ = c.cacher.Set(key, data)
	return data, http.StatusOK, nil
}

// GetTopCharacter to get top character list.
func (c *Cacher) GetTopCharacter(page int) (data []model.TopCharacter, code int, err error) {
	// Get from cache.
	key := internal.GetKey(internal.KeyTopCharacter, page)
	if c.cacher.Get(key, &data) == nil {
		return data, http.StatusOK, nil
	}

	// Parse.
	data, code, err = c.api.GetTopCharacter(page)
	if err != nil {
		return nil, code, err
	}

	// Save to cache. Won't return error.
	_ = c.cacher.Set(key, data)
	return data, http.StatusOK, nil
}

// GetTopPeople to get top people list.
func (c *Cacher) GetTopPeople(page int) (data []model.TopPeople, code int, err error) {
	// Get from cache.
	key := internal.GetKey(internal.KeyTopPeople, page)
	if c.cacher.Get(key, &data) == nil {
		return data, http.StatusOK, nil
	}

	// Parse.
	data, code, err = c.api.GetTopPeople(page)
	if err != nil {
		return nil, code, err
	}

	// Save to cache. Won't return error.
	_ = c.cacher.Set(key, data)
	return data, http.StatusOK, nil
}
