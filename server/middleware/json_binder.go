package middleware

import (
	"encoding/json"
	"io"

	"github.com/gin-gonic/gin"
)

// BindJSON binds a json reuqest to the obj
func BindJSON(c *gin.Context, obj interface{}) error {
	decoder := json.NewDecoder(c.Request.Body)
	err := decoder.Decode(obj)
	if err == io.EOF { // ignore EOF
		err = nil
	}
	return err
}
