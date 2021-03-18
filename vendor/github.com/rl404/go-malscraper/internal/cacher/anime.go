package cacher

import (
	"net/http"

	"github.com/rl404/go-malscraper/internal"
	"github.com/rl404/go-malscraper/model"
)

// GetAnime to get anime from cache.
func (c *Cacher) GetAnime(id int) (data *model.Anime, code int, err error) {
	// Get from cache.
	key := internal.GetKey(internal.KeyAnime, id)
	if c.cacher.Get(key, &data) == nil {
		return data, http.StatusOK, nil
	}

	// Parse.
	data, code, err = c.api.GetAnime(id)
	if err != nil {
		return nil, code, err
	}

	// Save to cache. Won't return error.
	_ = c.cacher.Set(key, data)
	return data, http.StatusOK, nil
}

// GetAnimeCharacter to get anime character list.
func (c *Cacher) GetAnimeCharacter(id int) (data []model.CharacterItem, code int, err error) {
	// Get from cache.
	key := internal.GetKey(internal.KeyAnimeCharacter, id)
	if c.cacher.Get(key, &data) == nil {
		return data, http.StatusOK, nil
	}

	// Parse.
	data, code, err = c.api.GetAnimeCharacter(id)
	if err != nil {
		return nil, code, err
	}

	// Save to cache. Won't return error.
	_ = c.cacher.Set(key, data)
	return data, http.StatusOK, nil
}

// GetAnimeStaff to get anime staff list.
func (c *Cacher) GetAnimeStaff(id int) (data []model.Role, code int, err error) {
	// Get from cache.
	key := internal.GetKey(internal.KeyAnimeStaff, id)
	if c.cacher.Get(key, &data) == nil {
		return data, http.StatusOK, nil
	}

	// Parse.
	data, code, err = c.api.GetAnimeStaff(id)
	if err != nil {
		return nil, code, err
	}

	// Save to cache. Won't return error.
	_ = c.cacher.Set(key, data)
	return data, http.StatusOK, nil
}

// GetAnimeVideo to get anime video list.
func (c *Cacher) GetAnimeVideo(id int, page int) (data *model.Video, code int, err error) {
	// Get from cache.
	key := internal.GetKey(internal.KeyAnimeVideo, id, page)
	if c.cacher.Get(key, &data) == nil {
		return data, http.StatusOK, nil
	}

	// Parse.
	data, code, err = c.api.GetAnimeVideo(id, page)
	if err != nil {
		return nil, code, err
	}

	// Save to cache. Won't return error.
	_ = c.cacher.Set(key, data)
	return data, http.StatusOK, nil
}

// GetAnimeEpisode to get anime episode list.
func (c *Cacher) GetAnimeEpisode(id int, page int) (data []model.Episode, code int, err error) {
	// Get from cache.
	key := internal.GetKey(internal.KeyAnimeEpisode, id, page)
	if c.cacher.Get(key, &data) == nil {
		return data, http.StatusOK, nil
	}

	// Parse.
	data, code, err = c.api.GetAnimeEpisode(id, page)
	if err != nil {
		return nil, code, err
	}

	// Save to cache. Won't return error.
	_ = c.cacher.Set(key, data)
	return data, http.StatusOK, nil
}

// GetAnimeStats to get anime stats.
func (c *Cacher) GetAnimeStats(id int) (data *model.Stats, code int, err error) {
	// Get from cache.
	key := internal.GetKey(internal.KeyAnimeStats, id)
	if c.cacher.Get(key, &data) == nil {
		return data, http.StatusOK, nil
	}

	// Parse.
	data, code, err = c.api.GetAnimeStats(id)
	if err != nil {
		return nil, code, err
	}

	// Save to cache. Won't return error.
	_ = c.cacher.Set(key, data)
	return data, http.StatusOK, nil
}

// GetAnimeReview to get anime review list.
func (c *Cacher) GetAnimeReview(id int, page int) (data []model.Review, code int, err error) {
	// Get from cache.
	key := internal.GetKey(internal.KeyAnimeReview, id, page)
	if c.cacher.Get(key, &data) == nil {
		return data, http.StatusOK, nil
	}

	// Parse.
	data, code, err = c.api.GetAnimeReview(id, page)
	if err != nil {
		return nil, code, err
	}

	// Save to cache. Won't return error.
	_ = c.cacher.Set(key, data)
	return data, http.StatusOK, nil
}

// GetAnimeRecommendation to get anime recommendation list.
func (c *Cacher) GetAnimeRecommendation(id int) (data []model.Recommendation, code int, err error) {
	// Get from cache.
	key := internal.GetKey(internal.KeyAnimeRecommendation, id)
	if c.cacher.Get(key, &data) == nil {
		return data, http.StatusOK, nil
	}

	// Parse.
	data, code, err = c.api.GetAnimeRecommendation(id)
	if err != nil {
		return nil, code, err
	}

	// Save to cache. Won't return error.
	_ = c.cacher.Set(key, data)
	return data, http.StatusOK, nil
}

// GetAnimeNews to get anime recommendation list.
func (c *Cacher) GetAnimeNews(id int) (data []model.NewsItem, code int, err error) {
	// Get from cache.
	key := internal.GetKey(internal.KeyAnimeNews, id)
	if c.cacher.Get(key, &data) == nil {
		return data, http.StatusOK, nil
	}

	// Parse.
	data, code, err = c.api.GetAnimeNews(id)
	if err != nil {
		return nil, code, err
	}

	// Save to cache. Won't return error.
	_ = c.cacher.Set(key, data)
	return data, http.StatusOK, nil
}

// GetAnimeArticle to get anime featured article list.
func (c *Cacher) GetAnimeArticle(id int) (data []model.ArticleItem, code int, err error) {
	// Get from cache.
	key := internal.GetKey(internal.KeyAnimeArticle, id)
	if c.cacher.Get(key, &data) == nil {
		return data, http.StatusOK, nil
	}

	// Parse.
	data, code, err = c.api.GetAnimeArticle(id)
	if err != nil {
		return nil, code, err
	}

	// Save to cache. Won't return error.
	_ = c.cacher.Set(key, data)
	return data, http.StatusOK, nil
}

// GetAnimeClub to get anime club list.
func (c *Cacher) GetAnimeClub(id int) (data []model.ClubItem, code int, err error) {
	// Get from cache.
	key := internal.GetKey(internal.KeyAnimeClub, id)
	if c.cacher.Get(key, &data) == nil {
		return data, http.StatusOK, nil
	}

	// Parse.
	data, code, err = c.api.GetAnimeClub(id)
	if err != nil {
		return nil, code, err
	}

	// Save to cache. Won't return error.
	_ = c.cacher.Set(key, data)
	return data, http.StatusOK, nil
}

// GetAnimePicture to get anime picture list.
func (c *Cacher) GetAnimePicture(id int) (data []string, code int, err error) {
	// Get from cache.
	key := internal.GetKey(internal.KeyAnimePicture, id)
	if c.cacher.Get(key, &data) == nil {
		return data, http.StatusOK, nil
	}

	// Parse.
	data, code, err = c.api.GetAnimePicture(id)
	if err != nil {
		return nil, code, err
	}

	// Save to cache. Won't return error.
	_ = c.cacher.Set(key, data)
	return data, http.StatusOK, nil
}

// GetAnimeMoreInfo to get anime more info.
func (c *Cacher) GetAnimeMoreInfo(id int) (data string, code int, err error) {
	// Get from cache.
	key := internal.GetKey(internal.KeyAnimeMoreInfo, id)
	if c.cacher.Get(key, &data) == nil {
		return data, http.StatusOK, nil
	}

	// Parse.
	data, code, err = c.api.GetAnimeMoreInfo(id)
	if err != nil {
		return "", code, err
	}

	// Save to cache. Won't return error.
	_ = c.cacher.Set(key, data)
	return data, http.StatusOK, nil
}
