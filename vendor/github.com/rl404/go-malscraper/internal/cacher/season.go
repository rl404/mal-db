package cacher

import (
	"net/http"

	"github.com/rl404/go-malscraper/internal"
	"github.com/rl404/go-malscraper/model"
)

// GetSeason to get seasonal anime list.
func (c *Cacher) GetSeason(season string, year int) (data []model.AnimeItem, code int, err error) {
	// Get from cache.
	key := internal.GetKey(internal.KeySeason, season, year)
	if c.cacher.Get(key, &data) == nil {
		return data, http.StatusOK, nil
	}

	// Parse.
	data, code, err = c.api.GetSeason(season, year)
	if err != nil {
		return nil, code, err
	}

	// Save to cache. Won't return error.
	_ = c.cacher.Set(key, data)
	return data, http.StatusOK, nil
}
