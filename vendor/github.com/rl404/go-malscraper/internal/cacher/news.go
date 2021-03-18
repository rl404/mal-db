package cacher

import (
	"net/http"

	"github.com/rl404/go-malscraper/internal"
	"github.com/rl404/go-malscraper/model"
)

// GetNews to get news detail information.
func (c *Cacher) GetNews(id int) (data *model.News, code int, err error) {
	// Get from cache.
	key := internal.GetKey(internal.KeyNews, id)
	if c.cacher.Get(key, &data) == nil {
		return data, http.StatusOK, nil
	}

	// Parse.
	data, code, err = c.api.GetNews(id)
	if err != nil {
		return nil, code, err
	}

	// Save to cache. Won't return error.
	_ = c.cacher.Set(key, data)
	return data, http.StatusOK, nil
}

// GetNewsList to get news list.
func (c *Cacher) GetNewsList(page int, tag string) (data []model.NewsItem, code int, err error) {
	// Get from cache.
	key := internal.GetKey(internal.KeyNewsList, page, tag)
	if c.cacher.Get(key, &data) == nil {
		return data, http.StatusOK, nil
	}

	// Parse.
	data, code, err = c.api.GetNewsList(page, tag)
	if err != nil {
		return nil, code, err
	}

	// Save to cache. Won't return error.
	_ = c.cacher.Set(key, data)
	return data, http.StatusOK, nil
}

// GetNewsTag to get news tag list.
func (c *Cacher) GetNewsTag() (data *model.NewsTag, code int, err error) {
	// Get from cache.
	key := internal.GetKey(internal.KeyNewsTag)
	if c.cacher.Get(key, &data) == nil {
		return data, http.StatusOK, nil
	}

	// Parse.
	data, code, err = c.api.GetNewsTag()
	if err != nil {
		return nil, code, err
	}

	// Save to cache. Won't return error.
	_ = c.cacher.Set(key, data)
	return data, http.StatusOK, nil
}
