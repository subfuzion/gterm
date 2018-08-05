package terminal

import (
	"fmt"

	"github.com/subfuzion/gterm"
	"github.com/subfuzion/gterm/driver"
	"github.com/subfuzion/gterm/logging"
	"github.com/subfuzion/gterm/style"
)

// New initializes a new Terminal instance.
func New() (Interface, error) {
	t := &term{
		log: gterm.Logger,
		driver: driver.NewTermbox(gterm.Logger),
	}

	t.log("Init")
	if err := t.driver.Init(); err != nil {
		return nil, &driverError{
			"term initialization failed",
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


// term implements the Terminal interface and controls the underlying
// term device through the Driver interface.
type term struct {
	driver driver.Driver
	log logging.LogPrinter
	fg style.CellStyle
	bg style.CellStyle
}

func (t term) End() {
	t.log("End")
	t.driver.End()
}

func (t term) Refresh() {
	t.log("Refresh")
	t.driver.Refresh()
}

func (t *term) SetStyle(fg, bg style.CellStyle) {
	t.log("SetStyle(%+v, %+v)", fg, bg)
	t.fg = fg
	t.bg = bg
}

func (t term) GetStyle() (fg, bg style.CellStyle) {
	return t.fg, t.bg
}

func (t term) Fill(r rune) {
	t.log("NOT IMPLEMENTED: Fill('%c')", r)
}

func (t term) Clear() {
	t.log("Clear(%+v, %+v)", t.fg, t.bg)
	t.driver.Clear(t.fg, t.bg)
}

