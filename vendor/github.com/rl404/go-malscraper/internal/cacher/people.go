package cacher

import (
	"net/http"

	"github.com/rl404/go-malscraper/internal"
	"github.com/rl404/go-malscraper/model"
)

// GetPeople to get people detail information.
func (c *Cacher) GetPeople(id int) (data *model.People, code int, err error) {
	// Get from cache.
	key := internal.GetKey(internal.KeyPeople, id)
	if c.cacher.Get(key, &data) == nil {
		return data, http.StatusOK, nil
	}

	// Parse.
	data, code, err = c.api.GetPeople(id)
	if err != nil {
		return nil, code, err
	}

	// Save to cache. Won't return error.
	_ = c.cacher.Set(key, data)
	return data, http.StatusOK, nil
}

// GetPeopleCharacter to get people anime character list.
func (c *Cacher) GetPeopleCharacter(id int) (data []model.PeopleCharacter, code int, err error) {
	// Get from cache.
	key := internal.GetKey(internal.KeyPeopleChar, id)
	if c.cacher.Get(key, &data) == nil {
		return data, http.StatusOK, nil
	}

	// Parse.
	data, code, err = c.api.GetPeopleCharacter(id)
	if err != nil {
		return nil, code, err
	}

	// Save to cache. Won't return error.
	_ = c.cacher.Set(key, data)
	return data, http.StatusOK, nil
}

// GetPeopleStaff to get people anime staff list.
func (c *Cacher) GetPeopleStaff(id int) (data []model.Role, code int, err error) {
	// Get from cache.
	key := internal.GetKey(internal.KeyPeopleStaff, id)
	if c.cacher.Get(key, &data) == nil {
		return data, http.StatusOK, nil
	}

	// Parse.
	data, code, err = c.api.GetPeopleStaff(id)
	if err != nil {
		return nil, code, err
	}

	// Save to cache. Won't return error.
	_ = c.cacher.Set(key, data)
	return data, http.StatusOK, nil
}

// GetPeopleManga to get people published manga list.
func (c *Cacher) GetPeopleManga(id int) (data []model.Role, code int, err error) {
	// Get from cache.
	key := internal.GetKey(internal.KeyPeopleManga, id)
	if c.cacher.Get(key, &data) == nil {
		return data, http.StatusOK, nil
	}

	// Parse.
	data, code, err = c.api.GetPeopleManga(id)
	if err != nil {
		return nil, code, err
	}

	// Save to cache. Won't return error.
	_ = c.cacher.Set(key, data)
	return data, http.StatusOK, nil
}

// GetPeopleNews to get people news list.
func (c *Cacher) GetPeopleNews(id int) (data []model.NewsItem, code int, err error) {
	// Get from cache.
	key := internal.GetKey(internal.KeyPeopleNews, id)
	if c.cacher.Get(key, &data) == nil {
		return data, http.StatusOK, nil
	}

	// Parse.
	data, code, err = c.api.GetPeopleNews(id)
	if err != nil {
		return nil, code, err
	}

	// Save to cache. Won't return error.
	_ = c.cacher.Set(key, data)
	return data, http.StatusOK, nil
}

// GetPeopleArticle to get people featured article list.
func (c *Cacher) GetPeopleArticle(id int) (data []model.ArticleItem, code int, err error) {
	// Get from cache.
	key := internal.GetKey(internal.KeyPeopleArticle, id)
	if c.cacher.Get(key, &data) == nil {
		return data, http.StatusOK, nil
	}

	// Parse.
	data, code, err = c.api.GetPeopleArticle(id)
	if err != nil {
		return nil, code, err
	}

	// Save to cache. Won't return error.
	_ = c.cacher.Set(key, data)
	return data, http.StatusOK, nil
}

// GetPeoplePicture to get people picture list.
func (c *Cacher) GetPeoplePicture(id int) (data []string, code int, err error) {
	// Get from cache.
	key := internal.GetKey(internal.KeyPeoplePicture, id)
	if c.cacher.Get(key, &data) == nil {
		return data, http.StatusOK, nil
	}

	// Parse.
	data, code, err = c.api.GetPeoplePicture(id)
	if err != nil {
		return nil, code, err
	}

	// Save to cache. Won't return error.
	_ = c.cacher.Set(key, data)
	return data, http.StatusOK, nil
}
