package server

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/creasty/go-server-boilerplate/server/util"
	"github.com/creasty/go-server-boilerplate/type/system"
)

// Timeout durations for reading a request body or writing a response body
const (
	ServerReadTimeout  = 120 * time.Second
	ServerWriteTimeout = 120 * time.Second
)

// Run initializes routings and serves the server
func Run(appContext *system.AppContext) {
	if appContext.Config.IsProduction() {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	drawRoutes(r, appContext)

	httpServer := &http.Server{
		Addr:           appContext.Config.Host,
		Handler:        util.NewMethodOverrider(r),
		ReadTimeout:    ServerReadTimeout,
		WriteTimeout:   ServerWriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	if err := httpServer.ListenAndServe(); err != nil {
		panic(err)
	}
}
