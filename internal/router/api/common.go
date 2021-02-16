package api

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/rl404/mal-db/internal/loader/api"
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
}

func (c *common) getEntryCount(w http.ResponseWriter, r *http.Request) {
	data, meta, code, err := c.api.GetEntryCount()
	utils.ResponseWithJSON(w, code, data, err, meta)
}

func (c *common) getYearSummary(w http.ResponseWriter, r *http.Request) {
	data, meta, code, err := c.api.GetYearSummary()
	utils.ResponseWithJSON(w, code, data, err, meta)
}

type enqueueRequest struct {
	Type string `json:"type"`
	ID   int    `json:"id"`
}

func (c *common) enqueue(w http.ResponseWriter, r *http.Request) {
	var request enqueueRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		utils.ResponseWithJSON(w, http.StatusBadRequest, nil, err, nil)
		return
	}
	code, err := c.api.Enqueue(request.Type, request.ID)
	utils.ResponseWithJSON(w, code, nil, err, nil)
}
