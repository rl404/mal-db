package api

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/rl404/mal-db/internal/loader/api"
	"github.com/rl404/mal-db/internal/pkg/utils"
)

type people struct {
	api api.API
}

func registerPeople(r chi.Router, api api.API) {
	p := people{api: api}
	r.Get("/people/{id}", p.getPeople)
	r.Get("/people/{id}/va", p.getPeopleVA)
	r.Get("/people/{id}/staff", p.getPeopleStaff)
	r.Get("/people/{id}/manga", p.getPeopleManga)
}

// @summary Get people details
// @tags people
// @accept json
// @produce json
// @param id path integer true "People ID"
// @success 200 {object} utils.Response{data=model.People}
// @router /people/{id} [get]
func (p *people) getPeople(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	data, meta, code, err := p.api.GetPeople(id)
	utils.ResponseWithJSON(w, code, data, err, meta)
}

// @summary Get people character role list
// @tags people
// @accept json
// @produce json
// @param id path integer true "People ID"
// @param page query integer false "Page"
// @param limit query integer false "Limit"
// @success 200 {object} utils.Response{data=[]model.VoiceActor}
// @router /people/{id}/va [get]
func (p *people) getPeopleVA(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	page, _ := strconv.Atoi(utils.GetQuery(r, "page", "1"))
	limit, _ := strconv.Atoi(utils.GetQuery(r, "limit", "10"))
	data, meta, code, err := p.api.GetPeopleVA(id, page, limit)
	utils.ResponseWithJSON(w, code, data, err, meta)
}

// @summary Get people staff role list
// @tags people
// @accept json
// @produce json
// @param id path integer true "People ID"
// @param page query integer false "Page"
// @param limit query integer false "Limit"
// @success 200 {object} utils.Response{data=[]model.Role}
// @router /people/{id}/staff [get]
func (p *people) getPeopleStaff(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	page, _ := strconv.Atoi(utils.GetQuery(r, "page", "1"))
	limit, _ := strconv.Atoi(utils.GetQuery(r, "limit", "10"))
	data, meta, code, err := p.api.GetPeopleStaff(id, page, limit)
	utils.ResponseWithJSON(w, code, data, err, meta)
}

// @summary Get people published manga list
// @tags people
// @accept json
// @produce json
// @param id path integer true "People ID"
// @param page query integer false "Page"
// @param limit query integer false "Limit"
// @success 200 {object} utils.Response{data=[]model.Role}
// @router /people/{id}/manga [get]
func (p *people) getPeopleManga(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	page, _ := strconv.Atoi(utils.GetQuery(r, "page", "1"))
	limit, _ := strconv.Atoi(utils.GetQuery(r, "limit", "10"))
	data, meta, code, err := p.api.GetPeopleManga(id, page, limit)
	utils.ResponseWithJSON(w, code, data, err, meta)
}
