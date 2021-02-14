package api

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/rl404/mal-db/internal/loader/api"
	"github.com/rl404/mal-db/internal/pkg/utils"
)

type summary struct {
	api api.API
}

func registerSummary(r chi.Router, api api.API) {
	s := summary{api: api}
	r.Get("/summary/total", s.getEntryCount)
	r.Get("/summary/year", s.getYearSummary)
}

func (s *summary) getEntryCount(w http.ResponseWriter, r *http.Request) {
	data, meta, code, err := s.api.GetEntryCount()
	utils.ResponseWithJSON(w, code, data, err, meta)
}

func (s *summary) getYearSummary(w http.ResponseWriter, r *http.Request) {
	data, meta, code, err := s.api.GetYearSummary()
	utils.ResponseWithJSON(w, code, data, err, meta)
}
