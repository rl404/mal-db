package api

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/rl404/mal-db/internal/constant"
	"github.com/rl404/mal-db/internal/loader/api"
	"github.com/rl404/mal-db/internal/pkg/utils"
)

type producerMagazine struct {
	api api.API
}

func registerProducerMagazine(r chi.Router, api api.API) {
	pm := producerMagazine{api: api}
	r.Get("/producers", pm.getProducers)
	r.Get("/magazines", pm.getMagazines)
}

// @summary Get producer list
// @tags producer
// @accept json
// @produce json
// @success 200 {object} utils.Response{data=[]model.Item}
// @router /producers [get]
func (pm *producerMagazine) getProducers(w http.ResponseWriter, r *http.Request) {
	data, meta, code, err := pm.api.GetProducerMagazine(constant.AnimeType)
	utils.ResponseWithJSON(w, code, data, err, meta)
}

// @summary Get magazine list
// @tags magazine
// @accept json
// @produce json
// @success 200 {object} utils.Response{data=[]model.Item}
// @router /magazines [get]
func (pm *producerMagazine) getMagazines(w http.ResponseWriter, r *http.Request) {
	data, meta, code, err := pm.api.GetProducerMagazine(constant.MangaType)
	utils.ResponseWithJSON(w, code, data, err, meta)
}
