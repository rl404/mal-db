package api

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/rl404/mal-db/internal/constant"
	"github.com/rl404/mal-db/internal/loader/api"
	"github.com/rl404/mal-db/internal/pkg/utils"
)

type genre struct {
	api api.API
}

func registerGenre(r chi.Router, api api.API) {
	g := genre{api: api}
	r.Get("/genres/anime", g.getAnimeGenres)
	r.Get("/genres/manga", g.getMangaGenres)
}

// @summary Get anime genre list
// @tags genre
// @accept json
// @produce json
// @success 200 {object} utils.Response{data=[]model.Item}
// @router /genres/anime [get]
func (g *genre) getAnimeGenres(w http.ResponseWriter, r *http.Request) {
	data, meta, code, err := g.api.GetGenres(constant.AnimeType)
	utils.ResponseWithJSON(w, code, data, err, meta)
}

// @summary Get manga genre list
// @tags genre
// @accept json
// @produce json
// @success 200 {object} utils.Response{data=[]model.Item}
// @router /genres/manga [get]
func (g *genre) getMangaGenres(w http.ResponseWriter, r *http.Request) {
	data, meta, code, err := g.api.GetGenres(constant.MangaType)
	utils.ResponseWithJSON(w, code, data, err, meta)
}
