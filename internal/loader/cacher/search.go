package cacher

import "github.com/rl404/mal-db/internal/model"

// SearchAnime to search anime.
func (c *Cacher) SearchAnime(q model.AnimeQuery) ([]model.Media, map[string]interface{}, int, error) {
	return c.api.SearchAnime(q)
}

// SearchManga to search manga.
func (c *Cacher) SearchManga(q model.MangaQuery) ([]model.Media, map[string]interface{}, int, error) {
	return c.api.SearchManga(q)
}

// SearchCharacter to search character.
func (c *Cacher) SearchCharacter(q model.EntryQuery) ([]model.Entry, map[string]interface{}, int, error) {
	return c.api.SearchCharacter(q)
}

// SearchPeople to search people.
func (c *Cacher) SearchPeople(q model.EntryQuery) ([]model.Entry, map[string]interface{}, int, error) {
	return c.api.SearchPeople(q)
}
