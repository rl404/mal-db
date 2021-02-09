package parser

import (
	"net/http"

	"github.com/rl404/go-malscraper/model"
	"github.com/rl404/go-malscraper/pkg/utils"
)

// GetSeason to get seasonal anime list.
func (p *Parser) GetSeason(season string, year int) ([]model.AnimeItem, int, error) {
	doc, code, err := p.getDoc(utils.BuildURL(malURL, "anime", "season", year, season), "#content")
	if err != nil {
		return nil, code, err
	}
	return p.producer.GetAnime(doc), http.StatusOK, nil
}
