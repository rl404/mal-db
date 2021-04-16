package api

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/rl404/mal-db/internal/constant"
	"github.com/rl404/mal-db/internal/loader"
	"github.com/rl404/mal-db/internal/pkg/utils"
)

type anime struct {
	api loader.API
}

func registerAnime(r chi.Router, api loader.API) {
	a := anime{api: api}
	r.Get("/anime/{id}", a.getAnime)
	r.Get("/anime/{id}/characters", a.getAnimeCharacter)
	r.Get("/anime/{id}/staff", a.getAnimeStaff)
	r.Get("/anime/{id}/stats", a.getAnimeStats)
}

// @summary Get anime details
// @tags anime
// @accept json
// @produce json
// @param id path integer true "Anime ID"
// @success 200 {object} utils.Response{data=model.Anime}
// @router /anime/{id} [get]
func (a *anime) getAnime(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	data, meta, code, err := a.api.GetAnime(id)
	utils.ResponseWithJSON(w, code, data, err, meta)
}

// @summary Get anime character list
// @tags anime
// @accept json
// @produce json
// @param id path integer true "Anime ID"
// @param page query integer false "Page"
// @param limit query integer false "Limit"
// @success 200 {object} utils.Response{data=[]model.AnimeCharacter}
// @router /anime/{id}/characters [get]
func (a *anime) getAnimeCharacter(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	page, _ := strconv.Atoi(utils.GetQuery(r, "page", "1"))
	limit, _ := strconv.Atoi(utils.GetQuery(r, "limit", "10"))
	data, meta, code, err := a.api.GetAnimeCharacter(id, page, limit)
	utils.ResponseWithJSON(w, code, data, err, meta)
}

// @summary Get anime staff list
// @tags anime
// @accept json
// @produce json
// @param id path integer true "Anime ID"
// @param page query integer false "Page"
// @param limit query integer false "Limit"
// @success 200 {object} utils.Response{data=[]model.Role}
// @router /anime/{id}/staff [get]
func (a *anime) getAnimeStaff(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	page, _ := strconv.Atoi(utils.GetQuery(r, "page", "1"))
	limit, _ := strconv.Atoi(utils.GetQuery(r, "limit", "10"))
	data, meta, code, err := a.api.GetAnimeStaff(id, page, limit)
	utils.ResponseWithJSON(w, code, data, err, meta)
}

// @summary Get anime stats
// @tags anime
// @accept json
// @produce json
// @param id path integer true "Anime ID"
// @success 200 {object} utils.Response{data=model.Stats}
// @router /anime/{id}/stats [get]
func (a *anime) getAnimeStats(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	data, meta, code, err := a.api.GetStats(constant.AnimeType, id)
	utils.ResponseWithJSON(w, code, data, err, meta)
}
