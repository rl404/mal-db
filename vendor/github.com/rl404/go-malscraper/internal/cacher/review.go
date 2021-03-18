package cacher

import (
	"net/http"

	"github.com/rl404/go-malscraper/internal"
	"github.com/rl404/go-malscraper/model"
)

// GetReview to get review detail information.
func (c *Cacher) GetReview(id int) (data *model.Review, code int, err error) {
	// Get from cache.
	key := internal.GetKey(internal.KeyReview, id)
	if c.cacher.Get(key, &data) == nil {
		return data, http.StatusOK, nil
	}

	// Parse.
	data, code, err = c.api.GetReview(id)
	if err != nil {
		return nil, code, err
	}

	// Save to cache. Won't return error.
	_ = c.cacher.Set(key, data)
	return data, http.StatusOK, nil
}

// GetReviews to get anime/manga/best review list.
func (c *Cacher) GetReviews(t string, page int) (data []model.Review, code int, err error) {
	// Get from cache.
	key := internal.GetKey(internal.KeyReviews, t, page)
	if c.cacher.Get(key, &data) == nil {
		return data, http.StatusOK, nil
	}

	// Parse.
	data, code, err = c.api.GetReviews(t, page)
	if err != nil {
		return nil, code, err
	}

	// Save to cache. Won't return error.
	_ = c.cacher.Set(key, data)
	return data, http.StatusOK, nil
}
