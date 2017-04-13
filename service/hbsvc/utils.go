package hbsvc

import (
	"strconv"

	"github.com/go-errors/errors"
	"github.com/honeybadger-io/honeybadger-go"
)

// CreateFramesFromError builds stack frames from the error object
func CreateFramesFromError(err interface{}) []*honeybadger.Frame {
	frames := make([]*honeybadger.Frame, 0)

	if goerr, ok := err.(*errors.Error); ok {
		err = goerr.Err

		for _, s := range goerr.StackFrames() {
			frame := &honeybadger.Frame{
				File:   s.File,
				Number: strconv.Itoa(s.LineNumber),
				Method: s.Package + "." + s.Name,
			}
			frames = append(frames, frame)
		}
	}

	return frames
}
