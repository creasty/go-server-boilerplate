package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping handles "ping-pong"
func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
