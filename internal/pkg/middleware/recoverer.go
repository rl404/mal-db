package middleware

import (
	"errors"
	"net/http"
	"runtime/debug"

	"github.com/go-chi/chi/middleware"
	"github.com/rl404/mal-db/internal/pkg/utils"
)

// Recoverer is custom recoverer middleware.
// Will return 500.
func Recoverer(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rvr := recover(); rvr != nil && rvr != http.ErrAbortHandler {

				logEntry := middleware.GetLogEntry(r)
				if logEntry != nil {
					logEntry.Panic(rvr, debug.Stack())
				} else {
					middleware.PrintPrettyStack(rvr)
				}

				utils.ResponseWithJSON(w, http.StatusInternalServerError, nil, errors.New("panic"))
			}
		}()

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
