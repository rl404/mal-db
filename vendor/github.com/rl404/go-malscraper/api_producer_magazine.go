package malscraper

import (
	"github.com/rl404/go-malscraper/model"
)

// GetProducers to get anime producer/studio/licensor list.
//
// Example: https://myanimelist.net/anime/producer.
func (m *Malscraper) GetProducers() ([]model.ItemCount, int, error) {
	return m.api.GetProducers()
}

// GetProducer to get producer anime list.
//
// Example: https://myanimelist.net/anime/producer/1/Studio_Pierrot.
func (m *Malscraper) GetProducer(id int, page ...int) ([]model.AnimeItem, int, error) {
	p := 1
	if len(page) > 0 {
		p = page[0]
	}
	return m.api.GetProducer(id, p)
}

// GetMagazines to get manga magazine/serialization list.
//
// Example: https://myanimelist.net/manga/magazine.
func (m *Malscraper) GetMagazines() ([]model.ItemCount, int, error) {
	return m.api.GetMagazines()
}

// GetMagazine to get magazine manga list.
//
// Example: https://myanimelist.net/manga/magazine/1/Big_Comic_Original.
func (m *Malscraper) GetMagazine(id int, page ...int) ([]model.MangaItem, int, error) {
	p := 1
	if len(page) > 0 {
		p = page[0]
	}
	return m.api.GetMagazine(id, p)
}
