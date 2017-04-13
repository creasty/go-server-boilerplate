package middleware

import (
	"github.com/gin-gonic/gin"

	"github.com/creasty/go-server-boilerplate/model"
)

const configContextName = "Config"

// SetConfigWrapper sets model.Config to the context
func SetConfigWrapper(config *model.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(configContextName, config)
		c.Next()
	}
}

// GetConfig retrieve model.Config from the context
func GetConfig(c *gin.Context) *model.Config {
	v := c.MustGet(configContextName)

	cfg, ok := v.(*model.Config)
	if !ok {
		panic("Cannot retrive value from context")
	}

	return cfg
}
