package server

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"github.com/creasty/go-server-boilerplate/model"
	"github.com/creasty/go-server-boilerplate/server/util"
)

// Timeout durations for reading a request body or writing a response body
const (
	ServerReadTimeout  = 120 * time.Second
	ServerWriteTimeout = 120 * time.Second
)

// Server holds the global contexts and clients that have any connections reused over the requests
type Server struct {
	Config   *model.Config
	SampleDB *gorm.DB
}

// Run initializes routings and serves the server
func (s *Server) Run() {
	if s.Config.IsProduction() {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	drawRoutes(s, r)

	httpServer := &http.Server{
		Addr:           s.Config.Host,
		Handler:        util.NewMethodOverrider(r),
		ReadTimeout:    ServerReadTimeout,
		WriteTimeout:   ServerWriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	if err := httpServer.ListenAndServe(); err != nil {
		panic(err)
	}
}
