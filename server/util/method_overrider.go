package util

import (
	"net/http"
	"strings"
)

// MethodOverrider is a http.Handler for gin engine that supports method overriding
type MethodOverrider struct {
	handler http.Handler
}

// NewMethodOverrider creates a new http.Handler
func NewMethodOverrider(handler http.Handler) MethodOverrider {
	return MethodOverrider{handler: handler}
}

// ServeHTTP implements an interface of http.Handler.
// See https://golang.org/pkg/net/http/#Handler
func (m MethodOverrider) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if method := r.URL.Query().Get("_method"); method != "" {
		method = strings.ToUpper(method)

		// See https://golang.org/src/net/http/method.go
		if method == http.MethodGet ||
			method == http.MethodHead ||
			method == http.MethodPost ||
			method == http.MethodPut ||
			method == http.MethodPatch ||
			method == http.MethodDelete ||
			method == http.MethodConnect ||
			method == http.MethodOptions ||
			method == http.MethodTrace {
			r.Method = method
		}
	}

	m.handler.ServeHTTP(w, r)
}
