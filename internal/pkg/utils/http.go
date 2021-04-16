package utils

import (
	"bytes"
	"encoding/json"
	"image"
	"image/png"
	"net/http"
	"strconv"
	"strings"
)

// Response is standard api response model.
type Response struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Data    interface{}            `json:"data" swaggertype:"object"`
	Meta    map[string]interface{} `json:"meta" swaggertype:"object"`
}

// ResponseWithJSON to write response with JSON format.
func ResponseWithJSON(w http.ResponseWriter, code int, data interface{}, err error, meta ...map[string]interface{}) {
	r := Response{
		Status:  code,
		Message: strings.ToLower(http.StatusText(code)),
	}

	r.Data = data
	if err != nil {
		r.Data = err.Error()
	}

	// Handle meta.
	if len(meta) > 0 {
		r.Meta = meta[0]
	}

	rJSON, _ := json.Marshal(r)

	// Set response header.
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(rJSON)))
	w.WriteHeader(code)

	_, _ = w.Write(rJSON)
}

// GetQuery to get URL query with default value.
func GetQuery(r *http.Request, key string, defaultValue ...string) string {
	v := r.URL.Query().Get(key)
	if v != "" {
		return v
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return ""
}

// ResponseWithPNG to write response with PNG format.
func ResponseWithPNG(w http.ResponseWriter, image image.Image, code int, err error) {
	if err != nil {
		ResponseWithJSON(w, code, nil, err)
		return
	}

	buffer := new(bytes.Buffer)
	if err := png.Encode(buffer, image); err != nil {
		ResponseWithJSON(w, http.StatusInternalServerError, nil, err)
		return
	}

	// Set response header.
	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))
	w.WriteHeader(code)

	_, _ = w.Write(buffer.Bytes())
}
