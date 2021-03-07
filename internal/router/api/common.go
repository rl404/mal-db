package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/rl404/mal-db/internal/loader/api"
	"github.com/rl404/mal-db/internal/model"
	"github.com/rl404/mal-db/internal/pkg/utils"
)

type common struct {
	api api.API
}

func registerCommon(r chi.Router, api api.API) {
	s := common{api: api}
	r.Get("/summary/total", s.getEntryCount)
	r.Get("/summary/year", s.getYearSummary)
	r.Post("/enqueue", s.enqueue)
	r.Get("/stats/history/{type}/{id}", s.getStatsHistory)
	r.Get("/compare/score", s.compareScore)
}

// @summary Get all entry count
// @tags common
// @accept json
// @produce json
// @success 200 {object} utils.Response{data=model.Total}
// @router /summary/total [get]
func (c *common) getEntryCount(w http.ResponseWriter, r *http.Request) {
	data, meta, code, err := c.api.GetEntryCount()
	utils.ResponseWithJSON(w, code, data, err, meta)
}

// @summary Get yearly summar count
// @tags common
// @accept json
// @produce json
// @success 200 {object} utils.Response{data=[]model.YearSummary}
// @router /summary/year [get]
func (c *common) getYearSummary(w http.ResponseWriter, r *http.Request) {
	data, meta, code, err := c.api.GetYearSummary()
	utils.ResponseWithJSON(w, code, data, err, meta)
}

type enqueueRequest struct {
	Type string `json:"type"`
	ID   int    `json:"id"`
}

// @summary Enqueue entry
// @tags common
// @accept json
// @produce json
// @param request body enqueueRequest true "entry type and id"
// @success 200 {object} utils.Response
// @router /enqueue [post]
func (c *common) enqueue(w http.ResponseWriter, r *http.Request) {
	var request enqueueRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		utils.ResponseWithJSON(w, http.StatusBadRequest, nil, err, nil)
		return
	}
	code, err := c.api.Enqueue(request.Type, request.ID)
	utils.ResponseWithJSON(w, code, nil, err, nil)
}

// @summary Entry stats history
// @tags common
// @accept json
// @produce json
// @param type path string true "Entry type" enums(anime,manga,character,people)
// @param id path integer true "Entry ID"
// @success 200 {object} utils.Response{data=[]model.StatsHistory}
// @router /stats/history/{type}/{id} [get]
func (c *common) getStatsHistory(w http.ResponseWriter, r *http.Request) {
	t := chi.URLParam(r, "type")
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	data, code, err := c.api.GetStatsHistory(t, id)
	utils.ResponseWithJSON(w, code, data, err, nil)
}

// @summary Entry stats history
// @tags common
// @accept json
// @produce json
// @param title query string false "Manga title"
// @param order query string false "Order (negative means descending)" Enums(title,-title,score,-score)
// @param page query integer false "Page"
// @param limit query integer false "Limit"
// @success 200 {object} utils.Response{data=[]model.ScoreComparison}
// @router /compare/score [get]
func (c *common) compareScore(w http.ResponseWriter, r *http.Request) {
	var query model.CompareQuery
	query.Title = r.URL.Query().Get("title")
	query.Order = r.URL.Query().Get("order")
	query.Page, _ = strconv.Atoi(utils.GetQuery(r, "page", "1"))
	query.Limit, _ = strconv.Atoi(utils.GetQuery(r, "limit", "20"))
	data, meta, code, err := c.api.CompareScore(query)
	utils.ResponseWithJSON(w, code, data, err, meta)
}
