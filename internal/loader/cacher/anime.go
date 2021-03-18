package cacher

import (
	"net/http"

	"github.com/rl404/mal-db/internal/constant"
	"github.com/rl404/mal-db/internal/model"
)

// GetAnime to get anime.
func (c *Cacher) GetAnime(id int) (*model.Anime, map[string]interface{}, int, error) {
	// Cache model.
	var data struct {
		Data *model.Anime
		Meta map[string]interface{}
	}

	// Get from cache.
	key := constant.GetKey(constant.KeyAnime, id)
	if c.cacher.Get(key, &data) == nil {
		return data.Data, data.Meta, http.StatusOK, nil
	}

	// Query db.
	d, meta, code, err := c.api.GetAnime(id)
	if err != nil {
		return nil, nil, code, err
	}

	// Save to cache. Won't return error.
	data.Data, data.Meta = d, meta
	go c.cacher.Set(key, data)

	return d, meta, http.StatusOK, nil
}

// GetAnimeCharacter to get anime character list.
func (c *Cacher) GetAnimeCharacter(id int, page int, limit int) ([]model.AnimeCharacter, map[string]interface{}, int, error) {
	// Cache model.
	var data struct {
		Data []model.AnimeCharacter
		Meta map[string]interface{}
	}

	// Get from cache.
	key := constant.GetKey(constant.KeyAnimeCharacter, id, limit, page)
	if c.cacher.Get(key, &data) == nil {
		return data.Data, data.Meta, http.StatusOK, nil
	}

	// Query db.
	d, meta, code, err := c.api.GetAnimeCharacter(id, page, limit)
	if err != nil {
		return nil, nil, code, err
	}

	// Save to cache. Won't return error.
	data.Data, data.Meta = d, meta
	go c.cacher.Set(key, data)

	return d, meta, http.StatusOK, nil
}

// GetAnimeStaff to get anime staff list.
func (c *Cacher) GetAnimeStaff(id int, page int, limit int) ([]model.Role, map[string]interface{}, int, error) {
	// Cache model.
	var data struct {
		Data []model.Role
		Meta map[string]interface{}
	}

	// Get from cache.
	key := constant.GetKey(constant.KeyAnimeStaff, id)
	if c.cacher.Get(key, &data) == nil {
		return data.Data, data.Meta, http.StatusOK, nil
	}

	// Query db.
	d, meta, code, err := c.api.GetAnimeStaff(id, page, limit)
	if err != nil {
		return nil, nil, code, err
	}

	// Save to cache. Won't return error.
	data.Data, data.Meta = d, meta
	go c.cacher.Set(key, data)

	return d, meta, http.StatusOK, nil
}
