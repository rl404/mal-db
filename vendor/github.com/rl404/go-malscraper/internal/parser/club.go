package parser

import (
	"net/http"

	"github.com/rl404/go-malscraper/model"
	"github.com/rl404/go-malscraper/pkg/utils"
)

// GetClubs to get club list.
func (p *Parser) GetClubs(page int) ([]model.ClubSearch, int, error) {
	q := map[string]interface{}{"p": page}
	doc, code, err := p.getDoc(utils.BuildURLWithQuery(q, malURL, "clubs.php"), "#contentWrapper")
	if err != nil {
		return nil, code, err
	}
	return p.search.GetClub(doc), http.StatusOK, nil
}

// GetClub to get club detail information.
func (p *Parser) GetClub(id int) (*model.Club, int, error) {
	q := map[string]interface{}{"cid": id}
	doc, code, err := p.getDoc(utils.BuildURLWithQuery(q, malURL, "clubs.php"), "#contentWrapper")
	if err != nil {
		return nil, code, err
	}
	return p.club.GetDetails(doc), http.StatusOK, nil
}

// GetClubMember to get club member list.
func (p *Parser) GetClubMember(id int, page int) ([]model.ClubMember, int, error) {
	q := map[string]interface{}{"id": id, "action": "view", "t": "members", "show": 36 * (page - 1)}
	doc, code, err := p.getDoc(utils.BuildURLWithQuery(q, malURL, "clubs.php"), "#content")
	if err != nil {
		return nil, code, err
	}
	return p.club.GetMembers(doc), http.StatusOK, nil
}

// GetClubPicture to get club picture list.
func (p *Parser) GetClubPicture(id int) ([]string, int, error) {
	q := map[string]interface{}{"id": id, "action": "view", "t": "pictures"}
	doc, code, err := p.getDoc(utils.BuildURLWithQuery(q, malURL, "clubs.php"), "#content")
	if err != nil {
		return nil, code, err
	}
	return p.club.GetPictures(doc), http.StatusOK, nil
}

// GetClubRelated to get club related list.
func (p *Parser) GetClubRelated(id int) (*model.ClubRelated, int, error) {
	q := map[string]interface{}{"cid": id}
	doc, code, err := p.getDoc(utils.BuildURLWithQuery(q, malURL, "clubs.php"), "#contentWrapper")
	if err != nil {
		return nil, code, err
	}
	return p.club.GetRelated(doc), http.StatusOK, nil
}
