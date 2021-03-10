package cacher

import (
	"net/http"

	"github.com/rl404/mal-db/internal/constant"
	"github.com/rl404/mal-db/internal/model"
)

// GetGenres to get all anime/manga genre list.
func (c *Cacher) GetGenres(t string) ([]model.Item, map[string]interface{}, int, error) {
	// Cache model.
	var data struct {
		Data []model.Item
		Meta map[string]interface{}
	}

	// Get from cache.
	key := constant.GetKey(constant.KeyGenres, t)
	if c.cacher.Get(key, &data) == nil {
		return data.Data, data.Meta, http.StatusOK, nil
	}

	// Query db.
	d, meta, code, err := c.api.GetGenres(t)
	if err != nil {
		return nil, nil, code, err
	}

	// Save to cache. Won't return error.
	data.Data, data.Meta = d, meta
	go c.cacher.Set(key, data)

	return d, meta, http.StatusOK, nil
}
