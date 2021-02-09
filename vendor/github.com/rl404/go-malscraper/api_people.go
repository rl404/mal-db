package malscraper

import (
	"github.com/rl404/go-malscraper/model"
)

// GetPeople to get people detail information.
//
// Example: https://myanimelist.net/people/1.
func (m *Malscraper) GetPeople(id int) (*model.People, int, error) {
	return m.api.GetPeople(id)
}

// GetPeopleCharacter to get people anime character list.
//
// Example: https://myanimelist.net/people/1.
func (m *Malscraper) GetPeopleCharacter(id int) ([]model.PeopleCharacter, int, error) {
	return m.api.GetPeopleCharacter(id)
}

// GetPeopleStaff to get people anime staff list.
//
// Example: https://myanimelist.net/people/1.
func (m *Malscraper) GetPeopleStaff(id int) ([]model.Role, int, error) {
	return m.api.GetPeopleStaff(id)
}

// GetPeopleManga to get people published manga list.
//
// Example: https://myanimelist.net/people/1868.
func (m *Malscraper) GetPeopleManga(id int) ([]model.Role, int, error) {
	return m.api.GetPeopleManga(id)
}

// GetPeopleNews to get people news list.
//
// Example: https://myanimelist.net/people/1/Tomokazu_Seki/news.
func (m *Malscraper) GetPeopleNews(id int) ([]model.NewsItem, int, error) {
	return m.api.GetPeopleNews(id)
}

// GetPeopleArticle to get people featured article list.
//
// Example: https://myanimelist.net/people/185/Kana_Hanazawa/featured.
func (m *Malscraper) GetPeopleArticle(id int) ([]model.ArticleItem, int, error) {
	return m.api.GetPeopleArticle(id)
}

// GetPeoplePicture to get people picture list.
//
// Example: https://myanimelist.net/people/1/Tomokazu_Seki/pictures.
func (m *Malscraper) GetPeoplePicture(id int) ([]string, int, error) {
	return m.api.GetPeoplePicture(id)
}
