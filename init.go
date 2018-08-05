package gterm

import (
	"os"
	"github.com/subfuzion/gterm/logging"
)

// Logger is the default LogPrinter for the gterm package.
// Logging is only enabled when the DEBUG environment variable is set (DEBUG=1)
var Logger logging.LogPrinter

func init() {
	prefix := ""
	if v := os.Getenv("DEBUG"); v == "1" {
		f, err := os.OpenFile("gterm.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
		if err != nil {
			panic(err)
		}
		Logger = logging.NewLogger(f, prefix)
	} else {
		Logger = logging.NewLogger(nil, prefix)
	}
}
