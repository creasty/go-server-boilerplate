package middleware

import (
	"github.com/gin-gonic/gin"

	"github.com/creasty/go-server-boilerplate/type/system"
)

const (
	appContextContextName = "AppContext"
)

// SetAppContextWrapper sets an AppContext object to the context
func SetAppContextWrapper(appContext *system.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(appContextContextName, appContext)
		c.Next()
	}
}

// GetAppContext retrives an AppContext object from the context
func GetAppContext(c *gin.Context) *system.AppContext {
	v := c.MustGet(appContextContextName)

	appContext, ok := v.(*system.AppContext)
	if !ok {
		panic("Cannot retrive value from context")
	}

	return appContext
}
