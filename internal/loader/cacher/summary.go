package cacher

import (
	"net/http"

	"github.com/rl404/mal-db/internal/constant"
	"github.com/rl404/mal-db/internal/model"
)

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
