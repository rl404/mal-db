package cacher

import (
	"net/http"

	"github.com/rl404/go-malscraper/internal"
	"github.com/rl404/go-malscraper/model"
)

// GetGenres to get anime/manga genre list.
func (c *Cacher) GetGenres(t string) (data []model.ItemCount, code int, err error) {
	// Get from cache.
	key := internal.GetKey(internal.KeyGenres, t)
	if c.cacher.Get(key, &data) == nil {
		return data, http.StatusOK, nil
	}

	// Parse.
	data, code, err = c.api.GetGenres(t)
	if err != nil {
		return nil, code, err
	}

	// Save to cache. Won't return error.
	_ = c.cacher.Set(key, data)
	return data, http.StatusOK, nil
}

// GetAnimeWithGenre to get anime list with specific genre.
func (c *Cacher) GetAnimeWithGenre(id int, page int) (data []model.AnimeItem, code int, err error) {
	// Get from cache.
	key := internal.GetKey(internal.KeyAnimeWithGenre, id, page)
	if c.cacher.Get(key, &data) == nil {
		return data, http.StatusOK, nil
	}

	// Parse.
	data, code, err = c.api.GetAnimeWithGenre(id, page)
	if err != nil {
		return nil, code, err
	}

	// Save to cache. Won't return error.
	_ = c.cacher.Set(key, data)
	return data, http.StatusOK, nil
}

// GetMangaWithGenre to get manga list with specific genre.
func (c *Cacher) GetMangaWithGenre(id int, page int) (data []model.MangaItem, code int, err error) {
	// Get from cache.
	key := internal.GetKey(internal.KeyMangaWithGenre, id, page)
	if c.cacher.Get(key, &data) == nil {
		return data, http.StatusOK, nil
	}

	// Parse.
	data, code, err = c.api.GetMangaWithGenre(id, page)
	if err != nil {
		return nil, code, err
	}

	// Save to cache. Won't return error.
	_ = c.cacher.Set(key, data)
	return data, http.StatusOK, nil
}
