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
	go c.cacher.Set(key, data)

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
	go c.cacher.Set(key, data)

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
	go c.cacher.Set(key, data)

	return data, meta, http.StatusOK, nil
}

// Enqueue to enqueue to be re-parsed.
func (c *Cacher) Enqueue(t string, id int) (int, error) {
	return c.api.Enqueue(t, id)
}

// GetStatsHistory to get entry stats history.
func (c *Cacher) GetStatsHistory(t string, id int, page int, limit int) (data []model.StatsHistory, code int, err error) {
	// Get from cache.
	key := constant.GetKey(constant.KeyStatsHistory, t, id)
	if c.cacher.Get(key, &data) == nil {
		return data, http.StatusOK, nil
	}

	// Query db.
	data, code, err = c.api.GetStatsHistory(t, id, page, limit)
	if err != nil {
		return nil, code, err
	}

	// Save to cache. Won't return error.
	go c.cacher.Set(key, data)

	return data, http.StatusOK, nil
}

// CompareScore to get entry score comparison.
func (c *Cacher) CompareScore(query model.CompareQuery) ([]model.ScoreComparison, map[string]interface{}, int, error) {
	// Cache model.
	var data struct {
		Data []model.ScoreComparison
		Meta map[string]interface{}
	}

	// Get from cache.
	key := constant.GetKey(constant.KeyScoreComparison, query.Order, query.Limit, query.Page, query.Title)
	if c.cacher.Get(key, &data) == nil {
		return data.Data, data.Meta, http.StatusOK, nil
	}

	// Query db.
	d, meta, code, err := c.api.CompareScore(query)
	if err != nil {
		return nil, nil, code, err
	}

	// Save to cache. Won't return error.
	data.Data, data.Meta = d, meta
	go c.cacher.Set(key, data)

	return d, meta, http.StatusOK, nil
}
