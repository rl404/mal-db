package cacher

import (
	"net/http"

	"github.com/rl404/mal-db/internal/constant"
	"github.com/rl404/mal-db/internal/model"
)

// GetPeople to get people.
func (c *Cacher) GetPeople(id int) (*model.People, map[string]interface{}, int, error) {
	// Cache model.
	var data struct {
		Data *model.People
		Meta map[string]interface{}
	}

	// Get from cache.
	key := constant.GetKey(constant.KeyPeople, id)
	if c.cacher.Get(key, &data) == nil {
		return data.Data, data.Meta, http.StatusOK, nil
	}

	// Query db.
	d, meta, code, err := c.api.GetPeople(id)
	if err != nil {
		return nil, nil, code, err
	}

	// Save to cache. Won't return error.
	data.Data, data.Meta = d, meta
	go c.cacher.Set(key, data)

	return d, meta, http.StatusOK, nil
}

// GetPeopleVA to get people voice actor role list.
func (c *Cacher) GetPeopleVA(id int, page int, limit int) ([]model.VoiceActor, map[string]interface{}, int, error) {
	// Cache model.
	var data struct {
		Data []model.VoiceActor
		Meta map[string]interface{}
	}

	// Get from cache.
	key := constant.GetKey(constant.KeyPeopleVA, id)
	if c.cacher.Get(key, &data) == nil {
		return data.Data, data.Meta, http.StatusOK, nil
	}

	// Query db.
	d, meta, code, err := c.api.GetPeopleVA(id, page, limit)
	if err != nil {
		return nil, nil, code, err
	}

	// Save to cache. Won't return error.
	data.Data, data.Meta = d, meta
	go c.cacher.Set(key, data)

	return d, meta, http.StatusOK, nil
}

// GetPeopleStaff to get people anime staff role list.
func (c *Cacher) GetPeopleStaff(id int, page int, limit int) ([]model.Role, map[string]interface{}, int, error) {
	// Cache model.
	var data struct {
		Data []model.Role
		Meta map[string]interface{}
	}

	// Get from cache.
	key := constant.GetKey(constant.KeyPeopleStaff, id)
	if c.cacher.Get(key, &data) == nil {
		return data.Data, data.Meta, http.StatusOK, nil
	}

	// Query db.
	d, meta, code, err := c.api.GetPeopleStaff(id, page, limit)
	if err != nil {
		return nil, nil, code, err
	}

	// Save to cache. Won't return error.
	data.Data, data.Meta = d, meta
	go c.cacher.Set(key, data)

	return d, meta, http.StatusOK, nil
}

// GetPeopleManga to get people published manga list.
func (c *Cacher) GetPeopleManga(id int, page int, limit int) ([]model.Role, map[string]interface{}, int, error) {
	// Cache model.
	var data struct {
		Data []model.Role
		Meta map[string]interface{}
	}

	// Get from cache.
	key := constant.GetKey(constant.KeyPeopleManga, id)
	if c.cacher.Get(key, &data) == nil {
		return data.Data, data.Meta, http.StatusOK, nil
	}

	// Query db.
	d, meta, code, err := c.api.GetPeopleManga(id, page, limit)
	if err != nil {
		return nil, nil, code, err
	}

	// Save to cache. Won't return error.
	data.Data, data.Meta = d, meta
	go c.cacher.Set(key, data)

	return d, meta, http.StatusOK, nil
}
