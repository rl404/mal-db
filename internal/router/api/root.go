package api

import (
	"github.com/go-chi/chi"
	"github.com/rl404/mal-db/internal/loader/api"
)

// API is api router.
type API struct {
	api api.API
}

// New to create new api router.
func New(api api.API) *API {
	return &API{
		api: api,
	}
}

// Register to register all api router endpoints.
func (a *API) Register(r chi.Router) {
	registerAnime(r, a.api)
	registerManga(r, a.api)
	registerCharacter(r, a.api)
	registerPeople(r, a.api)
	registerSummary(r, a.api)
	registerProducerMagazine(r, a.api)
	registerGenre(r, a.api)
	registerSearch(r, a.api)
}
