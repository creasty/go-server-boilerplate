package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
)

// CorsWrapper handles preflight requests for allowed origins
func CorsWrapper(corsAllowedOrigins []string) gin.HandlerFunc {
	mw := cors.New(cors.Options{
		AllowedOrigins:   corsAllowedOrigins,
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		Debug:            false,
	})

	return func(c *gin.Context) {
		if c.Request.Method != "OPTIONS" {
			c.Next()
			return
		}

		mw.HandlerFunc(c.Writer, c.Request)
		c.AbortWithStatus(http.StatusOK)
	}
}
