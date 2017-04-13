package hbsvc

import (
	"bytes"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/honeybadger-io/honeybadger-go"
	"github.com/maruel/panicparse/stack"
)

// NotifyPanic notifies a panic event with useful contextual informations.
// It analyzes a buffer string from the output of panic,
// and creates error and context objects that are eligible for honeybadger
func NotifyPanic(output string) {
	outputLines := strings.Split(output, "\n")
	hbErr := honeybadger.NewError(outputLines[0])

	in := bytes.NewBufferString(output)
	goroutines, err := stack.ParseDump(in, ioutil.Discard)
	if err != nil {
		honeybadger.Notify(
			hbErr,
			honeybadger.ErrorClass{Name: "panicwrap"},
			honeybadger.Context{
				"RawOutput": outputLines,
			},
		)
		return
	}

	frames := make([]*honeybadger.Frame, 0)
	buckets := stack.SortBuckets(stack.Bucketize(goroutines, stack.AnyValue))
	for _, b := range buckets {
		for _, r := range b.Routines {
			for _, c := range r.Stack.Calls {
				frame := &honeybadger.Frame{
					File:   c.SourcePath,
					Number: strconv.Itoa(c.Line),
					Method: c.Func.PkgDotName(),
				}
				frames = append(frames, frame)
			}
		}
	}

	if len(frames) > 0 {
		hbErr.Stack = frames
	}

	honeybadger.Notify(
		hbErr,
		honeybadger.ErrorClass{Name: "panicwrap"},
		honeybadger.Context{
			"RawOutput": outputLines,
		},
	)
}
