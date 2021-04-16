package api

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/rl404/mal-db/internal/constant"
	"github.com/rl404/mal-db/internal/loader"
	"github.com/rl404/mal-db/internal/pkg/utils"
)

type manga struct {
	api loader.API
}

func registerManga(r chi.Router, api loader.API) {
	m := manga{api: api}
	r.Get("/manga/{id}", m.getManga)
	r.Get("/manga/{id}/characters", m.getMangaCharacter)
	r.Get("/manga/{id}/stats", m.getMangaStats)
}

// @summary Get manga details
// @tags manga
// @accept json
// @produce json
// @param id path integer true "Manga ID"
// @success 200 {object} utils.Response{data=model.Manga}
// @router /manga/{id} [get]
func (m *manga) getManga(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	data, meta, code, err := m.api.GetManga(id)
	utils.ResponseWithJSON(w, code, data, err, meta)
}

// @summary Get manga character list
// @tags manga
// @accept json
// @produce json
// @param id path integer true "Manga ID"
// @param page query integer false "Page"
// @param limit query integer false "Limit"
// @success 200 {object} utils.Response{data=[]model.Role}
// @router /manga/{id}/characters [get]
func (m *manga) getMangaCharacter(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	page, _ := strconv.Atoi(utils.GetQuery(r, "page", "1"))
	limit, _ := strconv.Atoi(utils.GetQuery(r, "limit", "10"))
	data, meta, code, err := m.api.GetMangaCharacter(id, page, limit)
	utils.ResponseWithJSON(w, code, data, err, meta)
}

// @summary Get manga stats
// @tags manga
// @accept json
// @produce json
// @param id path integer true "Manga ID"
// @success 200 {object} utils.Response{data=model.Stats}
// @router /manga/{id}/stats [get]
func (m *manga) getMangaStats(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	data, meta, code, err := m.api.GetStats(constant.MangaType, id)
	utils.ResponseWithJSON(w, code, data, err, meta)
}
