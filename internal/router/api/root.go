package api

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/rl404/mal-db/internal/loader"
)

// API is api router.
type API struct {
	api loader.API
}

// New to create new api router.
func New(api loader.API) *API {
	return &API{
		api: api,
	}
}

// Register to register all api router endpoints.
func (a *API) Register(r chi.Router, mw func(http.Handler) http.Handler) {
	r2 := chi.NewRouter()
	r2.Use(mw)
	registerAnime(r2, a.api)
	registerManga(r2, a.api)
	registerCharacter(r2, a.api)
	registerPeople(r2, a.api)
	registerProducerMagazine(r2, a.api)
	registerGenre(r2, a.api)
	registerSearch(r2, a.api)
	registerCommon(r2, a.api)
	r.Mount("/", r2)
}
