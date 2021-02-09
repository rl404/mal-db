package parser

import (
	"net/http"

	"github.com/rl404/go-malscraper/model"
	"github.com/rl404/go-malscraper/pkg/utils"
)

// SearchAnime to search anime.
func (p *Parser) SearchAnime(query model.Query) ([]model.AnimeSearch, int, error) {
	doc, code, err := p.getDoc(utils.BuildURLWithQuery(p.queryToMap(query), malURL, "anime.php"), "div.js-categories-seasonal")
	if err != nil {
		return nil, code, err
	}
	return p.search.GetAnime(doc), http.StatusOK, nil
}

// SearchManga to search manga.
func (p *Parser) SearchManga(query model.Query) ([]model.MangaSearch, int, error) {
	doc, code, err := p.getDoc(utils.BuildURLWithQuery(p.queryToMap(query), malURL, "manga.php"), "div.js-categories-seasonal")
	if err != nil {
		return nil, code, err
	}
	return p.search.GetManga(doc), http.StatusOK, nil
}

// SearchCharacter to search character.
func (p *Parser) SearchCharacter(name string, page int) ([]model.CharacterSearch, int, error) {
	q := map[string]interface{}{"q": name, "show": 50 * (page - 1)}
	doc, code, err := p.getDoc(utils.BuildURLWithQuery(q, malURL, "character.php"), "#content")
	if err != nil {
		return nil, code, err
	}
	return p.search.GetCharacter(doc), http.StatusOK, nil
}

// SearchPeople to search people.
func (p *Parser) SearchPeople(name string, page int) ([]model.PeopleSearch, int, error) {
	q := map[string]interface{}{"q": name, "show": 50 * (page - 1)}
	doc, code, err := p.getDoc(utils.BuildURLWithQuery(q, malURL, "people.php"), "#content")
	if err != nil {
		return nil, code, err
	}
	return p.search.GetPeople(doc), http.StatusOK, nil
}

// SearchClub to search club.
func (p *Parser) SearchClub(query model.ClubQuery) ([]model.ClubSearch, int, error) {
	q := map[string]interface{}{
		"action": "find",
		"q":      query.Name,
		"p":      query.Page,
		"catid":  query.Category,
		"sort":   query.Sort,
	}
	doc, code, err := p.getDoc(utils.BuildURLWithQuery(q, malURL, "clubs.php"), "#content")
	if err != nil {
		return nil, code, err
	}
	return p.search.GetClub(doc), http.StatusOK, nil
}

// SearchUser to search user.
func (p *Parser) SearchUser(query model.UserQuery) ([]model.UserSearch, int, error) {
	q := map[string]interface{}{
		"q":       query.Username,
		"show":    24 * (query.Page - 1),
		"loc":     query.Location,
		"agelow":  query.MinAge,
		"agehigh": query.MaxAge,
		"g":       query.Gender,
	}
	doc, code, err := p.getDoc(utils.BuildURLWithQuery(q, malURL, "users.php"), "#content")
	if err != nil {
		return nil, code, err
	}
	return p.search.GetUser(doc), http.StatusOK, nil
}
