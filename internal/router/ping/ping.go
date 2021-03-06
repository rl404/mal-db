package ping

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/rl404/mal-db/internal/pkg/utils"
)

// Ping contains basic routes.
type Ping struct {
}

// New to create new ping and other base routes.
func New() *Ping {
	return &Ping{}
}

// Register to register common routes.
func (p Ping) Register(r chi.Router) {
	r.Get("/", p.handleRoot)
	r.Get("/ping", p.handlePing)
	r.NotFound(http.HandlerFunc(p.handleNotFound))
	r.MethodNotAllowed(http.HandlerFunc(p.handleMethodNotAllowed))
}

func (p Ping) handleRoot(w http.ResponseWriter, _ *http.Request) {
	utils.ResponseWithJSON(w, http.StatusOK, "ok", nil)
}

func (p Ping) handlePing(w http.ResponseWriter, _ *http.Request) {
	utils.ResponseWithJSON(w, http.StatusOK, "pong", nil)
}

func (p Ping) handleNotFound(w http.ResponseWriter, _ *http.Request) {
	utils.ResponseWithJSON(w, http.StatusNotFound, nil, nil)
}

func (p Ping) handleMethodNotAllowed(w http.ResponseWriter, _ *http.Request) {
	utils.ResponseWithJSON(w, http.StatusMethodNotAllowed, nil, nil)
}
