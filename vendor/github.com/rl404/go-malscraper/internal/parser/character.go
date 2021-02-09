package parser

import (
	"net/http"

	"github.com/rl404/go-malscraper/errors"
	"github.com/rl404/go-malscraper/model"
	"github.com/rl404/go-malscraper/pkg/utils"
)

// GetCharacter to get character details.
func (p *Parser) GetCharacter(id int) (*model.Character, int, error) {
	doc, code, err := p.getDoc(utils.BuildURL(malURL, "character", id), "#contentWrapper")
	if err != nil {
		return nil, code, err
	}
	d := p.character.GetDetails(doc)
	if d == nil {
		return nil, http.StatusNotFound, errors.ErrInvalidID
	}
	return d, http.StatusOK, nil
}

// GetCharacterArticle to get character featured article list.
func (p *Parser) GetCharacterArticle(id int) ([]model.ArticleItem, int, error) {
	doc, code, err := p.getDoc(utils.BuildURL(malURL, "character", id, "a", "featured"), "#content")
	if err != nil {
		return nil, code, err
	}
	d := p.character.GetArticle(doc)
	if d == nil {
		return nil, http.StatusNotFound, errors.ErrInvalidID
	}
	return d, http.StatusOK, nil
}

// GetCharacterOgraphy to get character animeography/mangaography list.
func (p *Parser) GetCharacterOgraphy(t string, id int) ([]model.Role, int, error) {
	doc, code, err := p.getDoc(utils.BuildURL(malURL, "character", id), "#contentWrapper")
	if err != nil {
		return nil, code, err
	}
	d := p.character.GetOgraphy(doc, t)
	if d == nil {
		return nil, http.StatusNotFound, errors.ErrInvalidID
	}
	return d, http.StatusOK, nil
}

// GetCharacterPicture to get character picture list.
func (p *Parser) GetCharacterPicture(id int) ([]string, int, error) {
	doc, code, err := p.getDoc(utils.BuildURL(malURL, "character", id, "a", "pictures"), "#content")
	if err != nil {
		return nil, code, err
	}
	return p.character.GetPictures(doc), http.StatusOK, nil
}

// GetCharacterClub to get character club list.
func (p *Parser) GetCharacterClub(id int) ([]model.ClubItem, int, error) {
	doc, code, err := p.getDoc(utils.BuildURL(malURL, "character", id, "a", "clubs"), "#content")
	if err != nil {
		return nil, code, err
	}
	d := p.character.GetClubs(doc)
	if d == nil {
		return nil, http.StatusNotFound, errors.ErrInvalidID
	}
	return d, http.StatusOK, nil
}

// GetCharacterVA to get character voice actor list.
func (p *Parser) GetCharacterVA(id int) ([]model.Role, int, error) {
	doc, code, err := p.getDoc(utils.BuildURL(malURL, "character", id), "#contentWrapper")
	if err != nil {
		return nil, code, err
	}
	d := p.character.GetVA(doc)
	if d == nil {
		return nil, http.StatusNotFound, errors.ErrInvalidID
	}
	return d, http.StatusOK, nil
}
