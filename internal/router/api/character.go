package api

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/rl404/mal-db/internal/constant"
	"github.com/rl404/mal-db/internal/loader"
	"github.com/rl404/mal-db/internal/pkg/utils"
)

type character struct {
	api loader.API
}

func registerCharacter(r chi.Router, api loader.API) {
	c := character{api: api}
	r.Get("/character/{id}", c.getCharacter)
	r.Get("/character/{id}/anime", c.getCharacterAnime)
	r.Get("/character/{id}/manga", c.getCharacterManga)
	r.Get("/character/{id}/va", c.getCharacterVA)
}

// @summary Get character details
// @tags character
// @accept json
// @produce json
// @param id path integer true "Character ID"
// @success 200 {object} utils.Response{data=model.Character}
// @router /character/{id} [get]
func (c *character) getCharacter(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	data, meta, code, err := c.api.GetCharacter(id)
	utils.ResponseWithJSON(w, code, data, err, meta)
}

// @summary Get character animeography list
// @tags character
// @accept json
// @produce json
// @param id path integer true "Character ID"
// @param page query integer false "Page"
// @param limit query integer false "Limit"
// @success 200 {object} utils.Response{data=[]model.Role}
// @router /character/{id}/anime [get]
func (c *character) getCharacterAnime(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	page, _ := strconv.Atoi(utils.GetQuery(r, "page", "1"))
	limit, _ := strconv.Atoi(utils.GetQuery(r, "limit", "10"))
	data, meta, code, err := c.api.GetCharacterOgraphy(id, constant.AnimeType, page, limit)
	utils.ResponseWithJSON(w, code, data, err, meta)
}

// @summary Get character mangaography list
// @tags character
// @accept json
// @produce json
// @param id path integer true "Character ID"
// @param page query integer false "Page"
// @param limit query integer false "Limit"
// @success 200 {object} utils.Response{data=[]model.Role}
// @router /character/{id}/manga [get]
func (c *character) getCharacterManga(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	page, _ := strconv.Atoi(utils.GetQuery(r, "page", "1"))
	limit, _ := strconv.Atoi(utils.GetQuery(r, "limit", "10"))
	data, meta, code, err := c.api.GetCharacterOgraphy(id, constant.MangaType, page, limit)
	utils.ResponseWithJSON(w, code, data, err, meta)
}

// @summary Get character voice actor list
// @tags character
// @accept json
// @produce json
// @param id path integer true "Character ID"
// @param page query integer false "Page"
// @param limit query integer false "Limit"
// @success 200 {object} utils.Response{data=[]model.Role}
// @router /character/{id}/va [get]
func (c *character) getCharacterVA(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	page, _ := strconv.Atoi(utils.GetQuery(r, "page", "1"))
	limit, _ := strconv.Atoi(utils.GetQuery(r, "limit", "10"))
	data, meta, code, err := c.api.GetCharacterVA(id, page, limit)
	utils.ResponseWithJSON(w, code, data, err, meta)
}
