package parser

import (
	"net/http"

	"github.com/rl404/go-malscraper/model"
	"github.com/rl404/go-malscraper/pkg/utils"
)

// GetPeople to get people details.
func (p *Parser) GetPeople(id int) (*model.People, int, error) {
	doc, code, err := p.getDoc(utils.BuildURL(malURL, "people", id), "#contentWrapper")
	if err != nil {
		return nil, code, err
	}
	return p.people.GetDetails(doc), http.StatusOK, nil
}

// GetPeopleCharacter to get people anime character list.
func (p *Parser) GetPeopleCharacter(id int) ([]model.PeopleCharacter, int, error) {
	doc, code, err := p.getDoc(utils.BuildURL(malURL, "people", id), "#contentWrapper")
	if err != nil {
		return nil, code, err
	}
	return p.people.GetCharacters(doc), http.StatusOK, nil
}

// GetPeopleStaff to get people anime staff list.
func (p *Parser) GetPeopleStaff(id int) ([]model.Role, int, error) {
	doc, code, err := p.getDoc(utils.BuildURL(malURL, "people", id), "#contentWrapper")
	if err != nil {
		return nil, code, err
	}
	return p.people.GetStaffManga(doc, "staff"), http.StatusOK, nil
}

// GetPeopleManga to get people published manga list.
func (p *Parser) GetPeopleManga(id int) ([]model.Role, int, error) {
	doc, code, err := p.getDoc(utils.BuildURL(malURL, "people", id), "#contentWrapper")
	if err != nil {
		return nil, code, err
	}
	return p.people.GetStaffManga(doc, "manga"), http.StatusOK, nil
}

// GetPeopleNews to get people news list.
func (p *Parser) GetPeopleNews(id int) ([]model.NewsItem, int, error) {
	doc, code, err := p.getDoc(utils.BuildURL(malURL, "people", id, "a", "news"), "#content table tr td")
	if err != nil {
		return nil, code, err
	}
	return p.people.GetNews(doc), http.StatusOK, nil
}

// GetPeopleArticle to get people featured article list.
func (p *Parser) GetPeopleArticle(id int) ([]model.ArticleItem, int, error) {
	doc, code, err := p.getDoc(utils.BuildURL(malURL, "people", id, "a", "featured"), ".news-list")
	if err != nil {
		return nil, code, err
	}
	return p.people.GetArticle(doc), http.StatusOK, nil
}

// GetPeoplePicture to get people picture list.
func (p *Parser) GetPeoplePicture(id int) ([]string, int, error) {
	doc, code, err := p.getDoc(utils.BuildURL(malURL, "people", id, "a", "pictures"), "#content table tr td")
	if err != nil {
		return nil, code, err
	}
	return p.people.GetPictures(doc), http.StatusOK, nil
}
