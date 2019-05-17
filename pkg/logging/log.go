package logging

import (
	"fmt"
	"io"
	"log"
	"path/filepath"
	"runtime"
)

// LogPrinter is the function returned by NewLogger for writing log messages
type LogPrinter func(format interface{}, v ...interface{})

// NewLogger returns a LogPrinter function for logging formatted messages.
func NewLogger(out io.Writer, prefix string) LogPrinter {
	if out != nil {
		logger := log.New(out, prefix, log.LstdFlags)
		return func(format interface{}, v ...interface{}) {
			// simulate log.Lshortfile and dynamically create file:line header
			// runtime.Caller(callerdepth = 1) to get file:line from caller
			_, file, line, ok := runtime.Caller(1)
			if !ok {
				file = "???"
				line = 0
			}
			fhdr := fmt.Sprintf("%s:%d", filepath.Base(file), line)
			format = fmt.Sprintf("%s: %s", fhdr, format)
			logger.Printf(format.(string), v...)
		}
	} else {
		return func(format interface{}, v ...interface{}) { /* noop */ }
	}
}
