package server

import (
	"github.com/creasty/gin-contrib/app_error"
	"github.com/creasty/gin-contrib/recovery"
	"github.com/gin-gonic/gin"

	"github.com/creasty/go-server-boilerplate/server/middleware"
	"github.com/creasty/go-server-boilerplate/server/route"
	system_route "github.com/creasty/go-server-boilerplate/server/route"
	hb_svc "github.com/creasty/go-server-boilerplate/service/hb"
)

func drawRoutes(s *Server, r *gin.Engine) {
	// Middlewares
	r.Use(recovery.WrapWithCallback(func(c *gin.Context, body []byte, err interface{}) {
		hb_svc.NotifyGinError(err, body, c)
	}))
	r.Use(app_error.WrapWithCallback(func(c *gin.Context, body []byte, err error) {
		hb_svc.NotifyGinError(err, body, c)
	}))
	r.Use(middleware.SetConfigWrapper(s.Config))
	r.Use(middleware.SetSampleDBWrapper(s.SampleDB))

	// Routes
	drawAPIRoutes(s, r)
	drawSystemRoutes(s, r.Group("/system"))
}

func drawAPIRoutes(s *Server, r gin.IRouter) {
	r.GET("/ping", route.Ping)
}

func drawSystemRoutes(s *Server, r gin.IRouter) {
	r.Use(gin.BasicAuth(gin.Accounts{s.Config.BasicAuthUsername: s.Config.BasicAuthPassword}))

	r.GET("/appinfo", system_route.GetAppInfo)
}
