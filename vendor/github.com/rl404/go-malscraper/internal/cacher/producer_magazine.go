package cacher

import (
	"net/http"

	"github.com/rl404/go-malscraper/internal"
	"github.com/rl404/go-malscraper/model"
)

// GetProducers to get anime producer/studio/licensor list.
func (c *Cacher) GetProducers() (data []model.ItemCount, code int, err error) {
	// Get from cache.
	key := internal.GetKey(internal.KeyProducers)
	if c.cacher.Get(key, &data) == nil {
		return data, http.StatusOK, nil
	}

	// Parse.
	data, code, err = c.api.GetProducers()
	if err != nil {
		return nil, code, err
	}

	// Save to cache. Won't return error.
	_ = c.cacher.Set(key, data)
	return data, http.StatusOK, nil
}

// GetProducer to get producer anime list.
func (c *Cacher) GetProducer(id int, page int) (data []model.AnimeItem, code int, err error) {
	// Get from cache.
	key := internal.GetKey(internal.KeyProducer, id, page)
	if c.cacher.Get(key, &data) == nil {
		return data, http.StatusOK, nil
	}

	// Parse.
	data, code, err = c.api.GetProducer(id, page)
	if err != nil {
		return nil, code, err
	}

	// Save to cache. Won't return error.
	_ = c.cacher.Set(key, data)
	return data, http.StatusOK, nil
}

// GetMagazines to get manga magazine/serialization list.
func (c *Cacher) GetMagazines() (data []model.ItemCount, code int, err error) {
	// Get from cache.
	key := internal.GetKey(internal.KeyMagazines)
	if c.cacher.Get(key, &data) == nil {
		return data, http.StatusOK, nil
	}

	// Parse.
	data, code, err = c.api.GetMagazines()
	if err != nil {
		return nil, code, err
	}

	// Save to cache. Won't return error.
	_ = c.cacher.Set(key, data)
	return data, http.StatusOK, nil
}

// GetMagazine to get magazine manga list.
func (c *Cacher) GetMagazine(id int, page int) (data []model.MangaItem, code int, err error) {
	// Get from cache.
	key := internal.GetKey(internal.KeyMagazine, id, page)
	if c.cacher.Get(key, &data) == nil {
		return data, http.StatusOK, nil
	}

	// Parse.
	data, code, err = c.api.GetMagazine(id, page)
	if err != nil {
		return nil, code, err
	}

	// Save to cache. Won't return error.
	_ = c.cacher.Set(key, data)
	return data, http.StatusOK, nil
}
