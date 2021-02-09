package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/rl404/go-malscraper/service"
	"github.com/tidwall/gjson"
)

// ILogger is middleware logger interface.
type ILogger interface {
	Send(string, interface{}) error
}

// Logger is middleware to log request to elasticsearch.
func Logger(l service.Logger, mwl ILogger) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			t := time.Now()
			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

			var bw bodyWriter
			ww.Tee(&bw)

			scheme := "http"
			if r.TLS != nil {
				scheme = "https"
			}

			f := field{
				Method:    r.Method,
				URL:       fmt.Sprintf("%s://%s%s", scheme, r.Host, r.RequestURI),
				From:      r.RemoteAddr,
				CreatedAt: t,
			}

			defer func() {
				f.Code = ww.Status()
				f.Size = ww.BytesWritten()
				f.Time = time.Since(t).Truncate(time.Microsecond).Seconds()
				l.Debug("%s %s - %v %vb %s", f.Method, f.URL, f.Code, f.Size, time.Since(t).Truncate(time.Microsecond))

				// Send to elasticsearch.
				go func() {
					if f.Code/100 != 2 {
						f.Data = gjson.Get(string(bw.Body), "data").String()
					}
					if err := mwl.Send("mal-db-api", f); err != nil {
						l.Error(err.Error())
					}
				}()
			}()

			h.ServeHTTP(ww, r)
		}
		return http.HandlerFunc(fn)
	}
}

type field struct {
	Method    string    `json:"method"`
	URL       string    `json:"url"`
	Code      int       `json:"code"`
	Size      int       `json:"size"`
	From      string    `json:"from"`
	Time      float64   `json:"time"`
	Data      string    `json:"data"`
	CreatedAt time.Time `json:"created_at"`
}

type bodyWriter struct {
	Body []byte
}

func (w *bodyWriter) Write(b []byte) (int, error) {
	w.Body = b
	return len(b), nil
}
