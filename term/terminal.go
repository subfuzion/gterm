package term

import (
	"github.com/subfuzion/gterm"
	"github.com/subfuzion/gterm/driver"
	"github.com/subfuzion/gterm/logging"
	"github.com/subfuzion/gterm/style"
	"fmt"
)

// New initializes a new Terminal instance.
func New() (gterm.Terminal, error) {
	t := &terminal{
		log: gterm.Logger,
		driver: driver.NewTermbox(gterm.Logger),
	}

	t.log("Init")
	if err := t.driver.Init(); err != nil {
		return nil, &driverError{
			"terminal initialization failed",
			err.Error(),
		}
	}
	return t, nil
}

type driverError struct {
	err string
	driverErr string
}

func (e *driverError) Error() string {
	return fmt.Sprintf("error: %s (driver error: '%s')", e.err, e.driverErr)
}


// terminal implements the Terminal interface and controls the underlying
// terminal device through the Driver interface.
type terminal struct {
	driver driver.Driver
	log logging.LogPrinter
	fg style.CellStyle
	bg style.CellStyle
}

func (t terminal) End() {
	t.log("End")
	t.driver.End()
}

func (t terminal) Refresh() {
	t.log("Refresh")
	t.driver.Refresh()
}

func (t *terminal) SetStyle(fg, bg style.CellStyle) {
	t.log("SetStyle(%+v, %+v)", fg, bg)
	t.fg = fg
	t.bg = bg
}

func (t terminal) GetStyle() (fg, bg style.CellStyle) {
	return t.fg, t.bg
}

func (t terminal) Fill(r rune) {
	t.log("NOT IMPLEMENTED: Fill('%c')", r)
}

func (t terminal) Clear() {
	t.log("Clear(%+v, %+v)", t.fg, t.bg)
	t.driver.Clear(t.fg, t.bg)
}

