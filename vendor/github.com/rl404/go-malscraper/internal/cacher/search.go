package cacher

import (
	"net/http"

	"github.com/rl404/go-malscraper/internal"
	"github.com/rl404/go-malscraper/model"
)

// SearchAnime to search anime (no caching).
func (c *Cacher) SearchAnime(query model.Query) (data []model.AnimeSearch, code int, err error) {
	return c.api.SearchAnime(query)
}

// SearchManga to search manga (no caching).
func (c *Cacher) SearchManga(query model.Query) (data []model.MangaSearch, code int, err error) {
	return c.api.SearchManga(query)
}

// SearchCharacter to search character.
func (c *Cacher) SearchCharacter(name string, page int) (data []model.CharacterSearch, code int, err error) {
	// Get from cache.
	key := internal.GetKey(internal.KeySearchCharacter, name, page)
	if c.cacher.Get(key, &data) == nil {
		return data, http.StatusOK, nil
	}

	// Parse.
	data, code, err = c.api.SearchCharacter(name, page)
	if err != nil {
		return nil, code, err
	}

	// Save to cache. Won't return error.
	_ = c.cacher.Set(key, data)
	return data, http.StatusOK, nil
}

// SearchPeople to search people.
func (c *Cacher) SearchPeople(name string, page int) (data []model.PeopleSearch, code int, err error) {
	// Get from cache.
	key := internal.GetKey(internal.KeySearchPeople, name, page)
	if c.cacher.Get(key, &data) == nil {
		return data, http.StatusOK, nil
	}

	// Parse.
	data, code, err = c.api.SearchPeople(name, page)
	if err != nil {
		return nil, code, err
	}

	// Save to cache. Won't return error.
	_ = c.cacher.Set(key, data)
	return data, http.StatusOK, nil
}

// SearchClub to search club.
func (c *Cacher) SearchClub(query model.ClubQuery) (data []model.ClubSearch, code int, err error) {
	// Get from cache.
	key := internal.GetKey(internal.KeySearchClub, query.Name, query.Page, query.Category, query.Sort)
	if c.cacher.Get(key, &data) == nil {
		return data, http.StatusOK, nil
	}

	// Parse.
	data, code, err = c.api.SearchClub(query)
	if err != nil {
		return nil, code, err
	}

	// Save to cache. Won't return error.
	_ = c.cacher.Set(key, data)
	return data, http.StatusOK, nil
}

// SearchUser to search user.
func (c *Cacher) SearchUser(query model.UserQuery) (data []model.UserSearch, code int, err error) {
	// Get from cache.
	key := internal.GetKey(internal.KeySearchUser, query.Username, query.Page, query.Location, query.MinAge, query.MaxAge, query.Gender)
	if c.cacher.Get(key, &data) == nil {
		return data, http.StatusOK, nil
	}

	// Parse.
	data, code, err = c.api.SearchUser(query)
	if err != nil {
		return nil, code, err
	}

	// Save to cache. Won't return error.
	_ = c.cacher.Set(key, data)
	return data, http.StatusOK, nil
}
