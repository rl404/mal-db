package cacher

import (
	"net/http"

	"github.com/rl404/mal-db/internal/constant"
	"github.com/rl404/mal-db/internal/model"
)

// GetManga to get manga.
func (c *Cacher) GetManga(id int) (*model.Manga, map[string]interface{}, int, error) {
	// Cache model.
	var data struct {
		Data *model.Manga
		Meta map[string]interface{}
	}

	// Get from cache.
	key := constant.GetKey(constant.KeyManga, id)
	if c.cacher.Get(key, &data) == nil {
		return data.Data, data.Meta, http.StatusOK, nil
	}

	// Query db.
	d, meta, code, err := c.api.GetManga(id)
	if err != nil {
		return nil, nil, code, err
	}

	// Save to cache. Won't return error.
	data.Data, data.Meta = d, meta
	go c.cacher.Set(key, data)

	return d, meta, http.StatusOK, nil
}

// GetMangaCharacter to get manga character list.
func (c *Cacher) GetMangaCharacter(id int, page int, limit int) ([]model.Role, map[string]interface{}, int, error) {
	// Cache model.
	var data struct {
		Data []model.Role
		Meta map[string]interface{}
	}

	// Get from cache.
	key := constant.GetKey(constant.KeyMangaCharacter, id)
	if c.cacher.Get(key, &data) == nil {
		return data.Data, data.Meta, http.StatusOK, nil
	}

	// Query db.
	d, meta, code, err := c.api.GetMangaCharacter(id, page, limit)
	if err != nil {
		return nil, nil, code, err
	}

	// Save to cache. Won't return error.
	data.Data, data.Meta = d, meta
	go c.cacher.Set(key, data)

	return d, meta, http.StatusOK, nil
}
