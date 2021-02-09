package parser

import (
	"net/http"

	"github.com/rl404/go-malscraper/model"
	"github.com/rl404/go-malscraper/pkg/utils"
)

var topAnimeTypes = []string{"", "airing", "upcoming", "tv", "movie", "ova", "ona", "special", "bypopularity", "favorite"}
var topMangaTypes = []string{"", "manga", "novels", "oneshots", "doujin", "manhwa", "manhua", "bypopularity", "favorite"}

// GetTopAnime to get top anime list.
func (p *Parser) GetTopAnime(t int, page int) ([]model.TopAnime, int, error) {
	q := map[string]interface{}{"type": topAnimeTypes[t], "limit": 50 * (page - 1)}
	doc, code, err := p.getDoc(utils.BuildURLWithQuery(q, malURL, "topanime.php"), "#content")
	if err != nil {
		return nil, code, err
	}
	return p.top.GetAnime(doc), http.StatusOK, nil
}

// GetTopManga to get top manga list.
func (p *Parser) GetTopManga(t int, page int) ([]model.TopManga, int, error) {
	q := map[string]interface{}{"type": topMangaTypes[t], "limit": 50 * (page - 1)}
	doc, code, err := p.getDoc(utils.BuildURLWithQuery(q, malURL, "topmanga.php"), "#content")
	if err != nil {
		return nil, code, err
	}
	return p.top.GetManga(doc), http.StatusOK, nil
}

// GetTopCharacter to get top character list.
func (p *Parser) GetTopCharacter(page int) ([]model.TopCharacter, int, error) {
	q := map[string]interface{}{"limit": 50 * (page - 1)}
	doc, code, err := p.getDoc(utils.BuildURLWithQuery(q, malURL, "character.php"), "#content")
	if err != nil {
		return nil, code, err
	}
	return p.top.GetCharacter(doc), http.StatusOK, nil
}

// GetTopPeople to get top people list.
func (p *Parser) GetTopPeople(page int) ([]model.TopPeople, int, error) {
	q := map[string]interface{}{"limit": 50 * (page - 1)}
	doc, code, err := p.getDoc(utils.BuildURLWithQuery(q, malURL, "people.php"), "#content")
	if err != nil {
		return nil, code, err
	}
	return p.top.GetPeople(doc), http.StatusOK, nil
}
