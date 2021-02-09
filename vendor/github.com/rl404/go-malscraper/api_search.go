package malscraper

import (
	"github.com/rl404/go-malscraper/model"
)

// SearchAnime to quick search anime.
//
// Example: https://myanimelist.net/anime.php?q=naruto.
func (m *Malscraper) SearchAnime(title string, page ...int) ([]model.AnimeSearch, int, error) {
	p := 1
	if len(page) > 0 {
		p = page[0]
	}
	return m.AdvSearchAnime(model.Query{Title: title, Page: p})
}

// AdvSearchAnime to search anime with advanced query.
//
// Available constant options.
//
//  Type          Status           Rating
//  -----------   --------------   -------------
//  TypeDefault   StatusDefault    RatingDefault
//  TypeTV        StatusOnGoing    RatingG
//  TypeOVA       StatusFinished   RatingPG
//  TypeMovie     StatusUpcoming   RatingPG13
//  TypeSpecial                    RatingR17
//  TypeONA                        RatingR
//  TypeMusic                      RatingRx
//
// ProducerID should be from `GetProducers()`.
//
// GenreID should be from `GetAnimeGenres()`.
//
// Example: https://myanimelist.net/anime.php?q=naruto.
func (m *Malscraper) AdvSearchAnime(query model.Query) ([]model.AnimeSearch, int, error) {
	if query.Page == 0 {
		query.Page = 1
	}
	return m.api.SearchAnime(query)
}

// SearchManga to quick search manga.
//
// Example: https://myanimelist.net/manga.php?q=naruto.
func (m *Malscraper) SearchManga(title string, page ...int) ([]model.MangaSearch, int, error) {
	p := 1
	if len(page) > 0 {
		p = page[0]
	}
	return m.AdvSearchManga(model.Query{Title: title, Page: p})
}

// AdvSearchManga to search manga with advanced query.
//
// Available constant options.
//
//  Type             Status
//  -----------      --------------
//  TypeDefault      StatusDefault
//  TypeManga        StatusOnGoing
//  TypeLightNovel   StatusFinished
//  TypeOneShot      StatusUpcoming
//  TypeDoujinshi    StatusHiatus
//  TypeManhwa       StatusDiscontinued
//  TypeManhua
//  TypeNovel
//
// MagazineID should be from `GetMagazines()`.
//
// GenreID should be from `GetMangaGenres()`.
//
// Example: https://myanimelist.net/manga.php?q=naruto.
func (m *Malscraper) AdvSearchManga(query model.Query) ([]model.MangaSearch, int, error) {
	if query.Page == 0 {
		query.Page = 1
	}
	return m.api.SearchManga(query)
}

// SearchCharacter to search character.
//
// Example: https://myanimelist.net/character.php?q=luffy.
func (m *Malscraper) SearchCharacter(name string, page ...int) ([]model.CharacterSearch, int, error) {
	p := 1
	if len(page) > 0 {
		p = page[0]
	}
	return m.api.SearchCharacter(name, p)
}

// SearchPeople to search people.
//
// Example: https://myanimelist.net/people.php?q=kana.
func (m *Malscraper) SearchPeople(name string, page ...int) ([]model.PeopleSearch, int, error) {
	p := 1
	if len(page) > 0 {
		p = page[0]
	}
	return m.api.SearchPeople(name, p)
}

// SearchClub to quick search club.
//
// Example: https://myanimelist.net/clubs.php?cat=club&catid=0&q=naruto&action=find.
func (m *Malscraper) SearchClub(name string, page ...int) ([]model.ClubSearch, int, error) {
	p := 1
	if len(page) > 0 {
		p = page[0]
	}
	return m.AdvSearchClub(model.ClubQuery{Name: name, Page: p})
}

// AdvSearchClub to search club with advanced query.
//
// Available constant options.
//
//  Category             Sort
//  ------------------   -----------
//  AllCategory          SortDefault
//  AnimeCategory        SortName
//  ConventionCategory   SortComment
//  ActorCategory        SortPost
//  CharacterCategory    SortMember
//  CompanyCategory
//  GameCategory
//  JapanCategory
//  CityCategory
//  MusicCategory
//  MangaCategory
//  SchoolCategory
//  OtherCategory
//
// Example: https://myanimelist.net/clubs.php?cat=club&catid=0&q=naruto&action=find.
func (m *Malscraper) AdvSearchClub(query model.ClubQuery) ([]model.ClubSearch, int, error) {
	if query.Page == 0 {
		query.Page = 1
	}
	return m.api.SearchClub(query)
}

// SearchUser to quick search user.
//
// Example: https://myanimelist.net/users.php?q=rl404.
func (m *Malscraper) SearchUser(username string, page ...int) ([]model.UserSearch, int, error) {
	p := 1
	if len(page) > 0 {
		p = page[0]
	}
	return m.AdvSearchUser(model.UserQuery{Username: username, Page: p})
}

// AdvSearchUser to search user with advanced query.
//
// Gender should be one of these constants.
//
//  GenderDefault
//  GenderMale
//  GenderFemale
//  GenderNonBinary
//
// Example: https://myanimelist.net/users.php?q=rl404.
func (m *Malscraper) AdvSearchUser(query model.UserQuery) ([]model.UserSearch, int, error) {
	if query.Page == 0 {
		query.Page = 1
	}
	return m.api.SearchUser(query)
}
