package cacher

import (
	"net/http"

	"github.com/rl404/go-malscraper/internal"
	"github.com/rl404/go-malscraper/model"
)

// GetArticle to get featured article detail information.
func (c *Cacher) GetArticle(id int) (data *model.Article, code int, err error) {
	// Get from cache.
	key := internal.GetKey(internal.KeyArticle, id)
	if c.cacher.Get(key, &data) == nil {
		return data, http.StatusOK, nil
	}

	// Parse.
	data, code, err = c.api.GetArticle(id)
	if err != nil {
		return nil, code, err
	}

	// Save to cache. Won't return error.
	_ = c.cacher.Set(key, data)
	return data, http.StatusOK, nil
}

// GetArticles to get featured article list.
func (c *Cacher) GetArticles(page int, tag string) (data []model.ArticleItem, code int, err error) {
	// Get from cache.
	key := internal.GetKey(internal.KeyArticleList, page, tag)
	if c.cacher.Get(key, &data) == nil {
		return data, http.StatusOK, nil
	}

	// Parse.
	data, code, err = c.api.GetArticles(page, tag)
	if err != nil {
		return nil, code, err
	}

	// Save to cache. Won't return error.
	_ = c.cacher.Set(key, data)
	return data, http.StatusOK, nil
}

// GetArticleTag to get featured article tag list.
func (c *Cacher) GetArticleTag() (data []model.ArticleTagItem, code int, err error) {
	// Get from cache.
	key := internal.GetKey(internal.KeyArticleTag)
	if c.cacher.Get(key, &data) == nil {
		return data, http.StatusOK, nil
	}

	// Parse.
	data, code, err = c.api.GetArticleTag()
	if err != nil {
		return nil, code, err
	}

	// Save to cache. Won't return error.
	_ = c.cacher.Set(key, data)
	return data, http.StatusOK, nil
}
