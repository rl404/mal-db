package cacher

import (
	"net/http"

	"github.com/rl404/mal-db/internal/constant"
	"github.com/rl404/mal-db/internal/model"
)

// GetProducerMagazine to get all producer/magazine list.
func (c *Cacher) GetProducerMagazine(t string) ([]model.Item, map[string]interface{}, int, error) {
	// Cache model.
	var data struct {
		Data []model.Item
		Meta map[string]interface{}
	}

	// Get from cache.
	key := constant.GetKey(constant.KeyProducerMagazine, t)
	if c.cacher.Get(key, &data) == nil {
		return data.Data, data.Meta, http.StatusOK, nil
	}

	// Query db.
	d, meta, code, err := c.api.GetProducerMagazine(t)
	if err != nil {
		return nil, nil, code, err
	}

	// Save to cache. Won't return error.
	data.Data, data.Meta = d, meta
	_ = c.cacher.Set(key, data)

	return d, meta, http.StatusOK, nil
}
