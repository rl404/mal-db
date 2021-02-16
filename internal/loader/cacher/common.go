package cacher

import (
	"net/http"

	"github.com/rl404/mal-db/internal/constant"
	"github.com/rl404/mal-db/internal/model"
)

// GetStats to get anime/manga stats.
func (c *Cacher) GetStats(t string, id int) (data *model.Stats, meta map[string]interface{}, code int, err error) {
	// Get from cache.
	key := constant.GetKey(constant.KeyStats, t, id)
	if c.cacher.Get(key, &data) == nil {
		return data, meta, http.StatusOK, nil
	}

	// Query db.
	data, meta, code, err = c.api.GetStats(t, id)
	if err != nil {
		return nil, nil, code, err
	}

	// Save to cache. Won't return error.
	_ = c.cacher.Set(key, data)

	return data, meta, http.StatusOK, nil
}

// GetEntryCount to get all entry count.
func (c *Cacher) GetEntryCount() (data *model.Total, meta map[string]interface{}, code int, err error) {
	// Get from cache.
	key := constant.GetKey(constant.KeyTotal)
	if c.cacher.Get(key, &data) == nil {
		return data, nil, http.StatusOK, nil
	}

	// Query db.
	data, meta, code, err = c.api.GetEntryCount()
	if err != nil {
		return nil, nil, code, err
	}

	// Save to cache. Won't return error.
	_ = c.cacher.Set(key, data)

	return data, meta, http.StatusOK, nil
}

// GetYearSummary to get yearly anime & manga summary.
func (c *Cacher) GetYearSummary() (data []model.YearSummary, meta map[string]interface{}, code int, err error) {
	// Get from cache.
	key := constant.GetKey(constant.KeyYearSummary)
	if c.cacher.Get(key, &data) == nil {
		return data, nil, http.StatusOK, nil
	}

	// Query db.
	data, meta, code, err = c.api.GetYearSummary()
	if err != nil {
		return nil, nil, code, err
	}

	// Save to cache. Won't return error.
	_ = c.cacher.Set(key, data)

	return data, meta, http.StatusOK, nil
}

// Enqueue to enqueue to be re-parsed.
func (c *Cacher) Enqueue(t string, id int) (int, error) {
	return c.api.Enqueue(t, id)
}
