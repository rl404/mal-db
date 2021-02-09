package parser

import (
	"net/http"

	"github.com/rl404/go-malscraper/model"
	"github.com/rl404/go-malscraper/pkg/utils"
)

// GetGenres to get anime/manga genre list.
func (p *Parser) GetGenres(t string) ([]model.ItemCount, int, error) {
	doc, code, err := p.getDoc(utils.BuildURL(malURL, t+".php"), ".anime-manga-search .genre-link")
	if err != nil {
		return nil, code, err
	}
	return p.genre.GetGenres(doc), http.StatusOK, nil
}

// GetAnimeWithGenre to get anime list with specific genre.
func (p *Parser) GetAnimeWithGenre(id int, page int) ([]model.AnimeItem, int, error) {
	q := map[string]interface{}{"page": page}
	doc, code, err := p.getDoc(utils.BuildURLWithQuery(q, malURL, "anime", "genre", id, "a"), "#contentWrapper")
	if err != nil {
		return nil, code, err
	}
	return p.producer.GetAnime(doc), http.StatusOK, nil
}

// GetMangaWithGenre to get manga list with specific genre.
func (p *Parser) GetMangaWithGenre(id int, page int) ([]model.MangaItem, int, error) {
	q := map[string]interface{}{"page": page}
	doc, code, err := p.getDoc(utils.BuildURLWithQuery(q, malURL, "manga", "genre", id, "a"), "#contentWrapper")
	if err != nil {
		return nil, code, err
	}
	return p.producer.GetManga(doc), http.StatusOK, nil
}
