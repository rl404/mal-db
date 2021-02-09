package parser

import (
	"net/http"

	"github.com/rl404/go-malscraper/model"
	"github.com/rl404/go-malscraper/pkg/utils"
)

// GetProducers to get anime producer/studio/licensor list.
func (p *Parser) GetProducers() ([]model.ItemCount, int, error) {
	doc, code, err := p.getDoc(utils.BuildURL(malURL, "anime", "producer"), ".anime-manga-search")
	if err != nil {
		return nil, code, err
	}
	return p.producer.GetProducers(doc), http.StatusOK, nil
}

// GetProducer to get producer anime list.
func (p *Parser) GetProducer(id int, page int) ([]model.AnimeItem, int, error) {
	q := map[string]interface{}{"page": page}
	doc, code, err := p.getDoc(utils.BuildURLWithQuery(q, malURL, "anime", "producer", id, "a"), "#content .js-categories-seasonal")
	if err != nil {
		return nil, code, err
	}
	return p.producer.GetAnime(doc), http.StatusOK, nil
}

// GetMagazines to get manga magazine/serialization list.
func (p *Parser) GetMagazines() ([]model.ItemCount, int, error) {
	doc, code, err := p.getDoc(utils.BuildURL(malURL, "manga", "magazine"), ".anime-manga-search")
	if err != nil {
		return nil, code, err
	}
	return p.producer.GetMagazines(doc), http.StatusOK, nil
}

// GetMagazine to get magazine manga list.
func (p *Parser) GetMagazine(id int, page int) ([]model.MangaItem, int, error) {
	q := map[string]interface{}{"page": page}
	doc, code, err := p.getDoc(utils.BuildURLWithQuery(q, malURL, "manga", "magazine", id, "a"), "#content .js-categories-seasonal")
	if err != nil {
		return nil, code, err
	}
	return p.producer.GetManga(doc), http.StatusOK, nil
}
