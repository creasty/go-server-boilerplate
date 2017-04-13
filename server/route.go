package server

import (
	"strconv"

	"github.com/creasty/gin-contrib/app_error"
	"github.com/creasty/gin-contrib/recovery"
	"github.com/creasty/panicsync"
	"github.com/gin-gonic/gin"
	"github.com/honeybadger-io/honeybadger-go"

	"github.com/creasty/go-server-boilerplate/server/middleware"
	"github.com/creasty/go-server-boilerplate/server/route"
	"github.com/creasty/go-server-boilerplate/service/hbsvc"
)

func drawRoutes(s *Server, r *gin.Engine) {
	// Middlewares
	r.Use(recovery.WrapWithCallback(func(c *gin.Context, body []byte, err interface{}) {
		hbsvc.NotifyGinError(err, body, c)
	}))
	r.Use(app_error.WrapWithCallback(func(c *gin.Context, body []byte, err error) {
		hbsvc.NotifyGinError(err, body, c)
	}))
	r.Use(middleware.SetConfigWrapper(s.Config))
	r.Use(middleware.SetSampleDBWrapper(s.SampleDB))

	// Routes
	drawAPIRoutes(s, r)
	drawSystemRoutes(s, r.Group("/system"), route.System{})
}

func drawAPIRoutes(s *Server, r *gin.Engine) {
	r.Use(middleware.CorsWrapper(s.Config.Cors.AllowedOrigins))

	r.GET("/ping", route.Ping)
}

func drawSystemRoutes(s *Server, r *gin.RouterGroup, system route.System) {
	r.Use(gin.BasicAuth(gin.Accounts{s.Config.BasicAuthUsername: s.Config.BasicAuthPassword}))

	r.GET("/appinfo", system.GetAppInfo)
}
