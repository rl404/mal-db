package malscraper

import (
	"github.com/rl404/go-malscraper/model"
)

// GetTopAnime to get top anime list.
//
// Type should be one of these constants.
//
//  TopDefault
//  TopAiring
//  TopUpcoming
//  TopTV
//  TopMovie
//  TopOVA
//  TopONA
//  TopSpecial
//  TopPopularAnime
//  TopFavoriteAnime
//
// Example: https://myanimelist.net/topanime.php.
func (m *Malscraper) GetTopAnime(typePage ...int) ([]model.TopAnime, int, error) {
	t, p := 0, 1
	for i, param := range typePage {
		switch i {
		case 0:
			t = param
		case 1:
			p = param
		}
	}
	return m.api.GetTopAnime(t, p)
}

// GetTopManga to get top manga list.
//
// Type should be one of these constants.
//
//  TopDefault
//  TopManga
//  TopNovel
//  TopOneshot
//  TopDoujin
//  TopManhwa
//  TopManhua
//  TopPopularManga
//  TopFavoriteManga
//
// Example: https://mymangalist.net/topmanga.php.
func (m *Malscraper) GetTopManga(typePage ...int) ([]model.TopManga, int, error) {
	t, p := 0, 1
	for i, param := range typePage {
		switch i {
		case 0:
			t = param
		case 1:
			p = param
		}
	}
	return m.api.GetTopManga(t, p)
}

// GetTopCharacter to get top character list.
//
// Example: https://myanimelist.net/character.php.
func (m *Malscraper) GetTopCharacter(page ...int) ([]model.TopCharacter, int, error) {
	p := 1
	if len(page) > 0 {
		p = page[0]
	}
	return m.api.GetTopCharacter(p)
}

// GetTopPeople to get top people list.
//
// Example: https://myanimelist.net/people.php.
func (m *Malscraper) GetTopPeople(page ...int) ([]model.TopPeople, int, error) {
	p := 1
	if len(page) > 0 {
		p = page[0]
	}
	return m.api.GetTopPeople(p)
}
