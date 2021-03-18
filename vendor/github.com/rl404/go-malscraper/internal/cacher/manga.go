package cacher

import (
	"net/http"

	"github.com/rl404/go-malscraper/internal"
	"github.com/rl404/go-malscraper/model"
)

// GetManga to get manga from cache.
func (c *Cacher) GetManga(id int) (data *model.Manga, code int, err error) {
	// Get from cache.
	key := internal.GetKey(internal.KeyManga, id)
	if c.cacher.Get(key, &data) == nil {
		return data, http.StatusOK, nil
	}

	// Parse.
	data, code, err = c.api.GetManga(id)
	if err != nil {
		return nil, code, err
	}

	// Save to cache. Won't return error.
	_ = c.cacher.Set(key, data)
	return data, http.StatusOK, nil
}

// GetMangaReview to get manga review list.
func (c *Cacher) GetMangaReview(id int, page int) (data []model.Review, code int, err error) {
	// Get from cache.
	key := internal.GetKey(internal.KeyMangaReview, id, page)
	if c.cacher.Get(key, &data) == nil {
		return data, http.StatusOK, nil
	}

	// Parse.
	data, code, err = c.api.GetMangaReview(id, page)
	if err != nil {
		return nil, code, err
	}

	// Save to cache. Won't return error.
	_ = c.cacher.Set(key, data)
	return data, http.StatusOK, nil
}

// GetMangaRecommendation to get manga recommendation list.
func (c *Cacher) GetMangaRecommendation(id int) (data []model.Recommendation, code int, err error) {
	// Get from cache.
	key := internal.GetKey(internal.KeyMangaRecommendation, id)
	if c.cacher.Get(key, &data) == nil {
		return data, http.StatusOK, nil
	}

	// Parse.
	data, code, err = c.api.GetMangaRecommendation(id)
	if err != nil {
		return nil, code, err
	}

	// Save to cache. Won't return error.
	_ = c.cacher.Set(key, data)
	return data, http.StatusOK, nil
}

// GetMangaStats to get manga stats list.
func (c *Cacher) GetMangaStats(id int) (data *model.Stats, code int, err error) {
	// Get from cache.
	key := internal.GetKey(internal.KeyMangaStats, id)
	if c.cacher.Get(key, &data) == nil {
		return data, http.StatusOK, nil
	}

	// Parse.
	data, code, err = c.api.GetMangaStats(id)
	if err != nil {
		return nil, code, err
	}

	// Save to cache. Won't return error.
	_ = c.cacher.Set(key, data)
	return data, http.StatusOK, nil
}

// GetMangaCharacter to get manga character list.
func (c *Cacher) GetMangaCharacter(id int) (data []model.Role, code int, err error) {
	// Get from cache.
	key := internal.GetKey(internal.KeyMangaCharacter, id)
	if c.cacher.Get(key, &data) == nil {
		return data, http.StatusOK, nil
	}

	// Parse.
	data, code, err = c.api.GetMangaCharacter(id)
	if err != nil {
		return nil, code, err
	}

	// Save to cache. Won't return error.
	_ = c.cacher.Set(key, data)
	return data, http.StatusOK, nil
}

// GetMangaNews to get manga news list.
func (c *Cacher) GetMangaNews(id int) (data []model.NewsItem, code int, err error) {
	// Get from cache.
	key := internal.GetKey(internal.KeyMangaNews, id)
	if c.cacher.Get(key, &data) == nil {
		return data, http.StatusOK, nil
	}

	// Parse.
	data, code, err = c.api.GetMangaNews(id)
	if err != nil {
		return nil, code, err
	}

	// Save to cache. Won't return error.
	_ = c.cacher.Set(key, data)
	return data, http.StatusOK, nil
}

// GetMangaArticle to get manga featured article list.
func (c *Cacher) GetMangaArticle(id int) (data []model.ArticleItem, code int, err error) {
	// Get from cache.
	key := internal.GetKey(internal.KeyMangaArticle, id)
	if c.cacher.Get(key, &data) == nil {
		return data, http.StatusOK, nil
	}

	// Parse.
	data, code, err = c.api.GetMangaArticle(id)
	if err != nil {
		return nil, code, err
	}

	// Save to cache. Won't return error.
	_ = c.cacher.Set(key, data)
	return data, http.StatusOK, nil
}

// GetMangaClub to get manga club list.
func (c *Cacher) GetMangaClub(id int) (data []model.ClubItem, code int, err error) {
	// Get from cache.
	key := internal.GetKey(internal.KeyMangaClub, id)
	if c.cacher.Get(key, &data) == nil {
		return data, http.StatusOK, nil
	}

	// Parse.
	data, code, err = c.api.GetMangaClub(id)
	if err != nil {
		return nil, code, err
	}

	// Save to cache. Won't return error.
	_ = c.cacher.Set(key, data)
	return data, http.StatusOK, nil
}

// GetMangaPicture to get manga picture list.
func (c *Cacher) GetMangaPicture(id int) (data []string, code int, err error) {
	// Get from cache.
	key := internal.GetKey(internal.KeyMangaPicture, id)
	if c.cacher.Get(key, &data) == nil {
		return data, http.StatusOK, nil
	}

	// Parse.
	data, code, err = c.api.GetMangaPicture(id)
	if err != nil {
		return nil, code, err
	}

	// Save to cache. Won't return error.
	_ = c.cacher.Set(key, data)
	return data, http.StatusOK, nil
}

// GetMangaMoreInfo to get manga more info.
func (c *Cacher) GetMangaMoreInfo(id int) (data string, code int, err error) {
	// Get from cache.
	key := internal.GetKey(internal.KeyMangaMoreInfo, id)
	if c.cacher.Get(key, &data) == nil {
		return data, http.StatusOK, nil
	}

	// Parse.
	data, code, err = c.api.GetMangaMoreInfo(id)
	if err != nil {
		return "", code, err
	}

	// Save to cache. Won't return error.
	_ = c.cacher.Set(key, data)
	return data, http.StatusOK, nil
}
