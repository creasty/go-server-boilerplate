package hbsvc

import (
	"github.com/gin-gonic/gin"
	"github.com/honeybadger-io/honeybadger-go"
)

// NotifyGinError does alomst the same as NotifyGinErrorWithFrames,
// except stack frames are automatically generated from the err itself
func NotifyGinError(err interface{}, body []byte, c *gin.Context) {
	frames := CreateFramesFromError(err)
	NotifyGinErrorWithFrames(err, body, c, frames)
}

// NotifyGinErrorWithFrames notifies the err and its arbitrary stack frames,
// along with a context of the certain http request
func NotifyGinErrorWithFrames(err interface{}, body []byte, c *gin.Context, frames []*honeybadger.Frame) {
	realStack := false

	hbErr := honeybadger.NewError(err)
	if frames != nil && len(frames) > 0 {
		hbErr.Stack = frames
		realStack = true
	}

	honeybadger.Notify(hbErr, honeybadger.Context{
		"ClientIP":  c.ClientIP(),
		"Params":    c.Params,
		"URL":       c.Request.URL,
		"Method":    c.Request.Method,
		"Body":      string(body[:]),
		"Headers":   c.Request.Header,
		"RealStack": realStack,
	})
}
