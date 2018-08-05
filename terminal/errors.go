package terminal

import (
	"errors"
	"fmt"

	"github.com/subfuzion/gterm/logging"
)

var ErrOutOfBoundsRequestMsg = "request to render out of terminal bounds"

func ErrOutOfBoundsRequest(log logging.LogPrinter, bounds interface{}, args ...interface{}) error {
	msg := fmt.Sprintf("ERROR: %s (bounds: %s, args: %s)", ErrOutOfBoundsRequestMsg, bounds, fmt.Sprint(args...))
	if log != nil {
		log(msg)
	}
	return errors.New(msg)
}
