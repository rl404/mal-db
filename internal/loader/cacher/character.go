package cacher

import (
	"net/http"

	"github.com/rl404/mal-db/internal/constant"
	"github.com/rl404/mal-db/internal/model"
)

// GetCharacter to get character.
func (c *Cacher) GetCharacter(id int) (*model.Character, map[string]interface{}, int, error) {
	// Cache model.
	var data struct {
		Data *model.Character
		Meta map[string]interface{}
	}

	// Get from cache.
	key := constant.GetKey(constant.KeyCharacter, id)
	if c.cacher.Get(key, &data) == nil {
		return data.Data, data.Meta, http.StatusOK, nil
	}

	// Query db.
	d, meta, code, err := c.api.GetCharacter(id)
	if err != nil {
		return nil, nil, code, err
	}

	// Save to cache. Won't return error.
	data.Data, data.Meta = d, meta
	_ = c.cacher.Set(key, data)

	return d, meta, http.StatusOK, nil
}

// GetCharacterOgraphy to get character anime/manga ography list.
func (c *Cacher) GetCharacterOgraphy(id int, t string, page int, limit int) ([]model.Role, map[string]interface{}, int, error) {
	// Cache model.
	var data struct {
		Data []model.Role
		Meta map[string]interface{}
	}

	// Get from cache.
	key := constant.GetKey(constant.KeyCharacterOgraphy, id, t)
	if c.cacher.Get(key, &data) == nil {
		return data.Data, data.Meta, http.StatusOK, nil
	}

	// Query db.
	d, meta, code, err := c.api.GetCharacterOgraphy(id, t, page, limit)
	if err != nil {
		return nil, nil, code, err
	}

	// Save to cache. Won't return error.
	data.Data, data.Meta = d, meta
	_ = c.cacher.Set(key, data)

	return d, meta, http.StatusOK, nil
}

// GetCharacterVA to get character voice actor list.
func (c *Cacher) GetCharacterVA(id int, page int, limit int) ([]model.Role, map[string]interface{}, int, error) {
	// Cache model.
	var data struct {
		Data []model.Role
		Meta map[string]interface{}
	}

	// Get from cache.
	key := constant.GetKey(constant.KeyCharacterVA, id)
	if c.cacher.Get(key, &data) == nil {
		return data.Data, data.Meta, http.StatusOK, nil
	}

	// Query db.
	d, meta, code, err := c.api.GetCharacterVA(id, page, limit)
	if err != nil {
		return nil, nil, code, err
	}

	// Save to cache. Won't return error.
	data.Data, data.Meta = d, meta
	_ = c.cacher.Set(key, data)

	return d, meta, http.StatusOK, nil
}
