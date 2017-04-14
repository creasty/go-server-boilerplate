package middleware

import (
	"github.com/gin-gonic/gin"

	"github.com/creasty/go-server-boilerplate/type/system"
)

const configContextName = "Config"

// SetConfigWrapper sets system.Config to the context
func SetConfigWrapper(config *system.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(configContextName, config)
		c.Next()
	}
}

// GetConfig retrieve system.Config from the context
func GetConfig(c *gin.Context) *system.Config {
	v := c.MustGet(configContextName)

	cfg, ok := v.(*system.Config)
	if !ok {
		panic("Cannot retrive value from context")
	}

	return cfg
}
