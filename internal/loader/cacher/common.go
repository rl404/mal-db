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
