package utils

import (
	"fmt"
	"net/url"
	"strings"
)

// BuildURL to build URL with its paths.
func BuildURL(host string, path ...interface{}) string {
	u := []string{host}
	for _, p := range path {
		u = append(u, fmt.Sprintf("%v", p))
	}
	return strings.Join(u, "/")
}

// BuildURLWithQuery to build URL with query params.
func BuildURLWithQuery(params map[string]interface{}, host string, path ...interface{}) string {
	u, _ := url.Parse(BuildURL(host, path...))
	q := u.Query()
	for k, v := range params {
		q.Set(k, fmt.Sprintf("%v", v))
	}
	u.RawQuery = q.Encode()
	return u.String()
}
