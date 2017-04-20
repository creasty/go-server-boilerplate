package server

import (
	"github.com/creasty/gin-contrib/app_error"
	"github.com/creasty/gin-contrib/recovery"
	"github.com/gin-gonic/gin"

	"github.com/creasty/go-server-boilerplate/server/middleware"
	"github.com/creasty/go-server-boilerplate/server/route"
	system_route "github.com/creasty/go-server-boilerplate/server/route/system"
	hb_service "github.com/creasty/go-server-boilerplate/service/hb_service"
	"github.com/creasty/go-server-boilerplate/type/system"
)

func drawRoutes(r *gin.Engine, appContext *system.AppContext) {
	r.Use(recovery.WrapWithCallback(func(c *gin.Context, body []byte, err interface{}) {
		hb_service.NotifyGinError(err, body, c)
	}))
	r.Use(app_error.WrapWithCallback(func(c *gin.Context, body []byte, err error) {
		hb_service.NotifyGinError(err, body, c)
	}))
	r.Use(middleware.SetConfigWrapper(appContext.Config))
	r.Use(middleware.SetSampleDBWrapper(appContext.SampleDB))

	{
		r.GET("/ping", route.Ping)
	}

	{
		r := r.Group("/system")
		r.Use(gin.BasicAuth(gin.Accounts{
			appContext.Config.BasicAuthUsername: appContext.Config.BasicAuthPassword,
		}))

		r.GET("/appinfo", system_route.GetAppInfo)
	}
}
