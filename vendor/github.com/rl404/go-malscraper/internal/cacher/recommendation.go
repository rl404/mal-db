package cacher

import (
	"net/http"

	"github.com/rl404/go-malscraper/internal"
	"github.com/rl404/go-malscraper/model"
)

// GetRecommendation to get recommendation detail information.
func (c *Cacher) GetRecommendation(rType string, id1, id2 int) (data *model.Recommendation, code int, err error) {
	// Get from cache.
	key := internal.GetKey(internal.KeyRecommendation, rType, id1, id2)
	if c.cacher.Get(key, &data) == nil {
		return data, http.StatusOK, nil
	}

	// Parse.
	data, code, err = c.api.GetRecommendation(rType, id1, id2)
	if err != nil {
		return nil, code, err
	}

	// Save to cache. Won't return error.
	_ = c.cacher.Set(key, data)
	return data, http.StatusOK, nil
}

// GetRecommendations to get anime/manga recommendation list.
func (c *Cacher) GetRecommendations(t string, page int) (data []model.Recommendation, code int, err error) {
	// Get from cache.
	key := internal.GetKey(internal.KeyRecommendations, t, page)
	if c.cacher.Get(key, &data) == nil {
		return data, http.StatusOK, nil
	}

	// Parse.
	data, code, err = c.api.GetRecommendations(t, page)
	if err != nil {
		return nil, code, err
	}

	// Save to cache. Won't return error.
	_ = c.cacher.Set(key, data)
	return data, http.StatusOK, nil
}
