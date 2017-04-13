package util

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// MethodOverrider is a http.Handler for gin engine that supports method overriding
type MethodOverrider struct {
	engine *gin.Engine
}

// NewMethodOverrider creates a new http.Handler
func NewMethodOverrider(engine *gin.Engine) MethodOverrider {
	return MethodOverrider{
		engine: engine,
	}
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

	m.engine.ServeHTTP(w, r)
}
